package main

import (
	"log"
	"os"

	"github.com/kvartalo/relay/commands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "relay"
	app.Version = "0.0.1-alpha"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "config"},
	}

	app.Commands = []cli.Command{}
	app.Commands = append(app.Commands, commands.WalletCommands...)
	app.Commands = append(app.Commands, commands.ServerCommands...)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
