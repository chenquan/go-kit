package log

import (
	"sync/atomic"
)

var (
	log atomic.Value
)

type (
	Logger interface {
		Log(level Level, s string, fields ...Field) error
	}
	logEntry struct {
		log Logger
	}
)

func SetLogger(logger Logger) {
	log.Store(&logEntry{log: logger})
}

func getLogger() Logger {
	return log.Load().(*logEntry).log
}

func Log(level Level, s string, fields ...Field) error {
	return getLogger().Log(level, s, fields...)
}
