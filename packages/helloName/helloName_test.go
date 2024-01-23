package helloname_test

import (
	"bytes"
	"testing"

	helloname "github.com/devashishRaj/goTools/helloName"
)

func TestIf_GivenNameIsGreeted(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	var names []string
	names = make([]string, 3)
	names[0] = "Raj"
	names[1] = "Dev"
	names[2] = "Shish"

	for _, name := range names {
		buf.Reset()
		helloname.Greet(buf, name)
		want := "Hello, " + name
		got := buf.String()
		if want != got {
			t.Errorf("Wanted %s , got %s", want, got)
		}
	}
}
