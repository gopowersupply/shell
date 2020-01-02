package shell

import (
	"errors"
	"fmt"
)

// ErrCommand - General text for all errors from this package
const ErrCommand = "command execution error"

// CommandError - Type for all errors from this package
type CommandError struct {
	inner error
}

// NewCommandError - Error generator
func NewCommandError(err error) CommandError {
	return CommandError{
		inner: err,
	}
}

func (err CommandError) Error() string {
	return fmt.Sprintf("%s: %s", ErrCommand, err.inner.Error())
}

// Unwrap - Returns inner error
func (err CommandError) Unwrap() error {
	return err.inner
}

// IsCommandError - Checks that error is error from this package
func IsCommandError(err error) bool {
	return errors.As(err, &CommandError{})
}
