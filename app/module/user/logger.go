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

// ReceivedSignUpRequest logs the received sign up request
func (l *Logger) ReceivedSignUpRequest() {
	l.logger.LogMessage("Received sign up request")
}

// SignUpSuccess logs the sign-up success
func (l *Logger) SignUpSuccess() {
	l.logger.LogMessage("Sign-up success")
}

// SignUpFailed logs the sign-up failure
func (l *Logger) SignUpFailed(err error) {
	l.logger.LogMessageWithDetails("Sign-up failed", err.Error())
}
