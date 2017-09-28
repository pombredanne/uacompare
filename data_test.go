package main

import "testing"

const testDataPath = "testdata/testdata.txt"

func TestLoadData(t *testing.T) {
	var d Data
	err := loadData(new(Logger), testDataPath, &d)
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if len(d) == 0 {
		t.Error("empty data")
	}
}
