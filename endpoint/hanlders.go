package endpoint

import (
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/kvartalo/relay/eth"
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
	check(c, err)
	c.JSON(200, gin.H{
		"addr":    addr,
		"balance": balance.String(),
	})
}

// no history for the moment
func handleGetHistory(c *gin.Context) {
	addrHex := c.Param("addr")
	addr := common.HexToAddress(addrHex)
	c.JSON(200, gin.H{
		"addr": addr,
	})
}

func handleGetTxNonce(c *gin.Context) {
	addrHex := c.Param("addr")
	addr := common.HexToAddress(addrHex)
	nonce, err := ethSrv.Token.NonceOf(nil, addr)
	check(c, err)
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
	check(c, err)
	var r32 [32]byte
	copy(r32[:], rBytes)
	sBytes, err := hex.DecodeString(tx.S)
	check(c, err)
	var s32 [32]byte
	copy(s32[:], sBytes)

	auth, err := eth.GetAuth()
	check(c, err)
	ethTx, err := ethSrv.Token.Transfer(auth, common.HexToAddress(tx.From), common.HexToAddress(tx.To), big.NewInt(int64(tx.Value)), r32, s32, byte(tx.V))
	check(c, err)
	c.JSON(200, gin.H{
		"ethTx": ethTx.Hash().Hex(),
	})
}

func handleGetPolls(c *gin.Context) {
	c.JSON(200, gin.H{
		"polls": "polls",
	})
}

func handleGetPoll(c *gin.Context) {
	pollId := c.Param("id")
	c.JSON(200, gin.H{
		"pollId": pollId,
	})
}
