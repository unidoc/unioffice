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
	"strconv"
)

type CT_ExternalSheetData struct {
	// Sheet Id
	SheetIdAttr uint32
	// Last Refresh Resulted in Error
	RefreshErrorAttr *bool
	// Row
	Row []*CT_ExternalRow
}

func NewCT_ExternalSheetData() *CT_ExternalSheetData {
	ret := &CT_ExternalSheetData{}
	return ret
}
func (m *CT_ExternalSheetData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sheetId"},
		Value: fmt.Sprintf("%v", m.SheetIdAttr)})
	if m.RefreshErrorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "refreshError"},
			Value: fmt.Sprintf("%v", *m.RefreshErrorAttr)})
	}
	e.EncodeToken(start)
	if m.Row != nil {
		serow := xml.StartElement{Name: xml.Name{Local: "x:row"}}
		e.EncodeElement(m.Row, serow)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_ExternalSheetData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "sheetId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.SheetIdAttr = uint32(parsed)
		}
		if attr.Name.Local == "refreshError" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RefreshErrorAttr = &parsed
		}
	}
lCT_ExternalSheetData:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "row":
				tmp := NewCT_ExternalRow()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Row = append(m.Row, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ExternalSheetData
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_ExternalSheetData) Validate() error {
	return m.ValidateWithPath("CT_ExternalSheetData")
}
func (m *CT_ExternalSheetData) ValidateWithPath(path string) error {
	for i, v := range m.Row {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Row[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
