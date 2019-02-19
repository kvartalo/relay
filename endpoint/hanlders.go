package endpoint

import (
	"context"
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/kvartalo/relay/eth"
	"github.com/kvartalo/relay/storage"
)

func handleInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"eth": "current amount",
	})
}

func handleGetBalance(c *gin.Context) {
	addrHex := c.Param("addr")
	addr := common.HexToAddress(addrHex)
	balance, err := ethSrv.Token.BalanceOf(nil, addr)
	if err != nil {
		fail(c, err)
		return
	}
	c.JSON(200, gin.H{
		"addr":    addr,
		"balance": balance.String(),
	})
}

const pageSize = 20

// no history for the moment
func handleGetHistory(c *gin.Context) {

	addrHex := c.Param("addr")
	addr := common.HexToAddress(addrHex)
	count, err := sto.GetTransferCount(&addr)
	if err != nil {
		fail(c, err)
		return
	}

	var countstart int64
	if count < pageSize {
		countstart = 0
	} else {
		countstart = int64(count) - pageSize
	}

	transfers := []storage.Transfer{}
	for i := int64(count) - 1; i >= countstart; i-- {
		transfer, err := sto.GetTransfer(&addr, uint64(i))
		if err != nil {
			fail(c, err)
			return
		}
		transfers = append(transfers, *transfer)
	}

	c.JSON(200, gin.H{
		"addr":      addr,
		"count":     count,
		"transfers": transfers,
	})
}

func handleGetTxNonce(c *gin.Context) {
	addrHex := c.Param("addr")
	addr := common.HexToAddress(addrHex)
	nonce, err := ethSrv.Token.NonceOf(nil, addr)
	if err != nil {
		fail(c, err)
		return
	}
	c.JSON(200, gin.H{
		"addr":  addr,
		"nonce": nonce,
	})
}

type TxMsg struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value int    `json:"value"`
	R     string `json:"r"`
	S     string `json:"s"`
	V     int    `json:"v"`
}

func handlePostTx(c *gin.Context) {
	var tx TxMsg
	c.BindJSON(&tx)

	rBytes, err := hex.DecodeString(tx.R)
	if err != nil {
		fail(c, err)
		return
	}
	var r32 [32]byte
	copy(r32[:], rBytes)
	sBytes, err := hex.DecodeString(tx.S)
	if err != nil {
		fail(c, err)
		return
	}
	var s32 [32]byte
	copy(s32[:], sBytes)

	fromAddr := common.HexToAddress(serverConfig.Keystorage.Address)
	nonce, err := ethSrv.Client().PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		fail(c, err)
		return
	}

	gasPrice, err := ethSrv.Client().SuggestGasPrice(context.Background())
	if err != nil {
		fail(c, err)
		return
	}

	auth, err := eth.GetAuth()
	if err != nil {
		fail(c, err)
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	ethTx, err := ethSrv.Token.Transfer(auth, common.HexToAddress(tx.From), common.HexToAddress(tx.To), big.NewInt(int64(tx.Value)), r32, s32, byte(tx.V))
	if err != nil {
		fail(c, err)
		return
	}
	c.JSON(200, gin.H{
		"ethTx": ethTx.Hash().Hex(),
	})
}

func handleGetPolls(c *gin.Context) {
	// TODO
	c.JSON(200, gin.H{
		"polls": "polls",
	})
}

func handleGetPoll(c *gin.Context) {
	// TODO
	pollId := c.Param("id")
	c.JSON(200, gin.H{
		"pollId": pollId,
	})
}
