package main

import (
	"testing"
)

const testRawDataPath = "testdata/testdata.txt"

func TestLoadRawData(t *testing.T) {
	const path = "testdata/testdata.txt"
	rd, err := loadRawData(testRawDataPath, new(Logger))
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if len(rd) == 0 {
		t.Error("empty data")
	}
}

func TestRun(t *testing.T) {
	if got := run(testRawDataPath, nil, nil); got != exitOK {
		t.Errorf("run(%q, ...) got %v, want %v", testRawDataPath, got, exitOK)
	}
	const invalidPath = "invalid-path"
	if got := run(invalidPath, nil, nil); got != exitErr {
		t.Errorf("run(%q, ...) got %v, want %v", invalidPath, got, exitErr)
	}
}
