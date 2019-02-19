package commands

import (
	"github.com/kvartalo/relay/config"
	"github.com/kvartalo/relay/endpoint"
	"github.com/kvartalo/relay/storage"
	"github.com/urfave/cli"
)

var ServerCommands = []cli.Command{
	{
		Name:    "init",
		Aliases: []string{},
		Usage:   "initialize the database",
		Action:  cmdInit,
	},
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
	sto := loadStorage(c)
	sto.RawDump()
	ethSrv := loadRelay(c, sto)
	ethSrv.Scanner.Start()
	// run the service
	apiService := endpoint.Serve(config.C, sto, ethSrv)
	apiService.Run(":" + config.C.Server.Port)

	return nil
}

func cmdInfo(c *cli.Context) error {
	storage := loadStorage(c)
	_ = loadRelay(c, storage)

	return nil
}

func cmdInit(c *cli.Context) error {
	sto := loadStorage(c)
	sto.SetSavePoint(storage.SavePoint{
		LastBlock:   config.C.Web3.StartScanBlock,
		LastTxIndex: 0,
	})

	return nil
}
