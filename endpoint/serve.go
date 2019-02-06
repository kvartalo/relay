package endpoint

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewApiService() *gin.Engine {
	api := gin.Default()
	api.Use(cors.Default())
	api.GET("/info", handleInfo)
	api.GET("/balance/:addr", handleGetBalance)
	api.GET("/history/:addr", handleGetHistory)
	api.GET("/tx/nonce/:addr", handleGetTxNonce)
	api.POST("/tx/:addr", handlePostTx)
	api.GET("/polls", handleGetPolls)
	api.GET("/polls/:id", handleGetPoll)
	return api
}
