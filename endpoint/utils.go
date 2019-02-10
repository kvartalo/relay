package endpoint

import (
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func check(c *gin.Context, err error) {
	if err != nil {
		color.Red("error: " + err.Error())
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
}
