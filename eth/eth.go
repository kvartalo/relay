//go:generate solc --abi --bin token/contracts/token.sol -o build
//go:generate abigen --bin=./build/Token.bin --abi=./build/Token.abi --pkg=token --out=token.go

package eth

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
	"github.com/kvartalo/relay/config"
	"github.com/kvartalo/relay/eth/token"
	"github.com/kvartalo/relay/storage"
	"github.com/kvartalo/relay/utils"
	log "github.com/sirupsen/logrus"
)

type EthService struct {
	storage      *storage.Storage
	ks           *keystore.KeyStore
	acc          *accounts.Account
	client       *ethclient.Client
	Token        *token.Token
	Scanner      *ScanEventDispatcher
	tokenAddress common.Address
	tokenAbi     *abi.ABI
}

func NewEthService(ks *keystore.KeyStore, acc *accounts.Account, storage *storage.Storage) *EthService {
	client, err := ethclient.Dial(config.C.Web3.Url)
	if err != nil {
		color.Red("Can not open connection to web3 (config.Web3.Url: " + config.C.Web3.Url + ")\n" + err.Error() + "\n")
		os.Exit(0)
	}
	color.Green("Connection to web3 server opened")

	tokenAbi, err := abi.JSON(strings.NewReader(token.TokenABI))
	if err != nil {
		log.Panic(err)
	}

	service := &EthService{
		storage:  storage,
		ks:       ks,
		acc:      acc,
		client:   client,
		tokenAbi: &tokenAbi,
	}

	service.Scanner = NewScanEventDispatcher(client, service.scannedTx, service)

	return service
}

func (ethSrv *EthService) Account() *accounts.Account {
	return ethSrv.acc
}

func (ethSrv *EthService) Client() *ethclient.Client {
	return ethSrv.client
}

func (ethSrv *EthService) ConnectWeb3() *ethclient.Client {
	client, err := ethclient.Dial(config.C.Web3.Url)
	if err != nil {
		color.Red("Can not open connection to web3 (config.Web3.Url: " + config.C.Web3.Url + ")\n" + err.Error() + "\n")
		os.Exit(0)
	}
	color.Green("Connection to web3 server opened")
	return client
}

func (ethSrv *EthService) GetBalance(address common.Address) (*big.Float, error) {
	balance, err := ethSrv.client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return nil, err
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethBalance := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return ethBalance, nil
}

// DeployTokenContract deploys the Token contract to eth network
func (ethSrv *EthService) DeployTokenContract() error {
	auth, err := GetAuth()
	if err != nil {
		return err
	}

	address, tx, _, err := token.DeployToken(auth, ethSrv.client)
	if err != nil {
		return err
	}
	ethSrv.tokenAddress = address

	color.Green("token contract deployed at address: " + address.Hex())
	fmt.Println("deployment transaction: " + tx.Hash().Hex())
	return nil
}

// LoadTokenContract loads already deployed Token contract
func (ethSrv *EthService) LoadTokenContract(contractAddr common.Address) {
	instance, err := token.NewToken(contractAddr, ethSrv.client)
	if err != nil {
		color.Red(err.Error())
	}
	ethSrv.tokenAddress = contractAddr
	ethSrv.Token = instance
}

type SignatureEthMsg [65]byte

// SignBytes performs a keccak256 hash over the bytes and signs it
func (ethSrv *EthService) SignBytes(b []byte) (*SignatureEthMsg, error) {
	// h := utils.EthHash(b)
	h := utils.HashBytes(b)
	sigBytes, err := ethSrv.ks.SignHash(*ethSrv.acc, h[:])
	if err != nil {
		return nil, err
	}
	sig := &SignatureEthMsg{}
	copy(sig[:], sigBytes)

	return sig, nil
}

func GetAuth() (*bind.TransactOpts, error) {
	file, err := os.Open(config.C.Keystorage.KeyJsonPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewTransactor(strings.NewReader(string(b)), config.C.Keystorage.Password)
	return auth, err
}

// scandipatcher events ----------------------------------------------

func (ethSrv *EthService) scannedTx(block *types.Block, tx *types.Transaction, rcpt *types.Receipt) error {

	transferEvent, ok := ethSrv.tokenAbi.Events["Transfer"]
	if !ok {
		panic(fmt.Errorf("Event Transfer not found"))
	}
	transferEventTopic := transferEvent.Id()

	to := tx.To()
	if to == nil || !bytes.Equal((*to)[:], ethSrv.tokenAddress[:]) {
		return nil
	}
	for _, logevent := range rcpt.Logs {
		if logevent.Removed {
			continue
		}
		if !bytes.Equal(logevent.Topics[0][:], transferEventTopic[:]) {
			continue
		}
		type TransferEvent struct {
			From  common.Address
			To    common.Address
			Value *big.Int
			IsTax bool
		}

		var event TransferEvent
		err := ethSrv.tokenAbi.Unpack(&event, "Transfer", logevent.Data)
		if err != nil {
			return err
		}

		err = ethSrv.storage.AddTransfer(&storage.Transfer{
			From:      event.From,
			To:        event.To,
			Value:     event.Value,
			Timestamp: block.Time().Uint64(),
		})
		if err != nil {
			return err
		}

	}
	return nil
}

func (ethSrv *EthService) SavePointLoad() (lastBlock uint64, lastTxIndex uint, err error) {

	sp, err := ethSrv.storage.LoadSavePoint()
	if err != nil {
		return 0, 0, nil
	}
	return sp.LastBlock, sp.LastTxIndex, nil
}

func (ethSrv *EthService) SavePointSave(lastBlock uint64, lastTxIndex uint) error {

	return ethSrv.storage.SetSavePoint(storage.SavePoint{
		LastBlock:   lastBlock,
		LastTxIndex: lastTxIndex,
	})
}
