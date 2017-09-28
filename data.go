package main

import (
	"bufio"
	"fmt"
	"os"
)

// Data holds all user agents.
type Data []string

func loadData(lgr *Logger, path string, d *Data) error {

	r, err := os.Open(path)
	if err != nil {
		return err
	}

	defer func() {
		if err = r.Close(); err != nil {
			lgr.Errorf("load raw data close: %v", err)
		}
	}()

	s := bufio.NewScanner(r)

	var t string
	for s.Scan() {
		t = s.Text()
		if t == "" {
			continue
		}
		*d = append(*d, t)
	}
	if err = s.Err(); err != nil {
		return fmt.Errorf("scan: %v", s.Err())
	}
	return nil
}
