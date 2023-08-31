package interfaces

import (
	"github.com/gin-gonic/gin"
)

type IErrorHandlerService interface {
	HandleError(webHandler func(c *gin.Context) (string, error)) (ginhandler func(c *gin.Context))
}
