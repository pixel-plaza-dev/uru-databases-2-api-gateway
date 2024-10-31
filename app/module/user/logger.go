package user

import (
	commonlogger "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/logger"
)

// Logger is the logger for the user service
type Logger struct {
	logger commonlogger.Logger
}

// NewLogger is the logger for the user service
func NewLogger(logger commonlogger.Logger) Logger {
	return Logger{logger: logger}
}
