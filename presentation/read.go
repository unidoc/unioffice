// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentation

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"

	"baliance.com/gooxml"
	"baliance.com/gooxml/zippkg"
)

// Read reads a document from an io.Reader.
func Read(r io.ReaderAt, size int64) (*Presentation, error) {
	doc := newEmpty()

	td, err := ioutil.TempDir("", "gooxml-pptx")
	if err != nil {
		return nil, err
	}
	doc.TmpPath = td

	zr, err := zip.NewReader(r, size)
	if err != nil {
		return nil, fmt.Errorf("parsing zip: %s", err)
	}

	files := []*zip.File{}
	files = append(files, zr.File...)

	decMap := zippkg.DecodeMap{}
	decMap.SetOnNewRelationshipFunc(doc.onNewRelationship)
	// we should discover all contents by starting with these two files
	decMap.AddTarget(gooxml.ContentTypesFilename, doc.ContentTypes.X(), "", 0)
	decMap.AddTarget(gooxml.BaseRelsFilename, doc.Rels.X(), "", 0)
	if err := decMap.Decode(files); err != nil {
		return nil, err
	}

	for _, f := range files {
		if f == nil {
			continue
		}
		if err := doc.AddExtraFileFromZip(f); err != nil {
			return nil, err
		}
	}
	return doc, nil
}
