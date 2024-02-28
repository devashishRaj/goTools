package hello

import (
	"fmt"
	"io"
	"os"
)

// creating a struct so that each instance of this could have it's own required type of io.writer
type Printer struct {
	Output io.Writer
}

// a constructor which gives os.stdout as io.writer .
func NewPrinter() *Printer {
	return &Printer{Output: os.Stdout}
}

func (p *Printer) Print() {
	fmt.Fprintln(p.Output, "Hello, World!")
}

// instead of hello.NewPrinter().Print() we can just use a wrapper , here we name is Main as
// main funciton is to print ...
func Main() int {
	NewPrinter().Print()
	return 0
}

// io.writer is stanard library interface that means “thing you can write to”

func PrintTo(w io.Writer) {
	fmt.Fprintf(w, "Hello, World!\n")
}
