package greet_test

import (
	"bytes"
	"testing"
)

func TestGreetCLi(t *testing.T) {
	// running tests concurrently can in detecting bugs releated to concurrency too
	t.Parallel()
	Greets := []string{"raj", "dev", "ashish"}
	for _, greet := range Greets {
		// to mimic user input from terminal
		input := bytes.NewBufferString(greet)
		// mimic output to terminal
		output := new(bytes.Buffer)
		// clear the buffer for next user name .
		output.Reset()
		//greet.Greet(input, output)
		greet.Greet(input, output)
		want := "What is your name?\nHello, " + greet + "." + "\n"
		got := output.String()
		if want != got {
			t.Errorf("Wanted \n%s \ngot \n%s", want, got)
		}
	}
}
