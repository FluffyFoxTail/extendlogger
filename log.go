package extendlogger

import (
	"time"
)

const (
	LogLevelInfo LogLevel = iota
	LogLevelWarning
	LogLevelError
)

type LogLevel int

type Log struct {
	Prefix  string
	Message string
	logLvl  LogLevel
	Time    string
}

func NewLogMsg(lvl LogLevel, prefix, msg string) *Log {
	time_stp := time.Now().Format("2006-01-02 15:04:05")

	return &Log{
		Prefix:  prefix,
		Message: msg,
		logLvl:  lvl,
		Time:    time_stp,
	}
}
