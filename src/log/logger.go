package log

import (
	"sync"

	zap "go.uber.org/zap"
)

var (
	doOnce sync.Once
	logger *zap.Logger = nil
)

func init() {
	doOnce.Do(func() {
		logger, _ = zap.NewProduction()
	})
}

func Info(args ...interface{}) {
}
