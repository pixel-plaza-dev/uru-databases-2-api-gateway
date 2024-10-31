package logger

import (
	"github.com/pixel-plaza-dev/uru-databases-2-api-gateway/app/module/user"
	commonenv "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/env"
	commonflag "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/flag"
	commonlistener "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/listener"
	commonlogger "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/logger"
)

const (
	// FlagLoggerName is the name of the flag logger
	FlagLoggerName = "Flag"

	// ListenerLoggerName is the name of the listener logger
	ListenerLoggerName = "Net Listener"

	// EnvironmentLoggerName is the name of the environment logger
	EnvironmentLoggerName = "Environment"

	// UserServiceLoggerName is the name of the user service logger
	UserServiceLoggerName = "User Service"
)

var (
	// FlagLogger is the logger for the flag
	FlagLogger = commonflag.NewLogger(commonlogger.NewDefaultLogger(FlagLoggerName))

	// ListenerLogger is the logger for the listener
	ListenerLogger = commonlistener.NewLogger(commonlogger.NewDefaultLogger(ListenerLoggerName))

	// EnvironmentLogger is the logger for the environment
	EnvironmentLogger = commonenv.NewLogger(commonlogger.NewDefaultLogger(EnvironmentLoggerName))

	// UserServiceLogger is the logger for the user service
	UserServiceLogger = user.NewLogger(commonlogger.NewDefaultLogger(UserServiceLoggerName))
)
