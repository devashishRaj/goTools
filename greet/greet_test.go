package greet_test

import (
	"bytes"
	"testing"

	"github.com/devashishRaj/goTools/greet"
)

func TestGreetCLi(t *testing.T) {
	// running tests concurrently can in detecting bugs releated to concurrency too
	t.Parallel()
	Greets := []string{"raj", "dev", "ashish"}
	for _, a_greet := range Greets {
		// to mimic user input from terminal
		input := bytes.NewBufferString(a_greet)
		// mimic output to terminal
		output := new(bytes.Buffer)
		// clear the buffer for next user name .
		output.Reset()
		greet.Greet(input, output)
		want := "What is your name?\nHello, " + a_greet + "." + "\n"
		got := output.String()
		if want != got {
			t.Errorf("Wanted \n%s \ngot \n%s", want, got)
		}
	}
}
