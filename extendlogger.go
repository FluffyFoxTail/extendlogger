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
		Logger:   log.New(out, "", 0),
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

func (l *ExtendLogger) Infoln(msg string) {
	l.println(LogLevelInfo, "[*] INFO ", msg)
}

func (l *ExtendLogger) Warnln(msg string) {
	l.println(LogLevelWarning, "[*] WARNING ", msg)
}

func (l *ExtendLogger) Errorln(msg string) {
	l.println(LogLevelError, "[*] ERROR ", msg)
}
