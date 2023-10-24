package main

import (
	"github.com/robertinho258/ProjetoGo/config"
	"github.com/robertinho258/ProjetoGo/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config inictializar error : %v", err)
		return
	}

	//Rota Inicial
	router.Inicio()
}
