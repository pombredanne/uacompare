package main

import (
	"io"
	"log"
	"os"
)

// Logger holds all log levels.
type Logger struct {
	il *log.Logger
	dl *log.Logger
	el *log.Logger
}

func newLogger(debug bool, stdout, stderr io.Writer) *Logger {

	if stdout == nil {
		stdout = os.Stdout
	}

	if stderr == nil {
		stderr = os.Stderr
	}

	lgr := Logger{
		il: log.New(stdout, "", 0),
		el: log.New(stderr, "", 0),
	}

	if debug {
		lgr.dl = log.New(stdout, "", 0)
	}

	return &lgr
}

func printLog(lvl *log.Logger, format string, v ...interface{}) {
	if lvl == nil {
		return
	}
	if format == "" {
		lvl.Print(v...)
		return
	}
	lvl.Printf(format, v...)
}

// Debug prints output to debug log.
func (lgr *Logger) Debug(v ...interface{}) {
	printLog(lgr.dl, "", v...)
}

// Debugf prints formatted output to debug log.
func (lgr *Logger) Debugf(format string, v ...interface{}) {
	printLog(lgr.dl, format, v...)
}
