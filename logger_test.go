package main

import (
	"bytes"
	"testing"
)

func TestNewLogger(t *testing.T) {
	l := newLogger(true, nil, nil)

	if l.dl == nil {
		t.Error("nil debug")
	}
}

func testSetup(debug bool) (*Logger, *bytes.Buffer, *bytes.Buffer) {
	var (
		bufOut = &bytes.Buffer{}
		bufErr = &bytes.Buffer{}
	)

	l := newLogger(debug, bufOut, bufErr)

	return l, bufOut, bufErr
}
