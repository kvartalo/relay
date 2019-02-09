package main

import (
	"os"

	"github.com/fatih/color"
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
	app.Commands = append(app.Commands, commands.ContractsCommands...)
	app.Commands = append(app.Commands, commands.ServerCommands...)
	err := app.Run(os.Args)
	if err != nil {
		color.Red(err.Error())
	}
}
