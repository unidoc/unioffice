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

type CT_TextboxInfo struct {
	IdAttr      *uint16
	TxbxContent *CT_TxbxContent
	ExtLst      *drawingml.CT_OfficeArtExtensionList
}

func NewCT_TextboxInfo() *CT_TextboxInfo {
	ret := &CT_TextboxInfo{}
	ret.TxbxContent = NewCT_TxbxContent()
	return ret
}
func (m *CT_TextboxInfo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	setxbxContent := xml.StartElement{Name: xml.Name{Local: "wp:txbxContent"}}
	e.EncodeElement(m.TxbxContent, setxbxContent)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "wp:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TextboxInfo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TxbxContent = NewCT_TxbxContent()
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 16)
			if err != nil {
				return err
			}
			pt := uint16(parsed)
			m.IdAttr = &pt
		}
	}
lCT_TextboxInfo:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "txbxContent":
				if err := d.DecodeElement(m.TxbxContent, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = drawingml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TextboxInfo
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_TextboxInfo) Validate() error {
	return m.ValidateWithPath("CT_TextboxInfo")
}
func (m *CT_TextboxInfo) ValidateWithPath(path string) error {
	if err := m.TxbxContent.ValidateWithPath(path + "/TxbxContent"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
