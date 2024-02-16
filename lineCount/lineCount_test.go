package linecount_test

import (
	"bytes"
	linecount "paperwork/lineCount"
	"testing"
)

func TestLine_Count(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2\n3")
	c, err := linecount.NewCounter(linecount.WithInput(inputBuf))
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
