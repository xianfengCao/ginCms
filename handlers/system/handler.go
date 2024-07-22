package system

import (
	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	//fmt.Println(utils.DB())
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
