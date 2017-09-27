package main

import (
	"bufio"
	"fmt"
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

	raw, err := loadRawData(path, lgr)
	if err != nil {
		lgr.Errorf("load raw data: %v", err)
		return exitErr
	}

	_ = raw

	return exitOK
}

// RawData holds all user agents.
type RawData []string

func loadRawData(path string, lgr *Logger) (RawData, error) {
	empty := RawData{}

	r, err := os.Open(path)
	if err != nil {
		return empty, err
	}

	defer func() {
		if err = r.Close(); err != nil {
			lgr.Errorf("load raw data close: %v", err)
		}
	}()

	s := bufio.NewScanner(r)

	var rd RawData
	var t string
	for s.Scan() {
		t = s.Text()
		if t == "" {
			continue
		}
		rd = append(rd, t)
	}
	if err = s.Err(); err != nil {
		err = fmt.Errorf("scan: %v", s.Err())
		return empty, err
	}
	return rd, nil
}
