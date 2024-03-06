//go:build integration

package battery_test

// A build tag is a special comment in a Go file that prevents it from being
//compiled unless that tag is defined. like //go:build darwin
// to run above you need : go test -tags=integration
/*
	why do need to isolate integration test?
	We don’t need to run integration tests so often, though,
	because the only way they could break is if something external changed:
	the pmset command was updated or removed, for example.
	Whearas unitTests are lightweight and if written properly they will only
	fail if there's something wrong with out code not some external reason.
*/

/*
	Purpose of integration test is to test what happens when we do use
	external dependencies.

	* It validates our assumptions about how they work .

*/
import (
	"bytes"
	"os/exec"
	"testing"

	"github.com/devashishRaj/goTools/battery"
)

func TestGetPmsetOutput_CapturesCmdOutput(t *testing.T) {
	t.Parallel()
	// if we cannot run the pmset command, maybe because we are not on
	// mac we will skip the test. Though
	data, err := exec.Command("usr/bin/pmset", "-g", "ps").CombinedOutput()
	if err != nil {
		t.Skipf("Unable to run 'pmset' command: %v", err)
	}
	// running on mac deivce without batteries like mac mini
	if !bytes.Contains(data, []byte("InternalBattery")) {
		t.Skip("no battery fitted")
	}

	text, err := battery.GetPmsetOutput()
	if err != nil {
		t.Fatal(err)
	}
	// we have unit test to check the GetPmsetOutput and ParsePmsetOutput
	// this is just check pmset is returning something parseable
	status, err := battery.ParsePmsetOutput(text)
	if err != nil {
		t.Fatal(err)
	}
	// just to what value we got from parsing
	// The output from this won’t be printed unless the test fails
	// or we run "go test -v"
	t.Logf("Charge: %d%%", status.ChargingPercent)
}
