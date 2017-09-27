package main

import (
	"bytes"
	"testing"
)

func TestNewLogger(t *testing.T) {
	lgr := newLogger(true, nil, nil)

	if lgr.dl == nil {
		t.Error("nil debug")
	}
}

func testSetup(debug bool) (*Logger, *bytes.Buffer, *bytes.Buffer) {
	var (
		bufOut = &bytes.Buffer{}
		bufErr = &bytes.Buffer{}
	)

	lgr := newLogger(debug, bufOut, bufErr)

	return lgr, bufOut, bufErr
}
