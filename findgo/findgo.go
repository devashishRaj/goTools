package findgo

import (
	"github.com/devashishRaj/goTools/ErrorHandle"
	"io/fs"
	"path/filepath"
)

// any set of objects addressable by hierarchical pathname can be
// represented by a fs.FS
// refer to a fs.FS value as a “filesystem” from now on

// Files returns []string containing path of go files under a given folder
func Files(fsys fs.FS) (paths []string) {
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".go" {
			paths = append(paths, path)
		}
		return nil
	})
	ErrorHandle.Panic(err)
	return paths
}

// FilesCount Returns total no of go files in a folder
func FilesCount(fsys fs.FS) int {
	var count int
	err := fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if filepath.Ext(p) == ".go" {
			count++
		}
		return nil
	})
	ErrorHandle.Panic(err)
	//fmt.Println(count)
	return count
}
