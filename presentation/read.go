// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package presentation

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/zippkg"
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

	addCustom := false
	for _, f := range files {
		if f.FileHeader.Name == "docProps/custom.xml" {
			addCustom = true
			break
		}
	}
	if addCustom {
		doc.createCustomProperties()
	}

	decMap := zippkg.DecodeMap{}
	decMap.SetOnNewRelationshipFunc(doc.onNewRelationship)
	// we should discover all contents by starting with these two files
	decMap.AddTarget(unioffice.ContentTypesFilename, doc.ContentTypes.X(), "", 0)
	decMap.AddTarget(unioffice.BaseRelsFilename, doc.Rels.X(), "", 0)
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

	if addCustom {
		customPropertiesExist := false
		for _, rel := range doc.Rels.X().Relationship {
			if rel.TargetAttr == "docProps/custom.xml" {
				customPropertiesExist = true
				break
			}
		}
		if !customPropertiesExist {
			doc.addCustomRelationships()
		}
	}

	return doc, nil
}
