package gotest

import (
	"errors"
)

func Ok() error {
	return nil
}

func Fail() error {
	return errors.New("Fail")
}
