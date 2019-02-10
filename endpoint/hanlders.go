package endpoint

import (
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
	From  common.Address
	To    common.Address
	Value *big.Int
	R     [32]byte
	S     [32]byte
	V     byte
}

func handlePostTx(c *gin.Context) {
	var tx TxMsg
	c.BindJSON(&tx)

	auth, err := eth.GetAuth()
	check(c, err)
	ethTx, err := ethSrv.Token.Transfer(auth, tx.From, tx.To, tx.Value, tx.R, tx.S, tx.V)
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
