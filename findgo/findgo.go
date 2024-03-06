package findgo

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

// any set of objects addressable by hierarchical pathnames can be
// represented by an fs.FS
// refer to an fs.FS value as a “filesystem” from now on

// returns []string containing path of go files under a given folder
func Files(fsys fs.FS) (paths []string) {
	fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".go" {
			paths = append(paths, path)
		}
		return nil
	})
	return paths
}

// Returns total no of go files in a folder
func FilesCount(fsys fs.FS) {
	var count int
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if filepath.Ext(p) == ".go" {
			count++
		}
		return nil
	})
	fmt.Println(count)
}
