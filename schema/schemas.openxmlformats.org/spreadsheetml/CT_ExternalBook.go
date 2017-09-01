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

type CT_ExternalBook struct {
	IdAttr string
	// Supporting Workbook Sheet Names
	SheetNames *CT_ExternalSheetNames
	// Named Links
	DefinedNames *CT_ExternalDefinedNames
	// Cached Worksheet Data
	SheetDataSet *CT_ExternalSheetDataSet
}

func NewCT_ExternalBook() *CT_ExternalBook {
	ret := &CT_ExternalBook{}
	return ret
}
func (m *CT_ExternalBook) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	if m.SheetNames != nil {
		sesheetNames := xml.StartElement{Name: xml.Name{Local: "x:sheetNames"}}
		e.EncodeElement(m.SheetNames, sesheetNames)
	}
	if m.DefinedNames != nil {
		sedefinedNames := xml.StartElement{Name: xml.Name{Local: "x:definedNames"}}
		e.EncodeElement(m.DefinedNames, sedefinedNames)
	}
	if m.SheetDataSet != nil {
		sesheetDataSet := xml.StartElement{Name: xml.Name{Local: "x:sheetDataSet"}}
		e.EncodeElement(m.SheetDataSet, sesheetDataSet)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_ExternalBook) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
		}
	}
lCT_ExternalBook:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "sheetNames":
				m.SheetNames = NewCT_ExternalSheetNames()
				if err := d.DecodeElement(m.SheetNames, &el); err != nil {
					return err
				}
			case "definedNames":
				m.DefinedNames = NewCT_ExternalDefinedNames()
				if err := d.DecodeElement(m.DefinedNames, &el); err != nil {
					return err
				}
			case "sheetDataSet":
				m.SheetDataSet = NewCT_ExternalSheetDataSet()
				if err := d.DecodeElement(m.SheetDataSet, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ExternalBook
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_ExternalBook) Validate() error {
	return m.ValidateWithPath("CT_ExternalBook")
}
func (m *CT_ExternalBook) ValidateWithPath(path string) error {
	if m.SheetNames != nil {
		if err := m.SheetNames.ValidateWithPath(path + "/SheetNames"); err != nil {
			return err
		}
	}
	if m.DefinedNames != nil {
		if err := m.DefinedNames.ValidateWithPath(path + "/DefinedNames"); err != nil {
			return err
		}
	}
	if m.SheetDataSet != nil {
		if err := m.SheetDataSet.ValidateWithPath(path + "/SheetDataSet"); err != nil {
			return err
		}
	}
	return nil
}
