package endpoint

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func handleInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"eth": "current amount",
	})
}

func handleGetBalance(c *gin.Context) {
	addrHex := c.Param("addr")
	addr := common.HexToAddress(addrHex)
	c.JSON(200, gin.H{
		"addr": addr,
	})
}

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
	c.JSON(200, gin.H{
		"addr":  addr,
		"nonce": "nonce",
	})
}

func handlePostTx(c *gin.Context) {
	addrHex := c.Param("addr")
	addr := common.HexToAddress(addrHex)
	c.JSON(200, gin.H{
		"addr": addr,
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
