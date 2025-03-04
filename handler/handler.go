package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Welcome ... Welcome
// @Summary Welcome
// @Description Welcome
// @Tags Welcome
// @Success 200 {object} object
// @Failure 404 {object} object
// @Router /welcome [get]
func (handler Handler) Welcome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Test",
	})
}
