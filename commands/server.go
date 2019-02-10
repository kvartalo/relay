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

// startRelay does:
// - reads the configuration from config.yaml
// - opens the KeyStorage specified in the configuration, creating a new keystorage and account
// - creates a new EthService
// - prints the balance of the Relay wallet
func startRelay(c *cli.Context) *eth.EthService {
	if err := config.MustRead(c); err != nil {
		color.Red(err.Error())
		os.Exit(0)
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
		os.Exit(0)
	}
	color.Cyan("current balance: " + balance.String() + " ETH\n")

	return ethSrv
}

func cmdStart(c *cli.Context) error {
	ethSrv := startRelay(c)
	tokenContractAddr := common.HexToAddress(config.C.Contracts.Token)
	ethSrv.LoadTokenContract(tokenContractAddr)

	// run the service
	apiService := endpoint.Serve(ethSrv)
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
