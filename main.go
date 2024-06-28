package main

import (
	"github.com/brunobotter/ecommerce-carrinho/configs"
	"github.com/brunobotter/ecommerce-carrinho/router"
)

var (
	logger *configs.Logger
)

func main() {
	logger = configs.GetLogger("main")
	//initialize configs
	err := configs.Init()
	if err != nil {
		logger.Errorf("config initializate error: %v", err)
		return
	}
	//initialize router
	router.Initialize()

}
