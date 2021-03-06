[![GitHub](https://img.shields.io/github/license/gopowersupply/shell)](https://github.com/gopowersupply/shell/blob/master/LICENSE)
[![Go version](https://img.shields.io/github/go-mod/go-version/gopowersupply/shell)](https://blog.golang.org/go1.13)
[![Build Status](https://travis-ci.org/gopowersupply/shell.svg?branch=master)](https://travis-ci.org/gopowersupply/shell)
[![Go Report Card](https://goreportcard.com/badge/gopowersupply/shell)](http://goreportcard.com/report/gopowersupply/shell)
[![Coverage Status](https://coveralls.io/repos/github/gopowersupply/shell/badge.svg?branch=master)](https://coveralls.io/github/gopowersupply/shell?branch=master)
[![code-coverage](http://gocover.io/_badge/github.com/gopowersupply/shell)](https://gocover.io/github.com/gopowersupply/shell)
[![GoDoc](https://godoc.org/github.com/gopowersupply/shell?status.svg)](https://godoc.org/github.com/gopowersupply/shell)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/gopowersupply/shell)](https://github.com/gopowersupply/shell/releases)
[![GitHub last commit](https://img.shields.io/github/last-commit/gopowersupply/shell)](https://github.com/gopowersupply/shell/commits/master)
[![GitHub issues](https://img.shields.io/github/issues/gopowersupply/shell)](https://github.com/gopowersupply/shell/issues)

The package to allow interaction with your system shell.

It should support **Windows**, **macOS** and all **Linux-based** distros.
Actually, was tested on **Ubuntu** only. I am open to your suggestions.

Get it from github:
```bash
go get -u https://github.com/gopowersupply/shell
```

Documentation can be [found here](https://godoc.org/github.com/gopowersupply/shell)

# Examples

This is a simplest way to execute a command from default shell:
```go
    res, err := shell.Cmd("echo hi")
    panic(err)
    // res: hi    
```
> :warning: Be aware that command result string truncates.
> This means that if the real output is `' sample output\n'` you will get `'sample output'`

You can make a command to use it later with additional params:
```go
    echo := shell.NewCommand("echo")
    // [...]
    res, _ := echo.Exec("text")
    // res: text
```

You can use a custom shell to execution:
```go
    res, _ := shell.New("/bin/sh", []string{"-c"}).Cmd("echo test")
    // res: test
```

Alternative to create a new command from custom shell:
```go
    sh := shell.New("/bin/sh", []string{"-c"})
    echo := sh.NewCommand("echo")
    res, _ := echo.Exec("hey", "there")
    // res: hey there
```

## Errors handling

This package has an own error type `CommandError`  
You can pass the package errors through your functions then detect it via `errors.As`:
```go
    func ExecUnexpected() error {
    	// [...] Here your other returns with own errors
        _, err := shell.Cmd("unexpected_command")
        if err != nil {
        	return err
        }
        // [...] Here your other returns with own errors
    }

    func main() {
    	err := ExecUnexpected()    	
    	if shell.IsCommandError(err) {
    		// [...] to do anything
    	} else {
    		// [...] to do something other    		
    	}
    }
```
And you can use `errors.As(err, &shell.CommandError{})` as alternative.