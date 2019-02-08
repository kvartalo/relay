package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
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

func cmdWalletNew(c *cli.Context) {
	keystoreId := c.Args().Get(0)
	if keystoreId == "" {
		fmt.Println("usage: wallet new [keystoreId] [password]")
		os.Exit(0)
	}
	password := c.Args().Get(1)
	if password == "" {
		fmt.Println("usage: wallet new [keystoreId] [password]")
		os.Exit(0)
	}

	ks := keystore.NewKeyStore(KEYSPATH+keystoreId, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("New Keystore created in '" + KEYSPATH + keystoreId + "'")
	fmt.Println("addr: ", account.Address.Hex())
}

func cmdWalletInfo(c *cli.Context) {
	files, err := ioutil.ReadDir(KEYSPATH)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("current Keystores:")
	for _, f := range files {
		fmt.Println("	", f.Name())
	}
}
