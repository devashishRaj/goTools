package greet

import (
	"bufio"
	"fmt"
	"github.com/devashishRaj/goTools/ErrorHandle"
	"io"
	"os"
)

func Greet(stdin io.Reader, stdout io.Writer) {
	name := "you"
	_, err := fmt.Fprintln(stdout, "What is your name?")
	ErrorHandle.Panic(err)
	// Creates a new Scanner to read input from stdin
	input := bufio.NewScanner(stdin)
	//  Reads a line of input until a newline character
	if input.Scan() {
		name = input.Text()
	}
	_, err = fmt.Fprintf(stdout, "Hello, %s.\n", name)
	ErrorHandle.Panic(err)
}

// Main :a wrapper with standard input and output stream
func Main() int {
	Greet(os.Stdin, os.Stdout)
	return 0
}
