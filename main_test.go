package main

import (
	"testing"
)

func TestText(t *testing.T) {
	want := "Hello World!"
	msg := Text()
	if msg != want {
		t.Errorf(`Want: %v --> return value: %v`, want, msg)
	}
}
