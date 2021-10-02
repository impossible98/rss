// import "rss.app/logger"
package logger

import (
	"fmt"
	"os"
	"time"
)

var requestedLevel = InfoLevel

type LogLevel uint32

const (
	FatalLevel LogLevel = iota

	ErrorLevel

	InfoLevel

	DebugLevel
)

func (level LogLevel) String() string {
	switch level {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

func EnableDebug() {
	requestedLevel = DebugLevel
	formatMessage(InfoLevel, "Debug mode enabled")
}

func Debug(format string, v ...interface{}) {
	if requestedLevel >= DebugLevel {
		formatMessage(DebugLevel, format, v...)
	}
}

func Info(format string, v ...interface{}) {
	if requestedLevel >= InfoLevel {
		formatMessage(InfoLevel, format, v...)
	}
}

func Error(format string, v ...interface{}) {
	if requestedLevel >= ErrorLevel {
		formatMessage(ErrorLevel, format, v...)
	}
}

func Fatal(format string, v ...interface{}) {
	if requestedLevel >= FatalLevel {
		formatMessage(FatalLevel, format, v...)
		os.Exit(1)
	}
}

func formatMessage(level LogLevel, format string, v ...interface{}) {
	var prefix = fmt.Sprintf("%s [%s] ", time.Now().Format("2006-01-02 15:04:05"), level)

	fmt.Fprintf(os.Stderr, prefix+format+"\n", v...)
}
