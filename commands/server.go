package commands

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fatih/color"
	"github.com/kvartalo/relay/config"
	"github.com/kvartalo/relay/endpoint"
	"github.com/kvartalo/relay/eth"
	"github.com/urfave/cli"
)

var ServerCommands = []cli.Command{
	{
		Name:    "start",
		Aliases: []string{},
		Usage:   "start the server",
		Action:  cmdStart,
	},
}

func cmdStart(c *cli.Context) error {
	if err := config.MustRead(c); err != nil {
		return err
	}

	ks, account, err := importKeystorage(config.C.Keystorage.Address, config.C.Keystorage.Password)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	color.Cyan("Keystore with addr " + account.Address.Hex() + " opened")

	ethSrv := eth.NewEthService(ks, account)

	// get relay balance
	balance, err := ethSrv.GetBalance(account.Address)
	if err != nil {
		color.Red(err.Error())
	}
	color.Cyan("current balance: " + balance.String() + " ETH\n")

	// run the service
	apiService := endpoint.NewApiService()
	apiService.Run(":" + config.C.Server.Port)

	return nil
}

func importKeystorage(addr string, password string) (*keystore.KeyStore, *accounts.Account, error) {
	file := KEYSPATH
	ks := keystore.NewKeyStore(file, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.Find(accounts.Account{
		Address: common.HexToAddress(addr),
	})

	return ks, &account, err
}
