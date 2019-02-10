package eth

import (
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
	"github.com/kvartalo/relay/config"
	"github.com/kvartalo/relay/eth/token"
	"github.com/kvartalo/relay/utils"
)

type EthService struct {
	ks     *keystore.KeyStore
	acc    *accounts.Account
	client *ethclient.Client
	Token  *token.Token
}

func NewEthService(ks *keystore.KeyStore, acc *accounts.Account) *EthService {
	client, err := ethclient.Dial(config.C.Web3.Url)
	if err != nil {
		color.Red("Can not open connection to web3 (config.Web3.Url: " + config.C.Web3.Url + ")\n" + err.Error() + "\n")
		os.Exit(0)
	}
	color.Green("Connection to web3 server opened")
	return &EthService{
		ks:     ks,
		acc:    acc,
		client: client,
	}
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
	// fromAddress := ethSrv.acc.Address
	// nonce, err := ethSrv.client.PendingNonceAt(context.Background(), fromAddress)
	// if err != nil {
	//         return err
	// }
	//
	// gasPrice, err := ethSrv.client.SuggestGasPrice(context.Background())
	// if err != nil {
	//         return err
	// }
	auth, err := GetAuth()
	if err != nil {
		return err
	}

	taxDestination := ethSrv.acc.Address
	address, tx, instance, err := token.DeployToken(auth, ethSrv.client, taxDestination)
	if err != nil {
		return err
	}
	_ = instance

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
	ethSrv.Token = instance
}

type SignatureEthMsg [65]byte

func (ethSrv *EthService) SignBytes(b []byte) (*SignatureEthMsg, error) {
	h := utils.EthHash(b)
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
