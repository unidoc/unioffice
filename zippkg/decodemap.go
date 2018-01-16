// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package zippkg

import (
	"archive/zip"
	"path"

	"baliance.com/gooxml/schema/soo/pkg/relationships"
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
type OnNewRelationshipFunc func(decMap *DecodeMap, target, typ string, files []*zip.File, rel *relationships.Relationship, src Target) error

// DecodeMap is used to walk a tree of relationships, decoding files and passing
// control back to the document.
type DecodeMap struct {
	pathsToIfcs map[string]Target
	basePaths   map[*relationships.Relationships]string
	rels        []Target
	decodeFunc  OnNewRelationshipFunc
	decoded     map[string]struct{}
	indices     map[string]int
}

// SetOnNewRelationshipFunc sets the function to be called when a new
// relationship has been discovered.
func (d *DecodeMap) SetOnNewRelationshipFunc(fn OnNewRelationshipFunc) {
	d.decodeFunc = fn
}

type Target struct {
	Path  string
	Typ   string
	Ifc   interface{}
	Index uint32
}

// AddTarget allows documents to register decode targets. Path is a path that
// will be found in the zip file and ifc is an XML element that the file will be
// unmarshaled to.  filePath is the absolute path to the target, ifc is the
// object to decode into, sourceFileType is the type of file that the reference
// was discovered in, and index is the index of the source file type.
func (d *DecodeMap) AddTarget(filePath string, ifc interface{}, sourceFileType string, idx uint32) bool {
	if d.pathsToIfcs == nil {
		d.pathsToIfcs = make(map[string]Target)
		d.basePaths = make(map[*relationships.Relationships]string)
		d.decoded = make(map[string]struct{})
		d.indices = make(map[string]int)
	}

	// we use path.Clean instead of filepath.Clean to ensure we
	// end up with forward separators
	fn := path.Clean(filePath)
	if _, ok := d.decoded[fn]; ok {
		// already decoded this file
		return false
	}
	d.decoded[fn] = struct{}{}
	d.pathsToIfcs[fn] = Target{Path: filePath, Typ: sourceFileType, Ifc: ifc, Index: idx}
	return true
}
func (d *DecodeMap) RecordIndex(path string, idx int) {
	d.indices[path] = idx
}
func (d *DecodeMap) IndexFor(path string) int {
	return d.indices[path]
}

// Decode loops decoding targets registered with AddTarget and calling th
func (d *DecodeMap) Decode(files []*zip.File) error {
	pass := 1
	for pass > 0 {

		// if we've loaded any relationships files, notify the document so it
		// can create elements to receive the decoded version
		for len(d.rels) > 0 {
			relSource := d.rels[len(d.rels)-1]
			d.rels = d.rels[0 : len(d.rels)-1]
			relRaw := relSource.Ifc.(*relationships.Relationships)
			for _, r := range relRaw.Relationship {
				bp, _ := d.basePaths[relRaw]
				d.decodeFunc(d, bp+r.TargetAttr, r.TypeAttr, files, r, relSource)
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
				if err := Decode(f, dest.Ifc); err != nil {
					return err
				}
				files[i] = nil

				// we decoded a relationships file, so we need to traverse it
				if drel, ok := dest.Ifc.(*relationships.Relationships); ok {
					d.rels = append(d.rels, dest)
					// find the path that any files mentioned in the
					// relationships file are relative to

					// we use path.Clean instead of filepath.Clean to ensure we
					// end up with forward separators
					basePath, _ := path.Split(path.Clean(f.Name + "/../"))
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
