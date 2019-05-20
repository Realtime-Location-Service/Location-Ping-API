package logger

import (
	"github.com/rls/ping-api/pkg/config"
	"github.com/rls/ping-api/utils/consts"
)

// ILogger provides the log method
type ILogger interface {
	Log(keyvals ...interface{}) error
}

// New returns a logger of a specifc type
func New(loggerType consts.LoggerType) ILogger {
	if _, ok := consts.SupportedLogger[loggerType]; !ok || !config.AppCfg().Debug {
		return &silentlogger{}
	}
	if loggerType == consts.KitLogger {
		return GetKitLogger()
	}
	return nil
}

type silentlogger struct{}

// Log will log nothing
// will be used if defug is false
func (*silentlogger) Log(keyvals ...interface{}) error {
	return nil
}
