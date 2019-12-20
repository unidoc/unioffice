// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.
package wml

import (
	"encoding/xml"
	"fmt"
	"github.com/unidoc/unioffice"
)

type CT_Imagedata struct {
	//r:id
	IdAttr *string

	TitleAttr *string
}

func NewCT_Imagedata() *CT_Imagedata {
	ret := &CT_Imagedata{}
	return ret
}

func (m *CT_Imagedata) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}

	if m.TitleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "title"},
			Value: fmt.Sprintf("%v", *m.TitleAttr)})
	}

	_ = e.EncodeToken(start)
	_ = e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Imagedata) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}

		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "title" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TitleAttr = &parsed
			continue
		}
	}
lCT_Imagedata:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			default:
				unioffice.Log("skipping unsupported element on CT_Imagedata %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Imagedata
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_OleObject and its children
func (m *CT_Imagedata) Validate() error {
	return m.ValidateWithPath("CT_Imagedata")
}

// ValidateWithPath validates the CT_OleObject and its children, prefixing error messages with path
func (m *CT_Imagedata) ValidateWithPath(path string) error {
	return nil
}
