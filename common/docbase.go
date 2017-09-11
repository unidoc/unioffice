// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package common

import (
	"archive/zip"
	"fmt"
	"image"

	"baliance.com/gooxml/zippkg"
)

// DocBase is the type embedded in in the Document/Workbook/Presentation types
// that contains members common to all.
type DocBase struct {
	ContentTypes   ContentTypes
	AppProperties  AppProperties
	Rels           Relationships
	CoreProperties CoreProperties
	Thumbnail      image.Image // thumbnail preview of the document

	Images     []ImageRef
	ExtraFiles []ExtraFile
	TmpPath    string // path where temporary files are stored when opening documents

}

// AddExtraFileFromZip is used when reading an unsupported file from an OOXML
// file. This ensures that unsupported file content will at least round-trip
// correctly.
func (d *DocBase) AddExtraFileFromZip(f *zip.File) error {
	path, err := zippkg.ExtractToDiskTmp(f, d.TmpPath)
	if err != nil {
		return fmt.Errorf("error extracting unsupported file: %s", err)
	}
	d.ExtraFiles = append(d.ExtraFiles, ExtraFile{
		ZipPath:  f.Name,
		DiskPath: path,
	})
	return nil
}

// WriteExtraFiles writes the extra files to the zip package.
func (d *DocBase) WriteExtraFiles(z *zip.Writer) error {
	for _, ef := range d.ExtraFiles {
		if err := zippkg.AddFileFromDisk(z, ef.ZipPath, ef.DiskPath); err != nil {
			return err
		}
	}
	return nil
}

// ExtraFile is an unsupported file type extracted from, or to be written to a
// zip package
type ExtraFile struct {
	ZipPath  string
	DiskPath string
}
