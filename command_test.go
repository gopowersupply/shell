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

func TestCommand_NewCommand(t *testing.T) {
	shell := NewDefault()
	cmdEcho := shell.NewCommand("echo")
	cmdInfoEcho := cmdEcho.NewCommand("[info]")
	_, err := cmdInfoEcho.Exec("test")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCommand_AttachShell(t *testing.T) {
	cmd := NewDefault().NewCommand()
	shell := New("sh", []string{"-c", "echo hi"})
	cmd = cmd.AttachShell(shell)
	res, err := shell.Cmd()
	if err != nil {
		t.Error(err)
	}
	if res != "hi" {
		t.Errorf("expected 'hi', got: %s", res)
	}
}

func TestCommand_Pipe(t *testing.T) {
	res, err := NewCommand("echo 'hi\nthere'").Pipe(NewCommand("grep hi")).Exec()
	if err != nil {
		t.Error(err)
	}
	if res != "hi" {
		t.Errorf("expected 'hi', got: %s", res)
	}
}

func TestCommand_And(t *testing.T) {
	cmdEcho := NewCommand("echo")
	cmdEchoHi := cmdEcho.NewCommand("hi")
	cmdEchoThere := cmdEcho.NewCommand("there")
	cmdHiThere := cmdEchoHi.And(cmdEchoThere)

	res, err := cmdHiThere.Exec()
	if err != nil {
		t.Fatal(err)
	}
	if res != "hi\nthere" {
		t.Errorf("expected 'hi\nthere', got: %s", res)
	}
}

func TestCommand_At(t *testing.T) {
	pkexec := NewCommand("pkexec")
	res, err := NewCommand("echo hi").At(pkexec).Exec()
	if err != nil {
		t.Fatal(err)
	}
	if res != "hi" {
		t.Errorf("expected 'hi', got: %s", res)
	}
}

func TestCommand_Run(t *testing.T) {
	NewCommand().Run()
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
