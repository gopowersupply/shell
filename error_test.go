package shell

import (
	"errors"
	"testing"
)

func TestIsCommandError(t *testing.T) {
	_, err := New("/unexpected/path", []string{}).Cmd()
	if !errors.As(err, &CommandError{}) {
		t.Fatalf("error shoul be as CommandError type")
	}
}
