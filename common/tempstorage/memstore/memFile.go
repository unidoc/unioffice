// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package memstore

import "io"

// memFile implements tempstorage.File interface by applying its methods to memDataCell items
type memFile struct {
	file       *memDataCell
	readOffset int64
}

// Read reads from the underlying memDataCell in order to implement Reader interface
func (f *memFile) Read(p []byte) (int, error) {
	readOffset := f.readOffset
	size := f.file.size
	incomingLength := int64(len(p))
	if incomingLength > size {
		incomingLength = size
		p = p[:incomingLength]
	}
	if readOffset >= size {
		return 0, io.EOF
	}
	newOffset := readOffset + incomingLength
	if newOffset >= size {
		newOffset = size
	}
	n := copy(p, f.file.data[readOffset:newOffset])
	f.readOffset = newOffset
	return n, nil
}

// Write writes to the end of the underlying memDataCell in order to implement Writer interface
func (f *memFile) Write(p []byte) (int, error) {
	f.file.data = append(f.file.data, p...)
	f.file.size += int64(len(p))
	return len(p), nil
}

// Close is not applicable in this implementation
func (f *memFile) Close() error {
	return nil
}

// Name returns the filename of the underlying memDataCell
func (f *memFile) Name() string {
	return f.file.name
}
