package extendlogger

import (
	"runtime"
	"time"
)

const (
	LogLevelInfo LogLevel = iota
	LogLevelWarning
	LogLevelError
)

type LogLevel int

type EventLocation struct {
	PC   string
	File string
	Line int
}

type Log struct {
	Prefix   string
	Message  string
	logLvl   LogLevel
	Location *EventLocation
	Time     string
}

func NewLogMsg(lvl LogLevel, prefix, msg string) *Log {
	pc, file, line, ok := runtime.Caller(3)
	time_stp := time.Now().Format("2006-01-02 15:04:05")

	if ok {
		return &Log{
			Prefix:  prefix,
			Message: msg,
			logLvl:  lvl,
			Location: &EventLocation{
				PC:   runtime.FuncForPC(pc).Name(),
				File: file,
				Line: line,
			},
			Time: time_stp,
		}
	}
	return &Log{
		Prefix:   prefix,
		Message:  msg,
		logLvl:   lvl,
		Location: &EventLocation{},
		Time:     time_stp,
	}
}
