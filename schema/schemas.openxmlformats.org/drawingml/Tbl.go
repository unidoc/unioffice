// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"log"
)

type Tbl struct {
	CT_Table
}

func NewTbl() *Tbl {
	ret := &Tbl{}
	ret.CT_Table = *NewCT_Table()
	return ret
}
func (m *Tbl) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "a:tbl"
	return m.CT_Table.MarshalXML(e, start)
}
func (m *Tbl) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Table = *NewCT_Table()
lTbl:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tblPr":
				m.TblPr = NewCT_TableProperties()
				if err := d.DecodeElement(m.TblPr, &el); err != nil {
					return err
				}
			case "tblGrid":
				if err := d.DecodeElement(m.TblGrid, &el); err != nil {
					return err
				}
			case "tr":
				tmp := NewCT_TableRow()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Tr = append(m.Tr, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lTbl
		case xml.CharData:
		}
	}
	return nil
}
func (m *Tbl) Validate() error {
	return m.ValidateWithPath("Tbl")
}
func (m *Tbl) ValidateWithPath(path string) error {
	if err := m.CT_Table.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
