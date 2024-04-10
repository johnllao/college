package log

import (
	"io"
	"log"
	"os"
)

const (
	LevelError = 0
	LevelInfo  = 1
	LevelDebug = 2
)

var (
	debuglogger *log.Logger
	infologger  *log.Logger
	errorlogger *log.Logger

	level int
)

func Init(w io.Writer) {
	debuglogger = log.New(w, "DEBUG ", log.LstdFlags)
	infologger = log.New(w, "INFO ", log.LstdFlags)
	errorlogger = log.New(w, "ERROR ", log.LstdFlags)
}

func SetLevel(v int) {
	level = v
}

func Debug(m string, v ...any) {
	if debuglogger == nil {
		return
	}
	if level >= LevelDebug {
		debuglogger.Printf(m, v...)
	}
}

func Info(m string, v ...any) {
	if infologger == nil {
		return
	}
	if level >= LevelInfo {
		infologger.Printf(m, v...)
	}
}

func Error(m string, v ...any) {
	if errorlogger == nil {
		return
	}
	if level >= LevelError {
		errorlogger.Printf(m, v...)
	}
}

func Fatal(m string, v ...any) {
	Error(m, v...)
	os.Exit(1)
}
