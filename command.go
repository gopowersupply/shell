package shell

import "strings"

// Command provides shell command
type Command struct {
	sh   *Shell
	args []string
}

// NewCommand - Creates a new command
func (sh *Shell) NewCommand(arguments ...string) *Command {
	return &Command{
		sh:   sh,
		args: arguments,
	}
}

// NewCommand - Creates a new command from default shell
func NewCommand(arguments ...string) *Command {
	return NewDefault().NewCommand(arguments...)
}

// Exec - Executes the command
func (cmd *Command) Exec(arguments ...string) (string, error) {
	cmdArgs := make([]string, 0, len(cmd.args)+len(arguments))
	cmdArgs = append(cmdArgs, cmd.args...)
	cmdArgs = append(cmdArgs, arguments...)

	args := make([]string, 0, len(cmd.sh.args)+1)
	args = append(args, cmd.sh.args...)
	args = append(args, strings.Join(cmdArgs, " "))

	return command(cmd.sh.path, args...)
}
