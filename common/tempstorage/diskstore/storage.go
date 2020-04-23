// Package diskstore implements tempStorage interface
// by using disk as a storage
package diskstore

import (
	"io/ioutil"
	"os"
	"github.com/unidoc/unioffice/common/tempstorage"
)

type diskStorage struct {}

// SetAsStorage sets temp storage as a disk storage
func SetAsStorage() {
	ds := diskStorage{}
	tempstorage.New(&ds)
}

// Open opens file from disk according to a path
func (d diskStorage) Open(path string) (tempstorage.File, error) {
	return os.Open(path)
}

// TempFile creates temp file by calling ioutil TempFile
func (d diskStorage) TempFile(dir, pattern string) (tempstorage.File, error) {
	return ioutil.TempFile(dir, pattern)
}

// TempFile creates temp directory by calling ioutil TempDir
func (d diskStorage) TempDir(pattern string) (string, error) {
	return ioutil.TempDir("", pattern)
}

// RemoveAll removes all files in the directory
func (d diskStorage) RemoveAll(dir string) error {
	return os.RemoveAll(dir)
}

// Add is not applicable in diskstore implementation
func (d diskStorage) Add(path string) error {
	return nil
}
