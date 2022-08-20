package extendlogger

type ExtendLogger struct {
	Logger   *LogWriter
	logLevel LogLevel
}

func NewExtendLogger() *ExtendLogger {
	return &ExtendLogger{
		Logger:   logWriter,
		logLevel: LogLevelInfo,
	}
}

func (l *ExtendLogger) SetLogLevel(logLvl LogLevel) {
	if !logLvl.IsValid() {
		return
	}
	l.logLevel = logLvl
}

func (l *ExtendLogger) println(lvl LogLevel, prefix, msg string) {
	l.Logger.Log(NewLogMsg(lvl, prefix, msg))
}

func (l *ExtendLogger) Info(msg string) {
	l.println(LogLevelInfo, "[i] INFO", msg)
}

func (l *ExtendLogger) Warn(msg string) {
	l.println(LogLevelWarning, "[!] WARNING", msg)
}

func (l *ExtendLogger) Error(msg string) {
	l.println(LogLevelError, "[X] ERROR", msg)
}
