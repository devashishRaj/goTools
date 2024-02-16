package stringmatcher_test

import (
	"bytes"
	linecount "paperwork/lineCount"
	"testing"
)

func TestInputFromArgs_setInputPath(t *testing.T) {
	t.Parallel()
	want := 3
	args := []string{"testdata/three_lines.txt"}
	c, err := linecount.NewCounter(linecount.WithInputFromArgs(args))
	if err != nil {
		t.Fatal(err)
	}
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestInputFromArgs_IgnoreEmptyArgs(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2\n3")
	c, err := linecount.NewCounter(linecount.WithInput(inputBuf),
		linecount.WithInputFromArgs([]string{}))
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
