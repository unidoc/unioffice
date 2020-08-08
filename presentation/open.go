// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

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
