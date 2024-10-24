package main

import "testing"

func TestMarshal(t *testing.T) {
	_, err := jsonData()
	if err != nil {
		t.Error("Got error when parse data")
	}
}
