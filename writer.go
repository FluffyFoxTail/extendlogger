package extendlogger

import (
	"fmt"
	"os"
)

func (l LogLevel) IsValid() bool {
	switch l {
	case LogLevelInfo, LogLevelWarning, LogLevelError:
		return true
	default:
		return false
	}
}

func CreateNewWriter(dest *os.File) OutputWriter {
	writer := &BaseWriter{
		Target: dest,
	}
	return writer
}

type BaseWriter struct {
	Target *os.File
}

func (bw *BaseWriter) Init() {

}

func (bw *BaseWriter) Write(log *Log) {
	if log.logLvl.IsValid() {
		fmt.Fprintf(bw.Target, bw.FormatMsg(log))
	}
}

func (bw *BaseWriter) FormatMsg(log *Log) string {
	switch log.logLvl {
	case LogLevelInfo:
		log.Prefix = fmt.Sprintf("\x1b[32m%s\x1b[0m", log.Prefix)
	case LogLevelWarning:
		log.Prefix = fmt.Sprintf("\x1b[33m%s\x1b[0m", log.Prefix)
	case LogLevelError:
		log.Prefix = fmt.Sprintf("\x1b[31m%s\x1b[0m", log.Prefix)
	}

	if (log.Location == &EventLocation{}) {
		return fmt.Sprintf("%s %s, place of exec func not found ¯ \\ _ (ツ) _ / ¯\n[**] %s\n",

			log.Prefix,
			log.Time,
			log.Message,
		)
	}

	return fmt.Sprintf("%s %s on %s line: %d, func: %s\n[**] %s\n",
		log.Prefix,
		fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, log.Time),
		log.Location.File,
		log.Location.Line,
		fmt.Sprintf("\x1b[%dm%s\x1b[0m", 35, log.Location.PC),
		log.Message,
	)
}
