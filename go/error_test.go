package gotest

import (
	"testing"
)

func TestOk(t *testing.T) {
	err := Ok()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestFail(t *testing.T) {
	err := Fail()
	if err != nil {
		t.Fatal(err.Error())
	}
}
