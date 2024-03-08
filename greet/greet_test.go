package greet_test

import (
	"bytes"
	"testing"

	"github.com/devashishRaj/goTools/greet"
)

func TestGreetCLi(t *testing.T) {
	// running tests concurrently can in detecting bugs related to concurrency too
	t.Parallel()
	Greets := []string{"raj", "dev", "ashish"}
	// mimic output to terminal
	for _, aGreet := range Greets {
		// to mimic user input from terminal
		input := bytes.NewBufferString(aGreet)
		// acts as output stream
		output := new(bytes.Buffer)
		// if output is declared outside loop
		//you can clear the buffer for next username using Reset .
		//output.Reset()
		greet.Greet(input, output)
		//input.Reset()
		want := "What is your name?\nHello, " + aGreet + "." + "\n"
		got := output.String()
		if want != got {
			t.Errorf("Wanted \n%s \ngot \n%s", want, got)
		}
	}
}
