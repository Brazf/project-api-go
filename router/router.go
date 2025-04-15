package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// INITIALIZE ROUTER
	router := gin.Default()

	// INITIALIZE ROUTES
	initializeRoutes(router)

	// RUN THE SERVER
	router.Run() // listen and serve on 0.0.0.0:8080
}
