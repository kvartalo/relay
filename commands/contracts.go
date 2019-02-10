package commands

import (
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fatih/color"
	"github.com/kvartalo/relay/config"
	"github.com/kvartalo/relay/eth"
	"github.com/urfave/cli"
)

var ContractsCommands = []cli.Command{
	{
		Name:  "contracts",
		Usage: "interact with contracts",
		Subcommands: []cli.Command{{
			Name:  "token",
			Usage: "for token smart contract",
			Subcommands: []cli.Command{
				{
					Name:   "deploy",
					Usage:  "deploy token smart contract",
					Action: cmdTokenDeploy,
				},
				{
					Name:   "mint",
					Usage:  "mint tokens that will go to this Relay address",
					Action: cmdTokenMint,
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

func cmdTokenMint(c *cli.Context) error {
	amountStr := c.Args().Get(0)
	if amountStr == "" {
		color.Red("No amount specified. Usage: contracts token mint [amount]")
		os.Exit(0)
	}
	amountInt, err := strconv.Atoi(amountStr)
	if err != nil {
		color.Red("Amount parsing error. Usage: contracts token mint [amount]")
		os.Exit(0)
	}
	amount := big.NewInt(int64(amountInt))

	ethSrv := startRelay(c)
	tokenContractAddr := common.HexToAddress(config.C.Contracts.Token)
	ethSrv.LoadTokenContract(tokenContractAddr)

	addr := common.HexToAddress(config.C.Keystorage.Address)

	auth, err := eth.GetAuth()
	if err != nil {
		return err
	}
	tx, err := ethSrv.Token.Mint(auth, addr, amount)
	if err != nil {
		return err
	}

	color.Green("mint success, tx: " + tx.Hash().Hex())
	return nil
}
