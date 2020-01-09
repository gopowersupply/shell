package shell

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

// Executes command and returns result or error
func command(shell string, args ...string) (string, error) {
	var bufError bytes.Buffer
	var bufOut bytes.Buffer

	c := exec.Command(shell, args...)
	c.Stderr = &bufError
	c.Stdout = &bufOut

	err := c.Run()
	if err != nil {
		return "", NewCommandError(err)
	}

	if bufError.Len() > 0 {
		return "", NewCommandError(errors.New(bufError.String()))
	}

	return strings.TrimSpace(bufOut.String()), nil
}

// ProcessCloser - closes a current terminal process
type ProcessCloser func()

// Opens out and error streams
func open(bufIn *bytes.Buffer, shell string, args ...string) (
	closer ProcessCloser, bufOut *bytes.Buffer, bufError *bytes.Buffer, err error) {
	bufOut = &bytes.Buffer{}
	bufError = &bytes.Buffer{}

	c := exec.Command(shell, args...)

	c.Stdin = bufIn
	c.Stderr = bufError
	c.Stdout = bufOut

	closer = func() {
		_ = c.Process.Release()
	}

	if err := c.Run(); err != nil {
		return closer, bufOut, bufError, NewCommandError(err)
	}

	return closer, bufOut, bufError, nil
}
