package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorHandlerService struct {
	logger *logrus.Logger
}

func NewErrorHandlerService(logger *logrus.Logger) *ErrorHandlerService {
	return &ErrorHandlerService{
		logger: logger,
	}
}

func (s *ErrorHandlerService) HandleError(webHandler func(c *gin.Context) (string, error)) (ginhandler func(c *gin.Context)) {
	return func(c *gin.Context) {
		if status, err := webHandler(c); err != nil {
			s.logger.Error(err)
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": status, "message": err.Error()})
		}
	}
}
