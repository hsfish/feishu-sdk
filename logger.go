package feishu_sdk

import (
	"fmt"

	"github.com/modern-go/reflect2"
)

type Logger interface {
	Info(format string, args ...interface{})
	Error(format string, args ...interface{})
}

func SetDefaultLogger(l Logger) {
	if !reflect2.IsNil(l) {
		logger = l
	}
}

type consoleLogger struct {
}

func (this *consoleLogger) Info(format string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(format, args...))
}

func (this *consoleLogger) Error(format string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(format, args...))
}

var logger Logger

func enablePrintln() bool {
	return logger != nil
}

func printInfo(format string, args ...interface{}) {
	if logger != nil {
		logger.Info(format, args...)
	}
}

func printError(format string, args ...interface{}) {
	if logger != nil {
		logger.Info(format, args...)
	}
}
