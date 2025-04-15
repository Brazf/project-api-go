package handler

import "github.com/gin-gonic/gin"

func UpdateOpeningHeadler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "PUT Opening",
	})
}
