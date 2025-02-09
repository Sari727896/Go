package logger

import (
	"log"
	"os"
)

type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
}

func NewLogger() *Logger {
	flags := log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
	return &Logger{
		infoLogger:  log.New(os.Stdout, "INFO: ", flags),
		errorLogger: log.New(os.Stderr, "ERROR: ", flags),
		debugLogger: log.New(os.Stdout, "DEBUG: ", flags),
	}
}

func (l *Logger) Info(format string, v ...interface{}) {
	l.infoLogger.Printf(format, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.errorLogger.Printf(format, v...)
}

func (l *Logger) Debug(format string, v ...interface{}) {
	l.debugLogger.Printf(format, v...)
}
