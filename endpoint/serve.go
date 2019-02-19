package endpoint

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kvartalo/relay/config"
	"github.com/kvartalo/relay/eth"
	"github.com/kvartalo/relay/storage"
)

var ethSrv *eth.EthService
var serverConfig config.Config
var sto *storage.Storage

func newApiService() *gin.Engine {
	api := gin.Default()
	api.Use(cors.Default())
	api.GET("/info", handleInfo)
	api.GET("/balance/:addr", handleGetBalance)
	api.GET("/history/:addr", handleGetHistory)
	api.GET("/tx/nonce/:addr", handleGetTxNonce)
	api.POST("/tx", handlePostTx)
	api.GET("/polls", handleGetPolls)    // TODO
	api.GET("/polls/:id", handleGetPoll) // TODO
	return api
}

func Serve(cnfg config.Config, stor *storage.Storage, ethService *eth.EthService) *gin.Engine {
	sto = stor
	ethSrv = ethService
	serverConfig = cnfg
	return newApiService()
}
