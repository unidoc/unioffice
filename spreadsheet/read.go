// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package spreadsheet

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/zippkg"
)

// Read reads a workbook from an io.Reader(.xlsx).
func Read(r io.ReaderAt, size int64) (*Workbook, error) {
	wb := New()

	zr, err := zip.NewReader(r, size)
	if err != nil {
		return nil, fmt.Errorf("parsing zip: %s", err)
	}

	files := []*zip.File{}
	files = append(files, zr.File...)

	addCustom := false
	for _, f := range files {
		if f.FileHeader.Name == "docProps/custom.xml" {
			addCustom = true
			break
		}
	}
	if addCustom {
		wb.createCustomProperties()
	}

	decMap := zippkg.DecodeMap{}
	decMap.SetOnNewRelationshipFunc(wb.onNewRelationship)
	// we should discover all contents by starting with these two files
	decMap.AddTarget(unioffice.ContentTypesFilename, wb.ContentTypes.X(), "", 0)
	decMap.AddTarget(unioffice.BaseRelsFilename, wb.Rels.X(), "", 0)
	if err := decMap.Decode(files); err != nil {
		return nil, err
	}

	// etra files are things we don't handle yet, or files that happened to have
	// been in the zip before.  We just round-trip them.
	for _, f := range files {
		if f == nil {
			continue
		}
		if err := wb.AddExtraFileFromZip(f); err != nil {
			return nil, err
		}
	}

	if addCustom {
		customPropertiesExist := false
		for _, rel := range wb.Rels.X().Relationship {
			if rel.TargetAttr == "docProps/custom.xml" {
				customPropertiesExist = true
				break
			}
		}
		if !customPropertiesExist {
			wb.addCustomRelationships()
		}
	}

	return wb, nil
}

// Open opens and reads a workbook from a file (.xlsx).
func Open(filename string) (*Workbook, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening %s: %s", filename, err)
	}
	defer f.Close()
	fi, err := os.Stat(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening %s: %s", filename, err)
	}
	wb, err := Read(f, fi.Size())
	if err != nil {
		return nil, err
	}
	dir, _ := filepath.Abs(filepath.Dir(filename))
	wb.filename = filepath.Join(dir, filename)
	return wb, nil
}
