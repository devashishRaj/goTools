package count

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type counter struct {
	input  io.Reader
	output io.Writer
}

// a function that takes the *counter itself as a parameter
// and sets the appropriate field to the value the user wants
type option func(*counter) error

// an option constructor : with these fields are now unexported and can only be set using valid inputs
// to constructor, like here only io.reader type is allowed not io.writer or anything else .
// you can use methods too but with drawback that user isn’t obliged to call these methods before
// using the object: they can call them at any time. You need to consider if user will need abiltiy
// to change the configuration while using it or not and is the object designed to have this ability
func WithInput(input io.Reader) option {
	return func(c *counter) error {
		if input == nil {
			return errors.New("Nil input reader")
		}
		c.input = input
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(c *counter) error {
		if output == nil {
			return errors.New("Nil output writer")
		}
		c.output = output
		return nil
	}
}

// []string comes handy here as os.Args is also a slice of strings , in test it becomes usefull
// as we can simply give path of test_data as a string.
func WithInputFromArgs(args []string) option {
	return func(c *counter) error {
		if len(args) < 1 {
			return nil
		}
		f, err := os.Open(args[0])
		if err != nil {
			return err
		}
		c.input = f
		return nil
	}
}

func NewCounter(opts ...option) (*counter, error) {
	// default values for fields are here. If a filename is not passed as argument stdin will be used
	// if both are present , file argument will be preferred
	c := &counter{
		input:  os.Stdin,
		output: os.Stdout,
	}
	// as opt is type of fucntion on counter strruct , it will be here ensure valid fields are set
	// whether they are default one defined above or by the ones passed in opt
	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *counter) Lines() int {
	lines := 0
	input := bufio.NewScanner(c.input)
	for input.Scan() {
		lines++
	}
	for _ ,f := range c.files
	return lines
}

// // wrapper
// func Main() int {
// 	// The first element of os.Args is always the pathname of the running binary
// 	c, err := NewCounter(WithInputFromArgs(os.Args[1:]))
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		//os.Exit(1) // instead of os.exit(1) we will send non-zero value to main that Main faced an
// 		// error:of some kind . person using the library might wanna not have a panic beviour and
// 		// might wanna prompt via main again for valid input ?
// 	}
// 	fmt.Println(c.Lines())
// 	return 0
// }

func MainLines() int {
	c, err := NewCounter(WithInputFromArgs(os.Args[1:]))
	if err != nil {
		fmt.Println(os.Stderr,err)
		return 1
	}
	fmt.Println(c.Lines())
	return 0

}

func MainWords() int {
    c, err := NewCounter(
        WithInputFromArgs(os.Args[1:]),
    )
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return 1
    }
    fmt.Println(c.Words())
    return 0
}

