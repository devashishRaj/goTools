package writer

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWriteToFile_WritesGivenDataToFile(t *testing.T) {
	t.Parallel()
	// use t.TempDir() to use a temporary folder which is deleted at end of test, no paperwork
	path := t.TempDir() + "/write_test.txt"
	want := []byte{1, 2, 3}
	err := WriteToFile(path, want)
	if err != nil {
		t.Fatal(err)
	}
	// check if file exits or not
	stat, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	// check if file was created with right permissions
	// 0o is for octal , then firsr digit for permissions to owner of file then
	//next digit is for group and next digit is for everyone else
	// 4 is read permission and 2 is write permission : 6
	perm := stat.Mode().Perm()
	if perm != 0o600 {
		t.Errorf("want file mode 0o600 got %v", perm)
	}
	// read contents of file
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	// comparision
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}

}

func TestWriteToFile_ReturnErrorForUnwritableFile(t *testing.T) {

	t.Parallel()
	path := "bogusDir/test.txt"
	want := []byte{1, 3, 1}
	err := WriteToFile(path, want)
	if err == nil {
		t.Fatal("want error when file not writable")
	}
}

func TestWriteToFile_ClobbersExistingFile(t *testing.T) {
	// if the file already exits , WriteToFile should overwrite it with new data
	//clobber it...
	t.Parallel()
	path := t.TempDir() + "/clobber_test.txt"
	err := os.WriteFile(path, []byte{4, 5, 6}, 0o600)
	if err != nil {
		t.Fatal(err)
	}
	want := []byte{1, 2, 3}
	err = WriteToFile(path, want)
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestWriteToFile_ChangesPermsOnExistingFile(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "test.txt"
	want := []byte{1, 2, 3}
	// pre-create an empty file with open perms
	err := os.WriteFile(path, []byte{}, 0o644)
	if err != nil {
		t.Fatal(err)
	}
	err = WriteToFile(path, want)
	if err != nil {
		t.Fatal(err)
	}
	stat, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	perms := stat.Mode().Perm()
	if perms != 0o600 {
		t.Errorf("Want file permission 0o600 got %v", perms)
	}

}
