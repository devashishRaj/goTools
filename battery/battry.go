package battery

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

type Status struct {
	ChargingPercent int
}

// MustCompile only needs to be done once per program run thus at package level
// using var
var pmsetOutput = regexp.MustCompile("([0-9]+)%")

// GetStatus :Return status struct instance with chargingPercent filed set
func GetStatus() (Status, error) {
	text, err := GetPmsetOutput()
	if err != nil {
		return Status{}, err
	}
	return ParsePmsetOutput(text)
}

// GetPmsetOutput :run command "pmset -g ps"
func GetPmsetOutput() (string, error) {
	// first string is taken as command and rest like "-g" will be arguments
	// supplied to that command
	data, err := exec.Command("/usr/bin/pmset", "-g", "ps").CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ParsePmsetOutput :find's charging percentage from output of pmset and extracts chargingPercent
// via regex int Status struct
func ParsePmsetOutput(text string) (Status, error) {
	matches := pmsetOutput.FindStringSubmatch(text)
	if len(matches) < 2 {
		return Status{}, fmt.Errorf("failed to parse pmset output: %q", text)
	}
	// regex returns two values like ["98%", "98"] we need 2nd one
	charge, err := strconv.Atoi(matches[1])
	if err != nil {
		return Status{}, fmt.Errorf("failed to parse charge percentage: %q", matches[1])
	}
	return Status{
		ChargingPercent: charge,
	}, nil
}
