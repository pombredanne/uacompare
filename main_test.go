package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	if got := run(testDataPath, false, nil, nil); got != exitOK {
		t.Errorf("run(%q, ...) got %v, want %v", testDataPath, got, exitOK)
	}
}
