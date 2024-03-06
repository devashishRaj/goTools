package battery_test

import (
	"os"
	"testing"

	"github.com/devashishRaj/goTools/battery"
	"github.com/google/go-cmp/cmp"
)

func TestParsePmsetOutput_GetsChargePercent(t *testing.T) {
	// First create pmset.txt
	// mkdir testdata
	// run : pmset -g ps >testdata/pmset.txt
	t.Parallel()
	data, err := os.ReadFile("testdata/pmset.txt")
	if err != nil {
		t.Fatal(err)
	}
	// assign value ChargingPercent from  pmset.txt
	want := battery.Status{
		ChargingPercent: 100,
	}
	got, err := battery.ParsePmsetOutput(string(data))
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}
