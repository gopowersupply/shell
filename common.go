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
