package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler Handler) Welcome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Test",
	})
}
