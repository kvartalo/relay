package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fatih/color"
	"github.com/kvartalo/relay/config"
	"github.com/kvartalo/relay/eth"
	"github.com/kvartalo/relay/storage"
	"github.com/urfave/cli"
)

func loadStorage(c *cli.Context) *storage.Storage {
	if err := config.MustRead(c); err != nil {
		color.Red(err.Error())
		os.Exit(0)
	}
	sto, err := storage.New(config.C.Storage.Path)
	if err != nil {
		color.Red(err.Error())
		os.Exit(0)
	}
	return sto
}

// loadRelay does:
// - reads the configuration from config.yaml
// - opens the KeyStorage specified in the configuration, creating a new keystorage and account
// - creates a new EthService
// - prints the balance of the Relay wallet
func loadRelay(c *cli.Context, storage *storage.Storage) *eth.EthService {
	if err := config.MustRead(c); err != nil {
		color.Red(err.Error())
		os.Exit(0)
	}

	ks, account, err := importKeystorage(config.C.Keystorage.Address, config.C.Keystorage.Password)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	fmt.Println("Keystore with addr " + account.Address.Hex() + " opened")

	ethSrv := eth.NewEthService(ks, account, storage)

	// get current block number
	header, err := ethSrv.Client().HeaderByNumber(context.Background(), nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	fmt.Println("block number: " + header.Number.String())

	// get relay balance
	balance, err := ethSrv.GetBalance(account.Address)
	if err != nil {
		color.Red(err.Error())
		os.Exit(0)
	}
	color.Cyan("current ether balance: " + balance.String() + " ETH\n")

	addr := common.HexToAddress(config.C.Keystorage.Address)
	tokenContractAddr := common.HexToAddress(config.C.Contracts.Token)
	ethSrv.LoadTokenContract(tokenContractAddr)
	tokenBalance, err := ethSrv.Token.BalanceOf(nil, addr)
	if err != nil {
		color.Red(err.Error())
		os.Exit(0)
	}
	color.Magenta("current token balance: " + tokenBalance.String() + " Tokens\n")

	return ethSrv
}

func importKeystorage(addr string, password string) (*keystore.KeyStore, *accounts.Account, error) {
	file := KEYSPATH
	ks := keystore.NewKeyStore(file, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.Find(accounts.Account{
		Address: common.HexToAddress(addr),
	})
	if err != nil {
		return nil, nil, err
	}

	err = ks.Unlock(account, config.C.Keystorage.Password)
	if err != nil {
		return nil, nil, err
	}

	return ks, &account, nil
}
