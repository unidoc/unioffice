// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package zippkg

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"baliance.com/gooxml/algo"
	"baliance.com/gooxml/schema/soo/pkg/relationships"
)

// RelationsPathFor returns the relations path for a given filename.
func RelationsPathFor(path string) string {
	sp := strings.Split(path, "/")
	pathPortion := strings.Join(sp[0:len(sp)-1], "/")
	filePortion := sp[len(sp)-1]
	pathPortion += "/_rels/"
	filePortion += ".rels"
	return pathPortion + filePortion
}

// Decode unmarshals the content of a *zip.File as XML to a given destination.
func Decode(f *zip.File, dest interface{}) error {
	rc, err := f.Open()
	if err != nil {
		return fmt.Errorf("error reading %s: %s", f.Name, err)
	}
	defer rc.Close()
	dec := xml.NewDecoder(rc)
	if err := dec.Decode(dest); err != nil {
		return fmt.Errorf("error decoding %s: %s", f.Name, err)
	}

	// this ensures that relationship ID is increasing, which we apparently rely
	// on....
	if ds, ok := dest.(*relationships.Relationships); ok {
		sort.Slice(ds.Relationship, func(i, j int) bool {
			lhs := ds.Relationship[i]
			rhs := ds.Relationship[j]
			return algo.NaturalLess(lhs.IdAttr, rhs.IdAttr)
		})
	}
	return nil
}

// AddFileFromDisk reads a file from disk and adds it at a given path to a zip file.
func AddFileFromDisk(z *zip.Writer, zipPath, diskPath string) error {
	w, err := z.Create(zipPath)
	if err != nil {
		return fmt.Errorf("error creating %s: %s", zipPath, err)
	}
	f, err := os.Open(diskPath)
	if err != nil {
		return fmt.Errorf("error opening %s: %s", diskPath, err)
	}
	_, err = io.Copy(w, f)
	return err
}

// ExtractToDiskTmp extracts a zip file to a temporary file in a given path,
// returning the name of the file.
func ExtractToDiskTmp(f *zip.File, path string) (string, error) {
	tmpFile, err := ioutil.TempFile(path, "zz")
	if err != nil {
		return "", err
	}
	defer tmpFile.Close()
	rc, err := f.Open()
	if err != nil {
		return "", err
	}
	defer rc.Close()
	_, err = io.Copy(tmpFile, rc)
	if err != nil {
		return "", err
	}
	return tmpFile.Name(), nil
}
