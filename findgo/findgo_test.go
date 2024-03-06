package findgo

import (
	"archive/zip"
	"os"
	"testing"
	"testing/fstest"

	// Don't import just "cmp" as it does not support Equal and Diff
	"github.com/google/go-cmp/cmp"
)

func TestFinfgoCorrectlyListsFilesInTree(t *testing.T) {
	//  The files, folders, and relationships between them,
	//  are collectively called a filesystem.
	t.Parallel()
	fsys := os.DirFS("testdata/tree")
	want := []string{
		"file.go",
		"subfolder/subfolder.go",
		"subfolder2/another.go",
		"subfolder2/file.go",
	}
	got := Files(fsys)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestFilesCorrectlyListsFilesInMapFS(t *testing.T) {
	// MapFS
	t.Parallel()
	fsys := fstest.MapFS{
		// filename.extension:{"you write some content of file "}
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}
	want := []string{
		"file.go",
		"subfolder/subfolder.go",
		"subfolder2/another.go",
		"subfolder2/file.go",
	}
	got := Files(fsys)
	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}

func TestFilesCorrectlyListFilesInZipArchive(t *testing.T) {
	t.Parallel()
	// commands to create zip , finder's compress action might give error
	// cd testdat
	// zip -r files.zip tree/
	fsys, err := zip.OpenReader("testdata/files.zip")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{
		"tree/file.go",
		"tree/subfolder/subfolder.go",
		"tree/subfolder2/another.go",
		"tree/subfolder2/file.go",
	}
	got := Files(fsys)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func BenchmarkFilesOnDisk(b *testing.B) {
	// time taken to test using files on main storage
	fsys := os.DirFS("testdata/tree")
	b.ResetTimer()
	for range b.N {
		_ = Files(fsys)
	}
}

func BenchmarkFilesOnMemrory(b *testing.B) {
	// time taken to test using filesystem created via MapFS on RAM
	fsys := fstest.MapFS{
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}
	b.ResetTimer()
	for range b.N {
		_ = Files(fsys)
	}
}
