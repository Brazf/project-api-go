package router

import (
	"github.com/Brazf/project-api-go/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHeadler)

		v1.POST("/opening", handler.CreateOpeningHeadler)

		v1.DELETE("/opening", handler.DeleteOpeningHeadler)

		v1.PUT("/opening", handler.UpdateOpeningHeadler)

		v1.GET("/openings", handler.ListOpeningHeadler)
	}

}
