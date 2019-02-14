package endpoint

import (
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func fail(c *gin.Context, err error) {
	color.Red("error: " + err.Error())
	c.JSON(400, gin.H{
		"error": err.Error(),
	})
}
