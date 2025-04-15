package handler

import "github.com/gin-gonic/gin"

func DeleteOpeningHeadler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "DELETE Opening",
	})
}
