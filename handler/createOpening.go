package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}
