package extendlogger

type OutputWriter interface {
	Init()
	Write(log *Log)
}

type LogWriter struct {
	writers []OutputWriter
}

func (lw *LogWriter) Log(log *Log) {
	for _, ow := range lw.writers {
		ow.Write(log)
	}
}
