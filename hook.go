package extendlogger

import "os"

var logWriter *LogWriter

func init() {
	logWriter = &LogWriter{
		writers: []OutputWriter{
			CreateNewWriter(os.Stderr),
		},
	}
}

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
