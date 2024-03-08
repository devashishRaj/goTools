package hello

import (
	"fmt"
	"github.com/devashishRaj/goTools/ErrorHandle"
	"io"
	"os"
)

// Printer : creating a struct so that each instance of this could have its own required type of io.writer
type Printer struct {
	Output io.Writer
}

// NewPrinter : a constructor which gives os.stdout as io.writer .
func NewPrinter() *Printer {
	return &Printer{Output: os.Stdout}
}

// Print :a method on Printer struct to print Hello World
func (p *Printer) Print() {
	fmt.Fprintln(p.Output, "Hello, World!")
}

// Main : instead of hello.NewPrinter().Print() we can just use a wrapper , here we name is Main as
// main function is to print ...
func Main() int {
	NewPrinter().Print()
	return 0
}

// PrintTo
// io.writer is standard library "interface" that means “thing you can write to”
// like bytes.buffer , os.stdin
func PrintTo(w io.Writer) {
	_, err := fmt.Fprintf(w, "Hello, World!\n")
	ErrorHandle.Panic(err)
}
