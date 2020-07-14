// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package tempstorage

import "io"

type storage interface {
	Open(path string) (File, error)
	TempFile(dir, pattern string) (File, error)
	TempDir(pattern string) (string, error)
	RemoveAll(dir string) error
	Add(path string) error
}

// File is a representation of a storage file
// with Read, Write, Close and Name methods identical to os.File.
type File interface {
	io.Reader
	io.Writer
	io.Closer
	Name() string
}

var s storage

// SetAsStorage changes temporary storage to newStorage.
func SetAsStorage(newStorage storage) {
	s = newStorage
}

// Open returns tempstorage File object by name.
func Open(path string) (File, error) {
	return s.Open(path)
}

// TempFile creates new empty file in the storage and returns it.
func TempFile(dir, pattern string) (File, error) {
	return s.TempFile(dir, pattern)
}

// TempDir creates a name for a new temp directory using a pattern argument.
func TempDir(pattern string) (string, error) {
	return s.TempDir(pattern)
}

// RemoveAll removes all files according to the dir argument prefix.
func RemoveAll(dir string) error {
	return s.RemoveAll(dir)
}

// Add reads a file from a disk and adds it to the storage.
func Add(path string) error {
	return s.Add(path)
}
