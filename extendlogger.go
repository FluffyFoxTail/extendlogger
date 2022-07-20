package extendlogger

import (
	"io"
	"log"
)

const (
	LogLevelInfo LogLevel = iota
	LogLevelWarning
	LogLevelError
)

type LogLevel int

func (l LogLevel) IsValid() bool {
	switch l {
	case LogLevelInfo, LogLevelWarning, LogLevelError:
		return true
	default:
		return false
	}
}

type ExtendLogger struct {
	*log.Logger
	logLevel LogLevel
}

func NewExtendLogger(out io.Writer) *ExtendLogger {
	return &ExtendLogger{
		Logger:   log.New(out, "", log.LstdFlags),
		logLevel: LogLevelError,
	}
}

func (l *ExtendLogger) SetLogLevel(logLvl LogLevel) {
	if !logLvl.IsValid() {
		return
	}
	l.logLevel = logLvl
}

func (l *ExtendLogger) println(srcLogLvl LogLevel, prefix, msg string) {
	if l.logLevel < srcLogLvl {
		return
	}

	l.Logger.Println(prefix + msg)
}
