package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HomeController struct {
	logger *logrus.Logger
}

func NewHomeController(logger *logrus.Logger) *HomeController {
	return &HomeController{
		logger: logger,
	}
}

func (c *HomeController) TestConnection(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Server status: OK")
}
