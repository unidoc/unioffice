// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

// Package memstore implements tempStorage interface
// by using memory as a storage
package memstore

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"sync"

	"github.com/unidoc/unioffice/common/tempstorage"
)

// memStorage contains and manages memdataCell items as operating system manages files
type memStorage struct {
	m sync.Map
}

// memDataCell is an imitation of file on disk
type memDataCell struct {
	name string
	data []byte
	size int64
}

// SetAsStorage sets temp storage as a memory storage
func SetAsStorage() {
	ts := memStorage{
		m: sync.Map{},
	}
	tempstorage.SetAsStorage(&ts)
}

// Open returns tempstorage File object by name
func (ts *memStorage) Open(path string) (tempstorage.File, error) {
	file, exists := ts.m.Load(path)
	if !exists {
		return nil, errors.New(fmt.Sprintf("Cannot open the file %s", path))
	}
	return &memFile{
		file: file.(*memDataCell),
	}, nil
}

// TempFile creates a new empty file in the storage and returns it
func (ts *memStorage) TempFile(dir, pattern string) (tempstorage.File, error) {
	path := dir + "/" + getNameFromPattern(pattern)
	newDataCell := &memDataCell{
		name: path,
		data: []byte{},
	}
	newFile := &memFile{
		file: newDataCell,
	}
	ts.m.Store(path, newDataCell)
	return newFile, nil
}

// TempDir creates a name for a new temp directory using a pattern argument
func (ts *memStorage) TempDir(pattern string) (string, error) {
	return getNameFromPattern(pattern), nil
}

// RemoveAll removes all files according to the dir argument prefix
func (ts *memStorage) RemoveAll(dir string) error {
	ts.m.Range(func(path, value interface{}) bool {
		ts.m.Delete(path)
		return true
	})
	return nil
}

// Add reads a file from a disk and adds it to the storage
func (ts *memStorage) Add(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	ts.m.Store(path, &memDataCell{name: path, data: data})
	return nil
}

func getNameFromPattern(pattern string) string {
	suffix, _ := random(6)
	return pattern + suffix
}

func random(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
