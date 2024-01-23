package hello_test

import (
	"bytes"
	"testing"
	"github.com/devashishRaj/goTools/hello"
)

func TestPrintTo_PrintToGivenWriter(t *testing.T) {
	// "t" a pointer to testint.T struct which is used access methods to control the outcomes as "t"
	// contains state of the test during execution

	// tell compile that this test should be run concurrently with other test
	t.Parallel()

	buf := new(bytes.Buffer)
	hello.PrintTo(buf)
	want := "Hello, World!"
	got := buf.String()
	if want != got {
		t.Errorf("wanted %s got %s ", want, got)
	}
}
