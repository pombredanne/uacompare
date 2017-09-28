package main

import (
	"flag"
	"io"
	"os"
)

func main() {
	d := flag.String("data", "", "path to data file")
	v := flag.Bool("v", false, "verbose output")
	os.Exit(run(*d, *v, os.Stdout, os.Stderr))
}

const (
	exitOK   = 0
	exitErr  = 1
	exitUser = 2
)

func run(path string, verbose bool, stdout, stderr io.Writer) int {
	lgr := newLogger(verbose, stdout, stderr)

	var data Data
	err := loadData(lgr, path, &data)
	if err != nil {
		lgr.Errorf("load raw data: %v", err)
		return exitErr
	}

	_ = data

	return exitOK
}
