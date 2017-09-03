// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"fmt"
	"image"
	"os"
	// Add image format support
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"baliance.com/gooxml"
)

// Image is a container for image information.
type Image struct {
	Size   image.Point
	Format string
	Path   string
}

type iref struct {
	path string
}

// ImageRef is a reference to an image in a document.
type ImageRef struct {
	d   *Document
	ref *iref
	img Image
}

// RelID returns the relationship ID.
func (i ImageRef) RelID() string {
	for imgIdx, ir := range i.d.images {
		if ir == i.ref {
			imgID := i.d.docRels.FindRIDForN(imgIdx, gooxml.ImageType)
			return imgID
		}
	}
	return ""
}

// Format returns the format of the underlying image
func (i ImageRef) Format() string {
	return i.img.Format
}

// Path returns the path to an image file
func (i ImageRef) Path() string {
	return i.img.Path
}

// ImageFromFile reads an image from a file on disk. It doesn't keep the image
// in memory and only reads it to determine the format and size.  You can also
// construct an Image directly if the file and size are known.
func ImageFromFile(path string) (Image, error) {
	f, err := os.Open(path)
	r := Image{}
	if err != nil {
		return r, fmt.Errorf("error reading image: %s", err)
	}
	defer f.Close()
	imgDec, ifmt, err := image.Decode(f)
	if err != nil {
		return r, fmt.Errorf("unable to parse image: %s", err)
	}

	r.Path = path
	r.Format = ifmt
	r.Size = imgDec.Bounds().Size()
	return r, nil
}
