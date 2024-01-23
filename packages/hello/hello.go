package hello

import (
	"fmt"
	"io"
)

func PrintTo(w io.Writer) {
	fmt.Fprintf(w, "Hello, World!")
}
