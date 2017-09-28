package main

import (
	"io"
	"os"
)

func main() {
	os.Exit(run("", os.Stdout, os.Stderr))
}

const (
	exitOK   = 0
	exitErr  = 1
	exitUser = 2
)

func run(path string, stdout, stderr io.Writer) int {
	lgr := newLogger(true, stdout, stderr)

	var data Data
	err := loadData(lgr, path, &data)
	if err != nil {
		lgr.Errorf("load raw data: %v", err)
		return exitErr
	}

	_ = data

	return exitOK
}
