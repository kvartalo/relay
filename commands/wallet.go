package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

const KEYSPATH = "./keystorage/"

var WalletCommands = []cli.Command{
	{
		Name:  "wallet",
		Usage: "create and manage eth wallet",
		Subcommands: []cli.Command{{
			Name:   "new",
			Usage:  "create new Keystorage. Args: keystoreId password",
			Action: cmdWalletNew,
		},
			{
				Name:   "info",
				Usage:  "display info about the current relay wallets",
				Action: cmdWalletInfo,
			}},
	},
}

func cmdWalletNew(c *cli.Context) error {
	password := c.Args().Get(0)
	if password == "" {
		color.Red("No password specified. Usage: wallet new [password]")
		os.Exit(0)
	}

	ks := keystore.NewKeyStore(KEYSPATH, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	color.Green("New Keystore with address '" + account.Address.Hex() + "' created in '" + KEYSPATH + "'")
	return nil
}

func cmdWalletInfo(c *cli.Context) error {
	files, err := ioutil.ReadDir(KEYSPATH)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("current Keystores:")
	for _, f := range files {
		fmt.Println("	", f.Name())
	}
	return nil
}
