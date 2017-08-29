// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"log"
)

type CT_OdsoFieldMapData struct {
	// Merge Field Mapping
	Type *CT_MailMergeOdsoFMDFieldType
	// Data Source Name for Column
	Name *CT_String
	// Predefined Merge Field Name
	MappedName *CT_String
	// Index of Column Being Mapped
	Column *CT_DecimalNumber
	// Merge Field Name Language ID
	Lid *CT_Lang
	// Use Country-Based Address Field Ordering
	DynamicAddress *CT_OnOff
}

func NewCT_OdsoFieldMapData() *CT_OdsoFieldMapData {
	ret := &CT_OdsoFieldMapData{}
	return ret
}
func (m *CT_OdsoFieldMapData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Type != nil {
		setype := xml.StartElement{Name: xml.Name{Local: "w:type"}}
		e.EncodeElement(m.Type, setype)
	}
	if m.Name != nil {
		sename := xml.StartElement{Name: xml.Name{Local: "w:name"}}
		e.EncodeElement(m.Name, sename)
	}
	if m.MappedName != nil {
		semappedName := xml.StartElement{Name: xml.Name{Local: "w:mappedName"}}
		e.EncodeElement(m.MappedName, semappedName)
	}
	if m.Column != nil {
		secolumn := xml.StartElement{Name: xml.Name{Local: "w:column"}}
		e.EncodeElement(m.Column, secolumn)
	}
	if m.Lid != nil {
		selid := xml.StartElement{Name: xml.Name{Local: "w:lid"}}
		e.EncodeElement(m.Lid, selid)
	}
	if m.DynamicAddress != nil {
		sedynamicAddress := xml.StartElement{Name: xml.Name{Local: "w:dynamicAddress"}}
		e.EncodeElement(m.DynamicAddress, sedynamicAddress)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_OdsoFieldMapData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_OdsoFieldMapData:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "type":
				m.Type = NewCT_MailMergeOdsoFMDFieldType()
				if err := d.DecodeElement(m.Type, &el); err != nil {
					return err
				}
			case "name":
				m.Name = NewCT_String()
				if err := d.DecodeElement(m.Name, &el); err != nil {
					return err
				}
			case "mappedName":
				m.MappedName = NewCT_String()
				if err := d.DecodeElement(m.MappedName, &el); err != nil {
					return err
				}
			case "column":
				m.Column = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.Column, &el); err != nil {
					return err
				}
			case "lid":
				m.Lid = NewCT_Lang()
				if err := d.DecodeElement(m.Lid, &el); err != nil {
					return err
				}
			case "dynamicAddress":
				m.DynamicAddress = NewCT_OnOff()
				if err := d.DecodeElement(m.DynamicAddress, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_OdsoFieldMapData
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_OdsoFieldMapData) Validate() error {
	return m.ValidateWithPath("CT_OdsoFieldMapData")
}
func (m *CT_OdsoFieldMapData) ValidateWithPath(path string) error {
	if m.Type != nil {
		if err := m.Type.ValidateWithPath(path + "/Type"); err != nil {
			return err
		}
	}
	if m.Name != nil {
		if err := m.Name.ValidateWithPath(path + "/Name"); err != nil {
			return err
		}
	}
	if m.MappedName != nil {
		if err := m.MappedName.ValidateWithPath(path + "/MappedName"); err != nil {
			return err
		}
	}
	if m.Column != nil {
		if err := m.Column.ValidateWithPath(path + "/Column"); err != nil {
			return err
		}
	}
	if m.Lid != nil {
		if err := m.Lid.ValidateWithPath(path + "/Lid"); err != nil {
			return err
		}
	}
	if m.DynamicAddress != nil {
		if err := m.DynamicAddress.ValidateWithPath(path + "/DynamicAddress"); err != nil {
			return err
		}
	}
	return nil
}
