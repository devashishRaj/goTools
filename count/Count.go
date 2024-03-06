package count

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type counter struct {
	files  []io.Reader
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
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
	}
}

// set required output stream
func WithOutput(output io.Writer) option {
	return func(c *counter) error {
		if output == nil {
			return errors.New("nil output writer")
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
		c.files = make([]io.Reader, len(args))
		for i, path := range args {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			c.files[i] = f
		}
		c.input = io.MultiReader(c.files...)
		return nil
	}
}

// create a new counter and apply any option(s) mentioned in parmeter.
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
	for _, f := range c.files {
		f.(io.Closer).Close()
	}
	return lines
}

func (c *counter) Words() int {
	words := 0
	input := bufio.NewScanner(c.input)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		words++
	}
	for _, f := range c.files {
		f.(io.Closer).Close()
	}
	return words
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
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Printf("Total number of lines %d\n", c.Lines())
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
	fmt.Printf("Total number of words %d\n", c.Words())
	return 0
}

func Main() int {
	// flag.Bool to declare a new boolean flag, called lines
	linemode := flag.Bool("lines", false, "Count lines, not words")
	//   go build -o cmd/flag/flag cmd/flag/main.go
	// then run ./cmd/flag/flag -h

	flag.Usage = func() {
		fmt.Printf("Usage: %s [-lines] [files...]\n", os.Args[0])
		fmt.Println("Counts words (or lines) from stdin (or files).")
		fmt.Println("Flags:")
		flag.PrintDefaults()
	}
	// parse the program’s command‐line arguments, and use them to set the value of any flags
	// previously defined.
	// NOTE:that the flag package stops parsing as soon as it sees a non‐flag argument
	flag.Parse()
	// flag.Args,  to extract all the arguments that are left after flag.parse
	c, err := NewCounter(WithInputFromArgs(flag.Args()))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	// flag.Parse needs to be able to modify each of our defined flag variables
	// so it must maintain a list of pointers to those variables
	if *linemode {
		fmt.Printf("Total number of lines %d\n", c.Lines())
	} else {
		fmt.Printf("Total number of words %d\n", c.Words())
	}
	return 0
}
