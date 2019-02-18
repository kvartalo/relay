package commands

import (
	"github.com/kvartalo/relay/config"
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
	{
		Name:    "info",
		Aliases: []string{},
		Usage:   "get info about the server",
		Action:  cmdInfo,
	},
}

func cmdStart(c *cli.Context) error {
	ethSrv := loadRelay(c)

	// run the service
	apiService := endpoint.Serve(config.C, ethSrv)
	apiService.Run(":" + config.C.Server.Port)

	return nil
}

func cmdInfo(c *cli.Context) error {
	_ = loadRelay(c)

	return nil
}
