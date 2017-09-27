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

func TestLoggerFalseDebug(t *testing.T) {
	lgr, bo, be := testSetupLogger(false)
	want := "message"
	lgr.Debug(want)
	if got := bo.String(); got != "" {
		t.Errorf("stdout: got %q, want empty", got)
	}
	if got := be.String(); got != "" {
		t.Errorf("stderr: got %q, want empty", got)
	}
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

func TestLoggerDebugf(t *testing.T) {
	lgr, bo, be := testSetupLogger(true)
	want := "message 1"
	lgr.Debugf(want, 1, 2, 3)
	got := bo.String()
	if !strings.HasPrefix(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
	if be.String() != "" {
		t.Error("empty stderr")
	}
}
