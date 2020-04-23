package tempstorage

import "io"

type storage interface {
	Open(path string) (File, error)
	TempFile(dir, pattern string) (File, error)
	TempDir(pattern string) (string, error)
	RemoveAll(dir string) error
	Add(path string) error
}

type File interface {
	io.Reader
	io.Writer
	io.Closer
	Name() string
}

var s storage

func New(newStorage storage) {
	s = newStorage
}

func Open(path string) (File, error) {
	return s.Open(path)
}

func TempFile(dir, pattern string) (File, error) {
	return s.TempFile(dir, pattern)
}

func TempDir(pattern string) (string, error) {
	return s.TempDir(pattern)
}

func RemoveAll(dir string) error {
	return s.RemoveAll(dir)
}

func Add(path string) error {
	return s.Add(path)
}
