package router

import (
	"github.com/gin-gonic/gin"
	"github.com/robertinho258/ProjetoGo/router/handler"
)

func inicializandoRotas(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.CreatOpeningHandler)
		v1.POST("/opening", handler.ShowOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
	}
}
