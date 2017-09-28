package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	if got := run(testDataPath, nil, nil); got != exitOK {
		t.Errorf("run(%q, ...) got %v, want %v", testDataPath, got, exitOK)
	}
	const invalidPath = "invalid-path"
	if got := run(invalidPath, nil, nil); got != exitErr {
		t.Errorf("run(%q, ...) got %v, want %v", invalidPath, got, exitErr)
	}
}
