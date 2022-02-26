package logger

import (
	"flag"
	"fmt"
	"sync"

	zap "go.uber.org/zap"
)

var (
	doOnce sync.Once

	logLevel             = flag.String("log_level", "DEBUG", "set loggging level")
	logger   *zap.Logger = nil
)

func init() {
	doOnce.Do(func() {
		if *logLevel == "DEBUG" {
			logger, _ = zap.NewDevelopment()
		} else {
			logger, _ = zap.NewProduction()
		}
	})
}

func Debug(args ...interface{}) {
	msg := fmt.Sprint(args...)
	logger.Debug(msg)
}

func Info(args ...interface{}) {
	msg := fmt.Sprint(args...)
	logger.Info(msg)
}

func Error(args ...interface{}) {
	msg := fmt.Sprint(args...)
	logger.Error(msg)
}

func Debugf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	logger.Debug(msg)
}

func Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	logger.Info(msg)
}

func Errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	logger.Error(msg)
}
