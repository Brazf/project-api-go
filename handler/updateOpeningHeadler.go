package handler

import (
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {

	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

}
