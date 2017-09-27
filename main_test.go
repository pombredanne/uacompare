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
