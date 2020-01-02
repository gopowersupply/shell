package shell

import "runtime"

// Shell provides access to shell
type Shell struct {
	path string
	args []string
}

// New creates a new shell provider
func New(path string, args []string) *Shell {
	return &Shell{
		path: path,
		args: args,
	}
}

// NewDefault creates a new shell from default provider
// For Windows:
//  cmd.exe /C
// For MacOS:
//  osascript -s h -e
// For other OS:
//  sh -c
func NewDefault() *Shell {
	switch runtime.GOOS {
	case "windows":
		return New("cmd.exe", []string{"/C"})
	case "darwin":
		return New("osascript", []string{"-s h", "-e"})
	default:
		return New("sh", []string{"-c"})
	}
}

// Cmd executes a command from shell
func (sh *Shell) Cmd(arguments ...string) (string, error) {
	return sh.NewCommand(arguments...).Exec()
}

// Cmd executes a command from default shell
func Cmd(arguments ...string) (string, error) {
	return NewCommand(arguments...).Exec()
}
