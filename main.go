package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {}

const (
	exitOK   = 0
	exitErr  = 1
	exitUser = 2
)

// RawData holds all user agents.
type RawData []string

func loadRawData(path string) (RawData, error) {
	empty := RawData{}

	r, err := os.Open(path)
	if err != nil {
		return empty, err
	}

	defer func() {
		if err = r.Close(); err != nil {
			err = fmt.Errorf("close error: %v", err)
			log.Fatal(err)
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
