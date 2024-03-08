package hello_test

import (
	"bytes"
	"testing"

	"github.com/devashishRaj/goTools/hello"
)

func TestPrintTo_PrintsHelloMessageToGivenWriter(t *testing.T) {
	// "t" a pointer to testing.T struct which is used access methods to control the outcomes as "t"
	// contains state of the test during execution

	// tell compiler that this test should be run concurrently with other test
	t.Parallel()
	// bytes.Buffer. It’s an all‐purpose io.Writer that remembers what we write to it
	buf := new(bytes.Buffer)
	hello.PrintTo(buf)
	want := "Hello, World!"
	// buf.String, which returns all the text written to buf since it was created.
	got := buf.String()
	if want != got {
		t.Errorf("wanted %s got %s ", want, got)
	}
}

func TestPrintWrapperHello(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	p := hello.NewPrinter()
	// customize struct field value as necessary , dangerous though as there is no checking .
	// check count package for way of un-export struct fields and using options
	p.Output = buf
	p.Print()
	want := "Hello, World!\n"
	got := buf.String()
	if want != got {
		t.Errorf("want %q, \ngot %q", want, got)
	}

}
