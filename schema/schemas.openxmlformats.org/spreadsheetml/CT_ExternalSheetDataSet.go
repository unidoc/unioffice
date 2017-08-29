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

type CT_ExternalSheetDataSet struct {
	// External Sheet Data Set
	SheetData []*CT_ExternalSheetData
}

func NewCT_ExternalSheetDataSet() *CT_ExternalSheetDataSet {
	ret := &CT_ExternalSheetDataSet{}
	return ret
}
func (m *CT_ExternalSheetDataSet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	sesheetData := xml.StartElement{Name: xml.Name{Local: "x:sheetData"}}
	e.EncodeElement(m.SheetData, sesheetData)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_ExternalSheetDataSet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ExternalSheetDataSet:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "sheetData":
				tmp := NewCT_ExternalSheetData()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SheetData = append(m.SheetData, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ExternalSheetDataSet
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_ExternalSheetDataSet) Validate() error {
	return m.ValidateWithPath("CT_ExternalSheetDataSet")
}
func (m *CT_ExternalSheetDataSet) ValidateWithPath(path string) error {
	for i, v := range m.SheetData {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/SheetData[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
