package router

import "github.com/gin-gonic/gin"

func Inicio() {
	//Inicializando Rota
	router := gin.Default()
	//Iniciando Rotas
	inicializandoRotas(router)

	//Rodando Servidor
	router.Run() // listen and serve on 0.0.0.0:8080
}
