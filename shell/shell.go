package shell

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/devashishRaj/goTools/ErrorHandle"
	"io"
	"os"
	"os/exec"
	"strings"
)

// Session prevents paperwork
type Session struct {
	Stdin          io.Reader
	Stdout, Stderr io.Writer
	// DryRun :commands aren’t run for real, just echoed to the output instead.

	DryRun bool
}

// NewSession : New... is used to create a object with particular configuration and hold the state of function
// We will name the object Session as a user starts a session with some configuration, interacts with it and exit
func NewSession(in io.Reader, out, errs io.Writer) *Session {
	return &Session{
		Stdin:  in,
		Stdout: out,
		Stderr: errs,
		DryRun: false,
	}
}

// Run :method on session struct to run the start and end the given session
func (s *Session) Run() {
	_, err := fmt.Fprintf(s.Stdout, "> ")
	ErrorHandle.Panic(err)

	input := bufio.NewScanner(s.Stdin)
	for input.Scan() {
		line := input.Text()
		cmd, err := CmdFromString(line)
		if err != nil {
			_, err = fmt.Fprintf(s.Stdout, "> ")
			ErrorHandle.Panic(err)
			//break
			// not using break as if user makes mistake no need to rerun whole program
			continue
		}
		if s.DryRun {
			// whole command with args is echoed into output not the output of command.
			_, err = fmt.Fprintf(s.Stdout, "%s\n> ", line)
			ErrorHandle.Panic(err)
			continue
		}
		output, err := cmd.CombinedOutput()
		if err != nil {
			_, err = fmt.Fprintf(s.Stderr, "Error:%v", err)
			ErrorHandle.Panic(err)
		}
		_, err = fmt.Fprintf(s.Stdout, "%s> ", output)
		ErrorHandle.Panic(err)
	}
	//type ctrl + d to close the input stream
	_, err = fmt.Fprintln(s.Stdout, "\nBe seeing you!")
	ErrorHandle.Panic(err)
}

func CmdFromString(cmdLine string) (*exec.Cmd, error) {
	// generate strings using " " as delimiter
	args := strings.Fields(cmdLine)
	if len(args) < 1 {
		return nil, errors.New("empty input")
	}
	//  use unroll operator "..." in-case no arguments is present in args[1:]
	// "args[1:]..." just unpacks(unroll) the slice to its individual elements
	// as exec.Command does not take slice it takes any number of strings
	return exec.Command(args[0], args[1:]...), nil
}

func Main() int {
	session := NewSession(os.Stdin, os.Stdout, os.Stderr)
	session.Run()
	return 0
}
