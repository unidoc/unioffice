// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"log"
)

type CT_PlotAreaChoice1 struct {
	ValAx  *CT_ValAx
	CatAx  *CT_CatAx
	DateAx *CT_DateAx
	SerAx  *CT_SerAx
}

func NewCT_PlotAreaChoice1() *CT_PlotAreaChoice1 {
	ret := &CT_PlotAreaChoice1{}
	return ret
}

func (m *CT_PlotAreaChoice1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAx != nil {
		sevalAx := xml.StartElement{Name: xml.Name{Local: "c:valAx"}}
		e.EncodeElement(m.ValAx, sevalAx)
	}
	if m.CatAx != nil {
		secatAx := xml.StartElement{Name: xml.Name{Local: "c:catAx"}}
		e.EncodeElement(m.CatAx, secatAx)
	}
	if m.DateAx != nil {
		sedateAx := xml.StartElement{Name: xml.Name{Local: "c:dateAx"}}
		e.EncodeElement(m.DateAx, sedateAx)
	}
	if m.SerAx != nil {
		seserAx := xml.StartElement{Name: xml.Name{Local: "c:serAx"}}
		e.EncodeElement(m.SerAx, seserAx)
	}
	return nil
}

func (m *CT_PlotAreaChoice1) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_PlotAreaChoice1:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "valAx":
				m.ValAx = NewCT_ValAx()
				if err := d.DecodeElement(m.ValAx, &el); err != nil {
					return err
				}
			case "catAx":
				m.CatAx = NewCT_CatAx()
				if err := d.DecodeElement(m.CatAx, &el); err != nil {
					return err
				}
			case "dateAx":
				m.DateAx = NewCT_DateAx()
				if err := d.DecodeElement(m.DateAx, &el); err != nil {
					return err
				}
			case "serAx":
				m.SerAx = NewCT_SerAx()
				if err := d.DecodeElement(m.SerAx, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_PlotAreaChoice1 %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PlotAreaChoice1
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PlotAreaChoice1 and its children
func (m *CT_PlotAreaChoice1) Validate() error {
	return m.ValidateWithPath("CT_PlotAreaChoice1")
}

// ValidateWithPath validates the CT_PlotAreaChoice1 and its children, prefixing error messages with path
func (m *CT_PlotAreaChoice1) ValidateWithPath(path string) error {
	if m.ValAx != nil {
		if err := m.ValAx.ValidateWithPath(path + "/ValAx"); err != nil {
			return err
		}
	}
	if m.CatAx != nil {
		if err := m.CatAx.ValidateWithPath(path + "/CatAx"); err != nil {
			return err
		}
	}
	if m.DateAx != nil {
		if err := m.DateAx.ValidateWithPath(path + "/DateAx"); err != nil {
			return err
		}
	}
	if m.SerAx != nil {
		if err := m.SerAx.ValidateWithPath(path + "/SerAx"); err != nil {
			return err
		}
	}
	return nil
}
