package handler

import "github.com/gin-gonic/gin"

func ListOpeningHeadler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "GET Openings",
	})
}
