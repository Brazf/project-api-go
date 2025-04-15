package router

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "GET Opening",
			})
		})
		v1.POST("/opening", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "POST Opening",
			})
		})
		v1.DELETE("/opening", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "DELETE Opening",
			})
		})
		v1.PUT("/opening", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "PUT Opening",
			})
		})
		v1.GET("/openings", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "GET Openings",
			})
		})
	}

}
