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

type CT_UpDownBars struct {
	GapWidth *CT_GapAmount
	UpBars   *CT_UpDownBar
	DownBars *CT_UpDownBar
	ExtLst   *CT_ExtensionList
}

func NewCT_UpDownBars() *CT_UpDownBars {
	ret := &CT_UpDownBars{}
	return ret
}
func (m *CT_UpDownBars) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.GapWidth != nil {
		segapWidth := xml.StartElement{Name: xml.Name{Local: "gapWidth"}}
		e.EncodeElement(m.GapWidth, segapWidth)
	}
	if m.UpBars != nil {
		seupBars := xml.StartElement{Name: xml.Name{Local: "upBars"}}
		e.EncodeElement(m.UpBars, seupBars)
	}
	if m.DownBars != nil {
		sedownBars := xml.StartElement{Name: xml.Name{Local: "downBars"}}
		e.EncodeElement(m.DownBars, sedownBars)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_UpDownBars) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_UpDownBars:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "gapWidth":
				m.GapWidth = NewCT_GapAmount()
				if err := d.DecodeElement(m.GapWidth, &el); err != nil {
					return err
				}
			case "upBars":
				m.UpBars = NewCT_UpDownBar()
				if err := d.DecodeElement(m.UpBars, &el); err != nil {
					return err
				}
			case "downBars":
				m.DownBars = NewCT_UpDownBar()
				if err := d.DecodeElement(m.DownBars, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
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
			break lCT_UpDownBars
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_UpDownBars) Validate() error {
	return m.ValidateWithPath("CT_UpDownBars")
}
func (m *CT_UpDownBars) ValidateWithPath(path string) error {
	if m.GapWidth != nil {
		if err := m.GapWidth.ValidateWithPath(path + "/GapWidth"); err != nil {
			return err
		}
	}
	if m.UpBars != nil {
		if err := m.UpBars.ValidateWithPath(path + "/UpBars"); err != nil {
			return err
		}
	}
	if m.DownBars != nil {
		if err := m.DownBars.ValidateWithPath(path + "/DownBars"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
