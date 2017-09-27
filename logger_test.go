package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewLogger(t *testing.T) {
	lgr := newLogger(true, nil, nil)

	if lgr.dl == nil {
		t.Error("nil debug")
	}
}

func testSetupLogger(debug bool) (*Logger, *bytes.Buffer, *bytes.Buffer) {
	var (
		bufOut = &bytes.Buffer{}
		bufErr = &bytes.Buffer{}
	)

	lgr := newLogger(debug, bufOut, bufErr)

	return lgr, bufOut, bufErr
}

func TestLoggerDebug(t *testing.T) {
	lgr, bo, be := testSetupLogger(true)
	want := "message"
	lgr.Debug(want)
	got := bo.String()
	if !strings.HasPrefix(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
	if be.String() != "" {
		t.Error("empty stderr")
	}
}
