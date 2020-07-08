// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package presentation

import (
	"fmt"
	"os"
)

// OpenTemplate opens a template file.
func OpenTemplate(fn string) (*Presentation, error) {
	p, err := Open(fn)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Open opens and reads a document from a file (.pptx).
func Open(filename string) (*Presentation, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening %s: %s", filename, err)
	}
	defer f.Close()
	fi, err := os.Stat(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening %s: %s", filename, err)
	}
	_ = fi
	return Read(f, fi.Size())
}
