package shell

import (
	"bytes"
	"fmt"
	"log"
	"strings"
)

// Command provides shell command
type Command struct {
	sh   *Shell
	args []string
}

// buildArgs - Makes a args list from args from command and additional args
func (cmd *Command) buildArgs(arguments ...string) []string {
	cmdArgs := make([]string, 0, len(cmd.args)+len(arguments))
	cmdArgs = append(cmdArgs, cmd.args...)
	cmdArgs = append(cmdArgs, arguments...)

	args := make([]string, 0, len(cmd.sh.args)+1)
	args = append(args, cmd.sh.args...)
	args = append(args, strings.Join(cmdArgs, " "))

	return args
}

// NewCommand - Creates a new command
func (sh *Shell) NewCommand(arguments ...string) *Command {
	return &Command{
		sh:   sh,
		args: arguments,
	}
}

// NewCommand - Creates a new command from other command with additional arguments
func (cmd *Command) NewCommand(arguments ...string) *Command {
	return &Command{
		sh:   cmd.sh,
		args: append(cmd.args, arguments...),
	}
}

// NewCommand - Creates a new command from default shell
func NewCommand(arguments ...string) *Command {
	return NewDefault().NewCommand(arguments...)
}

// AttachShell - Makes a new command and replaces shell
func (cmd *Command) AttachShell(sh *Shell) *Command {
	return &Command{
		sh:   sh,
		args: cmd.args,
	}
}

// Pipe - Makes pipeline for 2 commands
// Example:
//  cmd := NewCommand("echo 'hi\nthere'").Pipe(NewCommand("grep hi"))
//  // cmd: echo hi\nthere | grep hi
func (cmd *Command) Pipe(command *Command) *Command {
	args := make([]string, len(cmd.args)+len(command.args)+1)
	args = append(args, cmd.args...)
	args = append(args, "|")
	args = append(args, command.args...)
	return &Command{
		sh:   cmd.sh,
		args: args,
	}
}

// And - Makes && for 2 commands.
// Example:
//  cmd := NewCommand("echo", "hi").And(NewCommand("echo", "there"))
//  // cmd: echo hi && echo there
func (cmd *Command) And(command *Command) *Command {
	args := make([]string, len(cmd.args)+len(command.args)+1)
	args = append(args, cmd.args...)
	args = append(args, "&&")
	args = append(args, command.args...)
	return &Command{
		sh:   cmd.sh,
		args: args,
	}
}

// At - Makes a new command from current and adds argument command before current command.
// Example:
//  cmd := NewCommand("echo", "hi").At("pkexec")
//  // cmd: pkexec echo hi
func (cmd *Command) At(command *Command) *Command {
	args := make([]string, len(cmd.args)+len(command.args))
	args = append(args, command.args...)
	args = append(args, cmd.args...)
	return &Command{
		sh:   cmd.sh,
		args: args,
	}
}

//TODO: Работает и открывает интерфейс
// Сделать класс типо CommandLine в который можно отправлять Command и получать результат
// Будет полезно для pkexec
func (cmd *Command) Run(arguments ...string) {
	bufIn := &bytes.Buffer{}
	bufIn.WriteString("echo hi\n")
	bufIn.WriteString("echo test\n")
	_, bufOut, _, err := open(bufIn, cmd.sh.path, "-c", "pkexec bash")
	if err != nil {
		log.Fatal(err)
	}

	b := make([]byte, 4)
	bufOut.Read(b)
	fmt.Println(b)
}

// Exec - Executes the command
func (cmd *Command) Exec(arguments ...string) (string, error) {
	return command(cmd.sh.path, cmd.buildArgs(arguments...)...)
}
