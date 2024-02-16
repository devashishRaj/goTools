package greet

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Greet(stdin io.Reader, stdout io.Writer) {
	name := "you"
	fmt.Fprintln(stdout, "What is your name?")
	// Creates a new Scanner to read input from stdin
	input := bufio.NewScanner(stdin)
	//  Reads a line of input until a newline character
	if input.Scan() {
		name = input.Text()
	}
	fmt.Fprintf(stdout, "Hello, %s.\n", name)
}

func Main() int {
	Greet(os.Stdin, os.Stdout)
	return 0
}
