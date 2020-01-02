package shell

import (
	"fmt"
	"testing"
)

var (
	shellPath = "sh"
	shellArgs = []string{"-c"}
)

func TestNew(t *testing.T) {
	sh := New(shellPath, shellArgs)

	if sh.path != shellPath {
		t.Errorf("unexpected shellPath, expected %s, got %s", shellPath, sh.path)
	}

	if sh.args[0] != shellArgs[0] {
		t.Errorf("unexpected arg, expected %s, got %s", sh.args[0], shellArgs[0])
	}
}

func TestNewDefault(t *testing.T) {
	sh := NewDefault()

	if sh.path != shellPath {
		t.Errorf("unexpected shellPath, expected %s, got %s", shellPath, sh.path)
	}

	if sh.args[0] != shellArgs[0] {
		t.Errorf("unexpected arg, expected %s, got %s", sh.args[0], shellArgs[0])
	}
}

func TestCmd(t *testing.T) {
	res, err := Cmd("echo", echoText1)
	if err != nil {
		t.Fatal(err)
	}

	if res != echoText1 {
		t.Fatalf("unexpected result, expected %s, got %s", echoText1, res)
	}
}

func TestShell_Cmd(t *testing.T) {
	res, err := NewDefault().Cmd("echo", echoText1)
	if err != nil {
		t.Fatal(err)
	}

	if res != echoText1 {
		t.Fatalf("unexpected result, expected %s, got %s", echoText1, res)
	}
}

// This is simplest way to execute command as default shell
func ExampleCmd() {
	res, err := Cmd("echo", "this", "is test")
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

// This is a way to set your own shell and execute command from it
func ExampleNew() {
	sh := New("/bin/bash", []string{"-c"})

	res, err := sh.Cmd("echo", "text")
	if err != nil {
		panic(err)
	}

	fmt.Print(res)
}
