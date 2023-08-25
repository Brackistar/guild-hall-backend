package helpers

import (
	"runtime"

	"github.com/sirupsen/logrus"
)

type LogrusHelper struct {
}

func NewLogrusHelper() *LogrusHelper {
	return &LogrusHelper{}
}

func (c *LogrusHelper) GetCallerName() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	return frame.Function
}

func (c *LogrusHelper) ConfigureLogger() *logrus.Logger {
	return logrus.New()
}
