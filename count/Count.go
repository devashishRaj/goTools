package count

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/devashishRaj/goTools/ErrorHandle"
	"io"
	"os"
)

type Counter struct {
	files  []io.Reader
	input  io.Reader
	output io.Writer
}

// Option :a function that takes the *Counter itself as a parameter
// and sets the appropriate field to the value the user wants
type Option func(*Counter) error

// WithInput : an option constructor, with these fields are now unexported
// and can only be set using valid inputs to constructor,
// like here only io.reader type is allowed not io.writer or anything else .
// you can use methods too but with drawback that user isn’t obliged to call these methods before
// using the object: they can call them at any time. You need to consider if user will need ability
// to change the configuration while using it or not and is the object designed to have this ability
func WithInput(input io.Reader) Option {
	return func(c *Counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
	}
}

// WithOutput :set output stream
func WithOutput(output io.Writer) Option {
	return func(c *Counter) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		c.output = output
		return nil
	}
}

// WithInputFromArgs :[]string comes handy here as os.Args is also a slice of strings ,
// in test it become useful as we can simply give path of test_data as a string.
func WithInputFromArgs(args []string) Option {
	return func(c *Counter) error {
		if len(args) < 1 {
			//	if no arguments are provided
			//	WithInputFromArgs does not change the input field value
			//	and program reads input from console
			return nil
		}
		// slice of io.Reader with same length as number of arguments
		// the slice will hold references to opened files
		// references are essentially pointers to the File structs in memory.
		// structs contain metadata about the files, such as file descriptors and offsets.
		c.files = make([]io.Reader, len(args))
		for i, path := range args {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			// the files are stored as io.Reader in slice at appropriate index
			c.files[i] = f
		}
		// `Multireader` joins all files' references into a single stream
		c.input = io.MultiReader(c.files...)
		return nil
	}
}

// NewCounter :create a new Counter and apply any option(s) mentioned in parameter.
func NewCounter(opts ...Option) (*Counter, error) {
	// default values for fields are here. If a filename is not passed as argument stdin will be used
	// if both are present , file argument will be preferred
	c := &Counter{
		input:  os.Stdin,
		output: os.Stdout,
	}
	// as opt is type of function on Counter struct , it will be here ensure valid fields are set
	// whether they are default one defined above or by the ones passed in opt
	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *Counter) Lines() int {
	lines := 0
	input := bufio.NewScanner(c.input)
	for input.Scan() {
		lines++
	}
	for _, f := range c.files {
		err := f.(io.Closer).Close()
		ErrorHandle.Panic(err)
	}
	return lines
}

func (c *Counter) Words() int {
	words := 0
	input := bufio.NewScanner(c.input)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		words++
	}
	for _, f := range c.files {
		err := f.(io.Closer).Close()
		ErrorHandle.Panic(err)
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
// 		// error:of some kind . person using the library might wanna not have a panic behaviour and
// 		// might want prompt via main again for valid input ?
// 	}
// 	fmt.Println(c.Lines())
// 	return 0
// }

func MainLines() int {
	c, err := NewCounter(WithInputFromArgs(os.Args[1:]))
	if err != nil {
		_, err = fmt.Fprintln(os.Stderr, err)
		ErrorHandle.Panic(err)
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
		_, err = fmt.Fprintln(os.Stderr, err)
		ErrorHandle.Panic(err)
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
		_, err = fmt.Fprintln(os.Stderr, err)
		ErrorHandle.Panic(err)
		return 1
	}
	// flag.Parse needs to be able to modify each of our defined flag variables
	// ,so it must maintain a list of pointers to those variables
	if *linemode {
		fmt.Printf("Total number of lines %d\n", c.Lines())
	} else {
		fmt.Printf("Total number of words %d\n", c.Words())
	}
	return 0
}
