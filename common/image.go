// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package common

import (
	"fmt"
	"image"
	"os"

	"baliance.com/gooxml/measurement"
	// Add image format support
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

// Image is a container for image information. It's used as we need format and
// and size information to use images.
type Image struct {
	Size   image.Point
	Format string
	Path   string
}

// ImageRef is a reference to an image within a document.
type ImageRef struct {
	d     *DocBase
	rels  Relationships
	img   Image
	relID string
}

// MakeImageRef constructs an image reference which is a reference to a
// particular image file inside a document.  The same image can be used multiple
// times in a document by re-use the ImageRef.
func MakeImageRef(img Image, d *DocBase, rels Relationships) ImageRef {
	return ImageRef{img: img, d: d, rels: rels}
}

func (i *ImageRef) SetRelID(id string) {
	i.relID = id
}

// RelID returns the relationship ID.
func (i ImageRef) RelID() string {
	return i.relID
}

// Format returns the format of the underlying image
func (i ImageRef) Format() string {
	return i.img.Format
}

// Path returns the path to an image file
func (i ImageRef) Path() string {
	return i.img.Path
}

// Size returns the size of an image
func (i ImageRef) Size() image.Point {
	return i.img.Size
}

// RelativeHeight returns the relative height of an image given a fixed width.
// This is used when setting image to a fixed width to calculate the height
// required to keep the same image aspect ratio.
func (i ImageRef) RelativeHeight(w measurement.Distance) measurement.Distance {
	hScale := float64(i.Size().Y) / float64(i.Size().X)
	return w * measurement.Distance(hScale)
}

// RelativeWidth returns the relative width of an image given a fixed height.
// This is used when setting image to a fixed height to calculate the width
// required to keep the same image aspect ratio.
func (i ImageRef) RelativeWidth(h measurement.Distance) measurement.Distance {
	wScale := float64(i.Size().X) / float64(i.Size().Y)
	return h * measurement.Distance(wScale)
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
