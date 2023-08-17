package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestConnection(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Server status: OK")
}
