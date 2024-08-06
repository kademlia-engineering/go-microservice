/*
log.go
v0.1.0
1/2/24

This file initializes the web server
*/
package server

import "github.com/sirupsen/logrus"

// GetLogLevel
func GetLogLevel(level string) logrus.Level {
	switch level {
	case "ERROR":
		return logrus.ErrorLevel
	case "WARN":
		return logrus.WarnLevel
	case "INFO":
		return logrus.InfoLevel
	case "DEBUG":
		return logrus.DebugLevel
	case "TRACE":
		return logrus.TraceLevel
	default:
		return logrus.DebugLevel
	}
}
