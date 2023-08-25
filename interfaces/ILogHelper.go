package interfaces

import (
	"github.com/sirupsen/logrus"
)

type ILogHelper interface {
	GetCallerName() string
	ConfigureLogger() *logrus.Logger
}
