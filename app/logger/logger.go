package logger

import (
	authmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/middleware/auth"
	commongcloud "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/cloud/gcloud"
	commonenv "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/config/env"
	commonflag "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/config/flag"
	commonlistener "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/listener"
	commonlogger "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils/logger"
)

var (
	// FlagLogger is the logger for the flag
	FlagLogger = commonflag.NewLogger(commonlogger.NewDefaultLogger("Flag"))

	// ListenerLogger is the logger for the listener
	ListenerLogger = commonlistener.NewLogger(commonlogger.NewDefaultLogger("Net Listener"))

	// EnvironmentLogger is the logger for the environment
	EnvironmentLogger = commonenv.NewLogger(commonlogger.NewDefaultLogger("Environment"))

	// GCloudLogger is the logger for the Google Cloud
	GCloudLogger = commongcloud.NewLogger(commonlogger.NewDefaultLogger("Google Cloud"))

	// AuthMiddlewareLogger is the logger for the Auth middleware
	AuthMiddlewareLogger = authmiddleware.NewLogger(commonlogger.NewDefaultLogger("Auth Middleware"))
)
