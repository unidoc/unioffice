// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package zippkg

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"time"

	"github.com/unidoc/unioffice"
)

// XMLHeader is a header that MarshalXML uses to prefix the XML files it creates.
const XMLHeader = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>` + "\n"

var nl = []byte{'\r', '\n'}

func MarshalXMLByTypeIndex(z *zip.Writer, dt unioffice.DocType, typ string, idx int, v interface{}) error {
	fn := unioffice.AbsoluteFilename(dt, typ, idx)
	return MarshalXML(z, fn, v)
}

func MarshalXMLByType(z *zip.Writer, dt unioffice.DocType, typ string, v interface{}) error {
	fn := unioffice.AbsoluteFilename(dt, typ, 0)
	return MarshalXML(z, fn, v)
}

// MarshalXML creates a file inside of a zip and marshals an object as xml, prefixing it
// with a standard XML header.
func MarshalXML(z *zip.Writer, filename string, v interface{}) error {
	fh := &zip.FileHeader{}
	fh.Method = zip.Deflate
	fh.Name = filename
	fh.SetModTime(time.Now())

	w, err := z.CreateHeader(fh)
	if err != nil {
		return fmt.Errorf("creating %s in zip: %s", filename, err)
	}
	_, err = w.Write([]byte(XMLHeader))
	if err != nil {
		return fmt.Errorf("creating xml header to %s: %s", filename, err)
	}
	if err = xml.NewEncoder(SelfClosingWriter{w}).Encode(v); err != nil {
		return fmt.Errorf("marshaling %s: %s", filename, err)
	}
	_, err = w.Write(nl)
	return err
}
