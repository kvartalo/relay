package eth

import (
	"context"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
	"github.com/kvartalo/relay/config"
)

func ConnectWeb3(ks *keystore.KeyStore, acc *accounts.Account) *ethclient.Client {
	client, err := ethclient.Dial(config.C.Web3.Url)
	if err != nil {
		color.Red("Can not open connection to web3 (config.Web3.Url: " + config.C.Web3.Url + ")\n" + err.Error() + "\n")
		os.Exit(0)
	}
	color.Green("Connection to web3 server opened")
	return client
}

func GetBalance(web3client *ethclient.Client, address common.Address) (*big.Float, error) {
	balance, err := web3client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return nil, err
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethBalance := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return ethBalance, nil
}
