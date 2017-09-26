package main

import (
	"testing"
)

func TestLoadRawData(t *testing.T) {
	const path = "testdata/testdata.txt"
	rd, err := loadRawData(path)
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if len(rd) == 0 {
		t.Error("empty data")
	}
}
