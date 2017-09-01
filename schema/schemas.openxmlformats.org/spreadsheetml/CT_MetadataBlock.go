// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_MetadataBlock struct {
	// Metadata Record
	Rc []*CT_MetadataRecord
}

func NewCT_MetadataBlock() *CT_MetadataBlock {
	ret := &CT_MetadataBlock{}
	return ret
}

func (m *CT_MetadataBlock) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	serc := xml.StartElement{Name: xml.Name{Local: "x:rc"}}
	e.EncodeElement(m.Rc, serc)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_MetadataBlock) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_MetadataBlock:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "rc":
				tmp := NewCT_MetadataRecord()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rc = append(m.Rc, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_MetadataBlock
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_MetadataBlock and its children
func (m *CT_MetadataBlock) Validate() error {
	return m.ValidateWithPath("CT_MetadataBlock")
}

// ValidateWithPath validates the CT_MetadataBlock and its children, prefixing error messages with path
func (m *CT_MetadataBlock) ValidateWithPath(path string) error {
	for i, v := range m.Rc {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rc[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
