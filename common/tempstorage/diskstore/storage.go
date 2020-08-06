// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

// Package diskstore implements tempStorage interface
// by using disk as a storage
package diskstore

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/unidoc/unioffice/common/tempstorage"
)

type diskStorage struct{}

// SetAsStorage sets temp storage as a disk storage
func SetAsStorage() {
	ds := diskStorage{}
	tempstorage.SetAsStorage(&ds)
}

// Open opens file from disk according to a path
func (d diskStorage) Open(path string) (tempstorage.File, error) {
	return os.Open(path)
}

// TempFile creates a new temp file by calling ioutil TempFile
func (d diskStorage) TempFile(dir, pattern string) (tempstorage.File, error) {
	return ioutil.TempFile(dir, pattern)
}

// TempFile creates a new temp directory by calling ioutil TempDir
func (d diskStorage) TempDir(pattern string) (string, error) {
	return ioutil.TempDir("", pattern)
}

// RemoveAll removes all files in the directory
func (d diskStorage) RemoveAll(dir string) error {
	if strings.HasPrefix(dir, os.TempDir()) {
		return os.RemoveAll(dir)
	}
	return nil
}

// Add is not applicable in the diskstore implementation
func (d diskStorage) Add(path string) error {
	return nil
}
