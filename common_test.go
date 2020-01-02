package shell

import (
	"strings"
	"testing"
)

func TestShell_Command(t *testing.T) {
	const (
		cmd = "echo"
		arg = "test"
	)

	args := make([]string, 0, len(shellArgs)+1)

	args = append(args, shellArgs...)
	args = append(args, cmd+" "+arg)

	t.Logf("args: %s", strings.Join(args, ", "))

	res, err := command(shellPath, args...)
	if err != nil {
		t.Fatal(err)
	}

	if res != arg {
		t.Fatalf("unexpected result, expected %s, got: %s", arg, res)
	}
}
