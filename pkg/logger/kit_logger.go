package logger

import (
	"os"
	"sync"

	kitlog "github.com/go-kit/kit/log"
)

var kitLogger kitlog.Logger
var once sync.Once

// GetKitLogger returns go-kit logger instance
func GetKitLogger() kitlog.Logger {
	once.Do(func() {
		kitLogger = kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
		kitLogger = kitlog.With(kitLogger, "ts", kitlog.DefaultTimestampUTC, "caller", kitlog.DefaultCaller)
	})
	return kitLogger
}
