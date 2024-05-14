package handler

import (
	"github.com/gin-gonic/gin"
	"test-task-sw/lib/thttp"
)

// Ping godoc
// @Summary     Пинг сервиса
// @Tags		Служебные
// @Accept      json
// @Produce     json
// @Success     200 {object} thttp.ResponseWithDetails[string]
// @Router      /api/ping [get]
func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		thttp.OkResponseWithResult(c, "pong")
	}
}
