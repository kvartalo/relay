package commands

import (
	"fmt"
	"os"

	"github.com/kvartalo/relay/endpoint"
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

func cmdStart(c *cli.Context) {
	keystoreId := c.Args().Get(0)
	if keystoreId == "" {
		fmt.Println("no keystoreId specified")
		fmt.Println("usage: start [keystoreId] [password]")
		os.Exit(0)
	}
	password := c.Args().Get(1)
	if password == "" {
		fmt.Println("no password specified")
		fmt.Println("usage: start [keystoreId] [password]")
		os.Exit(0)
	}
	apiService := endpoint.NewApiService()
	apiService.Run(":3000")
}
