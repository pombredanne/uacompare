package main

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	var (
		bo = &bytes.Buffer{}
		be = &bytes.Buffer{}
	)

	if got := run(testDataPath, false, bo, be); got != exitOK {
		t.Errorf("run(%q, ...) got %v, want %v", testDataPath, got, exitOK)
	}
}
