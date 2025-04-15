package handler

import "github.com/gin-gonic/gin"

func ShowOpeningHeadler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "GET Opening",
	})
}
