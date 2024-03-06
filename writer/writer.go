package writer

import "os"

func WriteToFile(path string, data []byte) error {

	err := os.WriteFile(path, data, 0o600)
	if err != nil {
		return err
	}
	// attacker might be able to read some or all data in timeperiod the
	// function writes data to file and changes the permission on.
	return os.Chmod(path, 00600)
}
