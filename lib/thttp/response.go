package thttp

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func SendOkResponse(c *gin.Context, response any) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)

	if reflect.TypeOf(response) == reflect.TypeOf(nil) {
		return
	}

	responseJson, _ := json.Marshal(response)
	_, _ = c.Writer.Write(responseJson)
}

func SendErrorResponse(c *gin.Context, statusCode int, message string) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(statusCode)

	responseJson, _ := json.Marshal(ErrorResponse{
		Message: message,
	})

	_, _ = c.Writer.Write(responseJson)
}
