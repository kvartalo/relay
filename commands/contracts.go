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
				{
					Name:   "transfer",
					Usage:  "tranfer tokens to a specified address",
					Action: cmdTokenTransfer,
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

func stringToBigInt(s string) (*big.Int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}
	bigi := big.NewInt(int64(i))
	return bigi, nil
}

func cmdTokenMint(c *cli.Context) error {
	amountStr := c.Args().Get(0)
	if amountStr == "" {
		color.Red("No amount specified. Usage: contracts token mint [amount]")
		os.Exit(0)
	}
	amount, err := stringToBigInt(amountStr)
	if err != nil {
		color.Red("Amount parsing error. Usage: contracts token mint [amount]")
		os.Exit(0)
	}

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

func cmdTokenTransfer(c *cli.Context) error {
	amountStr := c.Args().Get(1)
	if amountStr == "" {
		os.Exit(0)
	}
	amount, err := stringToBigInt(amountStr)
	if err != nil {
		color.Red("No amount specified. Usage: contracts token transfer [addr] [amount]")
		os.Exit(0)
	}

	toAddrHex := c.Args().Get(0)
	if toAddrHex == "" {
		color.Red("No address specified. Usage: contracts token transfer [addr] [amount]")
		os.Exit(0)
	}
	toAddr := common.HexToAddress(toAddrHex)

	ethSrv := startRelay(c)

	tokenContractAddr := common.HexToAddress(config.C.Contracts.Token)
	ethSrv.LoadTokenContract(tokenContractAddr)

	fromAddr := common.HexToAddress(config.C.Keystorage.Address)
	tokenAddr := common.HexToAddress(config.C.Contracts.Token)

	auth, err := eth.GetAuth()
	if err != nil {
		return err
	}

	nonce, err := ethSrv.Token.NonceOf(nil, fromAddr)
	if err != nil {
		return err
	}

	// build tx msg
	var msgBytes []byte
	msgBytes = append(msgBytes, []byte{0x19}...)
	msgBytes = append(msgBytes, []byte{0}...)
	msgBytes = append(msgBytes, tokenAddr.Bytes()...)
	msgBytes = append(msgBytes, nonce.Bytes()...)
	msgBytes = append(msgBytes, fromAddr.Bytes()...)
	msgBytes = append(msgBytes, toAddr.Bytes()...)
	msgBytes = append(msgBytes, amount.Bytes()...)

	// sign msg
	sig, err := ethSrv.SignBytes(msgBytes)
	if err != nil {
		return err
	}
	// get r, s, v from sig
	var r, s [32]byte
	copy(r[:], sig[:32])
	copy(s[:], sig[32:64])
	v := sig[64]

	// transfer token
	tx, err := ethSrv.Token.Transfer(auth, fromAddr, toAddr, amount, r, s, v)
	if err != nil {
		return err
	}

	color.Green("Token transfer success, tx: " + tx.Hash().Hex())
	return nil
}
