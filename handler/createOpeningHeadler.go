package handler

import "github.com/gin-gonic/gin"

func CreateOpeningHeadler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "POST Opening",
	})
}
