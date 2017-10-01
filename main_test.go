package main

import (
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	var (
		debug  = true
		stdout = os.Stdout // &bytes.Buffer{}
		stderr = os.Stderr // &bytes.Buffer{}
	)

	if got := run(testDataPath, debug, stdout, stderr); got != exitOK {
		t.Errorf("run(%q, ...) got %v, want %v", testDataPath, got, exitOK)
	}
}
