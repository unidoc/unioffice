// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingDrawing

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_LinkedTextboxInformation struct {
	IdAttr  uint16
	SeqAttr uint16
	ExtLst  *drawingml.CT_OfficeArtExtensionList
}

func NewCT_LinkedTextboxInformation() *CT_LinkedTextboxInformation {
	ret := &CT_LinkedTextboxInformation{}
	return ret
}

func (m *CT_LinkedTextboxInformation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "seq"},
		Value: fmt.Sprintf("%v", m.SeqAttr)})
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "wp:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_LinkedTextboxInformation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 16)
			if err != nil {
				return err
			}
			m.IdAttr = uint16(parsed)
		}
		if attr.Name.Local == "seq" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 16)
			if err != nil {
				return err
			}
			m.SeqAttr = uint16(parsed)
		}
	}
lCT_LinkedTextboxInformation:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "extLst":
				m.ExtLst = drawingml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_LinkedTextboxInformation %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_LinkedTextboxInformation
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_LinkedTextboxInformation and its children
func (m *CT_LinkedTextboxInformation) Validate() error {
	return m.ValidateWithPath("CT_LinkedTextboxInformation")
}

// ValidateWithPath validates the CT_LinkedTextboxInformation and its children, prefixing error messages with path
func (m *CT_LinkedTextboxInformation) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
