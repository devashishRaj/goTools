package helloname

import (
	"fmt"
	"io"
	"log"
)

func Greet(w io.Writer, name string) {
	_, err := fmt.Fprintf(w, "Hello, %s", name)
	if err != nil {
		log.Fatal(err)
	}
}
