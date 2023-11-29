package main

import (
	"github.com/maikonalexandre/govagas/config"
	"github.com/maikonalexandre/govagas/router"
)

var logger *config.Logger

func main() {
	logger = config.GetLogger("main")
	err := config.InitDB()

	if err != nil {
		logger.Errf("config inicialization error: %v", err)
		return
	}

	router.Initialize()
}
