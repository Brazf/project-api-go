package main

import (
	"github.com/Brazf/project-api-go/config"
	"github.com/Brazf/project-api-go/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")
	//Initializer config
	err := config.Init()

	//Check config error
	if err != nil {
		logger.Errorf("DEU ERRRO: %v", err)
		return
	}
	//Initializer routes
	router.Initialize()
}
