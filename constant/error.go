package constant

import (
	"fmt"
	"github.com/pkg/errors"
)

func Error(text string) error {
	return errors.New(text)
}

func Errorf(format string, a ...interface{}) error {
	return errors.New(fmt.Sprintf(format, a...))
}

var (
	ErrProfileFilePathNil = Error("profile file path is nil.")
	ErrProfileNodeIdNil   = Error("NodeId is nil.")
)
