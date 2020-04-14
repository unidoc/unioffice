package storage

import "io"

type tempStorageItem struct {
	name       string
	size       int64
	bytes      []byte
	readOffset int64
}

// Read reads from tempStorageItem.bytes in order to implement Reader interface
func (tsi *tempStorageItem) Read(p []byte) (int, error) {
	readOffset := tsi.readOffset
	size := tsi.size
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
	n := copy(p, tsi.bytes[readOffset:newOffset])
	tsi.readOffset = newOffset
	return n, nil
}

// Write writes to the end of tempStorageItem.bytes in order to implement Writer interface
func (tsi *tempStorageItem) Write(p []byte) (int, error) {
	tsi.bytes = append(tsi.bytes, p...)
	return len(p), nil
}

// Name returns the filename of the storage item
func (tsi *tempStorageItem) Name() string {
	return tsi.name
}
