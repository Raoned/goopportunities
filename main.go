package main

import (
	"github.com/Raoned/goopportunities/config"
	"github.com/Raoned/goopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initialize config
	err := config.Init
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	//Initialize router - maiusculo variavel global e minusculo local
	router.Initialize()

}
