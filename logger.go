package main

import "log"

// Logger holds all log levels.
type Logger struct {
	il *log.Logger
	dl *log.Logger
	el *log.Logger
}
