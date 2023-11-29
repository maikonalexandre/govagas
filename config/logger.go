package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	info    *log.Logger
	debug   *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)

	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "[DEBUG]", logger.Flags()),
		info:    log.New(writer, "[INFO]", logger.Flags()),
		warning: log.New(writer, "[WARNING]", logger.Flags()),
		err:     log.New(writer, "[ERR]", logger.Flags()),
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Err(v ...interface{}) {
	l.err.Println(v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.warning.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Errf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}
