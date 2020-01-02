package shell

import "testing"

const (
	echoText1 = "next"
	echoText2 = "test"
)

func TestCommand_Exec(t *testing.T) {
	res, err := NewDefault().NewCommand("echo", echoText1).Exec(echoText2)
	if err != nil {
		t.Fatal(err)
	}

	echoRes := echoText1 + " " + echoText2
	if res != echoRes {
		t.Fatalf("unexpected result, expected: %s, got: %s", echoRes, res)
	}
}

func TestNewCommand(t *testing.T) {
	res, err := NewCommand("echo", echoText1).Exec(echoText2)
	if err != nil {
		t.Fatal(err)
	}

	echoRes := echoText1 + " " + echoText2
	if res != echoRes {
		t.Fatalf("unexpected result, expected: %s, got: %s", echoRes, res)
	}
}

// This is a way to save your command and use it later
func ExampleNewCommand() {
	// Creating a command from default shell
	cmd := NewCommand("echo")
	// ...
	res, err := cmd.Exec("text 1")
	// ...
	res, err = cmd.Exec("text 2")
	// ...
	_, _ = res, err
}

// This is a way to save your command and use it later
func ExampleShell_NewCommand() {
	// Creating specific shell
	sh := New("/bin/bash", []string{"-c"})

	// Creating a command from specific shell
	cmd := sh.NewCommand("echo")
	// ...
	res, err := cmd.Exec("text 1")
	// ...
	res, err = cmd.Exec("text 2")
	// ...
	_, _ = res, err
}
