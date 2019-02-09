package commands

import (
	"github.com/urfave/cli"
)

var ContractsCommands = []cli.Command{
	{
		Name:  "contracts",
		Usage: "deploy contracts",
		Subcommands: []cli.Command{{
			Name:  "token",
			Usage: "for token smart contract",
			Subcommands: []cli.Command{{
				Name:   "deploy",
				Usage:  "deploy token smart contract",
				Action: cmdTokenDeploy,
			},
			},
		}, // in the future here will come more contracts
		},
	},
}

func cmdTokenDeploy(c *cli.Context) error {
	ethSrv := startRelay(c)
	err := ethSrv.DeployTokenContract()
	return err
}
