package safeWriter_test

import (
	"bytes"
	"errors"
	"github.com/devashishRaj/goTools/safeWriter"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestWriter_GivesExpectedOutput(t *testing.T) {
	t.Parallel()
	want := "hello"
	output := new(bytes.Buffer)
	sw := safeWriter.NewSafeWriter(output)
	sw.Write([]byte(want))
	got := output.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestWrite_DoesNothingIfErrorPresent(t *testing.T) {
	t.Parallel()
	want := "hello"
	output := new(bytes.Buffer)
	sw := safeWriter.NewSafeWriter(output)
	sw.WErr = errors.New("oh no")
	sw.Write([]byte(want))
	got := output.String()
	if got != "" {
		t.Errorf("Suppose to be empty but got %v", got)
	}
}
