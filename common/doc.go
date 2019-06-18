// Package common contains wrapper types and utilities common to all of the
// OOXML document formats.
package common

import (
	"archive/zip"
	"strings"

	"github.com/mec07/unioffice/zippkg"
	"github.com/unidoc/unioffice"
)

//AddImageToZip adds an image (either from bytes or from disk) and adds it to the zip file.
func AddImageToZip(z *zip.Writer, img ImageRef, imageNum int, dt unioffice.DocType) error {
	fn := unioffice.AbsoluteFilename(dt, unioffice.ImageType, imageNum)
	fn = fn[0:len(fn)-3] + strings.ToLower(img.Format())
	if img.Data() != nil && len(*img.Data()) > 0 {
		if err := zippkg.AddFileFromBytes(z, fn, *img.Data()); err != nil {
			return err
		}
	} else if img.Path() != "" {
		if err := zippkg.AddFileFromDisk(z, fn, img.Path()); err != nil {
			return err
		}
	} else {
		unioffice.Log("unsupported image source: %+v", img)
	}
	return nil
}
