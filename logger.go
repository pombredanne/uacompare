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

	l := Logger{
		il: log.New(stdout, "", 0),
		el: log.New(stderr, "", 0),
	}

	if debug {
		l.dl = log.New(stdout, "", 0)
	}

	return &l
}
