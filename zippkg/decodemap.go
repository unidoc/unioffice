// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package zippkg

import (
	"archive/zip"
	"path/filepath"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/package/2006/relationships"
)

// OnNewRelationshipFunc is called when a new relationship has been discovered.
//
// target is a resolved path that takes into account the location of the
// relationships file source and should be the path in the zip file.
//
// files are passed so non-XML files that can't be handled by AddTarget can be
// decoded directly (e.g. images)
//
// rel is the actual relationship so its target can be modified if the source
// target doesn't match where gooxml will write the file (e.g. read in
// 'xl/worksheets/MyWorksheet.xml' and we'll write out
// 'xl/worksheets/sheet1.xml')
type OnNewRelationshipFunc func(decMap *DecodeMap, target, typ string, files []*zip.File, rel *relationships.Relationship) error

// DecodeMap is used to walk a tree of relationships, decoding files and passing
// control back to the document.
type DecodeMap struct {
	pathsToIfcs map[string]interface{}
	basePaths   map[*relationships.Relationships]string
	rels        []*relationships.Relationships
	decFunc     OnNewRelationshipFunc
}

// SetOnNewRelationshipFunc sets the function to be called when a new
// relationship has been discovered.
func (d *DecodeMap) SetOnNewRelationshipFunc(fn OnNewRelationshipFunc) {
	d.decFunc = fn
}

// AddTarget allows documents to register decode targets. Path is a path that
// will be found in the zip file and ifc is an XML element that the file will be
// unmarshaled to.
func (d *DecodeMap) AddTarget(path string, ifc interface{}) {
	if d.pathsToIfcs == nil {
		d.pathsToIfcs = make(map[string]interface{})
		d.basePaths = make(map[*relationships.Relationships]string)
	}
	d.pathsToIfcs[filepath.Clean(path)] = ifc
}

// Decode loops decoding targets registered with AddTarget and calling th
func (d *DecodeMap) Decode(files []*zip.File) error {
	pass := 1
	for pass > 0 {

		// if we've loaded any relationships files, notify the document so it
		// can create elements to receive the decoded version
		for len(d.rels) > 0 {
			relFile := d.rels[len(d.rels)-1]
			d.rels = d.rels[0 : len(d.rels)-1]
			for _, r := range relFile.Relationship {
				bp, _ := d.basePaths[relFile]
				d.decFunc(d, bp+r.TargetAttr, r.TypeAttr, files, r)
			}
		}

		for i, f := range files {
			if f == nil {
				continue
			}
			// if there is a registered target for the file
			if dest, ok := d.pathsToIfcs[f.Name]; ok {
				delete(d.pathsToIfcs, f.Name)
				// decode to the target and mark the file as nil so we'll skip
				// it later
				if err := Decode(f, dest); err != nil {
					return err
				}
				files[i] = nil

				// we decoded a relationships file, so we need to traverse it
				if drel, ok := dest.(*relationships.Relationships); ok {
					d.rels = append(d.rels, drel)
					// find the path that any files mentioned in the
					// relationships file are relative to
					basePath, _ := filepath.Split(filepath.Clean(f.Name + "/../"))
					d.basePaths[drel] = basePath
					// ensure we make another decode pass
					pass++
				}
			}
		}
		pass--
	}
	return nil
}
