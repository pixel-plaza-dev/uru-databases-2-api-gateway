package logger

import (
	commonenv "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/env"
	commonflag "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/flag"
	commonlistener "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/listener"
	commonlogger "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/logger"
)

var (
	// FlagLogger is the logger for the flag
	FlagLogger = commonflag.NewLogger(commonlogger.NewDefaultLogger("Flag"))

	// ListenerLogger is the logger for the listener
	ListenerLogger = commonlistener.NewLogger(commonlogger.NewDefaultLogger("Net Listener"))

	// EnvironmentLogger is the logger for the environment
	EnvironmentLogger = commonenv.NewLogger(commonlogger.NewDefaultLogger("Environment"))
)
