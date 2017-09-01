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

type CT_FFCheckBox struct {
	Choice *CT_FFCheckBoxChoice
	// Default Checkbox Form Field State
	Default *CT_OnOff
	// Checkbox Form Field State
	Checked *CT_OnOff
}

func NewCT_FFCheckBox() *CT_FFCheckBox {
	ret := &CT_FFCheckBox{}
	return ret
}
func (m *CT_FFCheckBox) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Choice != nil {
		m.Choice.MarshalXML(e, start)
	}
	if m.Default != nil {
		sedefault := xml.StartElement{Name: xml.Name{Local: "w:default"}}
		e.EncodeElement(m.Default, sedefault)
	}
	if m.Checked != nil {
		sechecked := xml.StartElement{Name: xml.Name{Local: "w:checked"}}
		e.EncodeElement(m.Checked, sechecked)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_FFCheckBox) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_FFCheckBox:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "size":
				m.Choice = NewCT_FFCheckBoxChoice()
				if err := d.DecodeElement(&m.Choice.Size, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "sizeAuto":
				m.Choice = NewCT_FFCheckBoxChoice()
				if err := d.DecodeElement(&m.Choice.SizeAuto, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "default":
				m.Default = NewCT_OnOff()
				if err := d.DecodeElement(m.Default, &el); err != nil {
					return err
				}
			case "checked":
				m.Checked = NewCT_OnOff()
				if err := d.DecodeElement(m.Checked, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_FFCheckBox
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_FFCheckBox) Validate() error {
	return m.ValidateWithPath("CT_FFCheckBox")
}
func (m *CT_FFCheckBox) ValidateWithPath(path string) error {
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	if m.Default != nil {
		if err := m.Default.ValidateWithPath(path + "/Default"); err != nil {
			return err
		}
	}
	if m.Checked != nil {
		if err := m.Checked.ValidateWithPath(path + "/Checked"); err != nil {
			return err
		}
	}
	return nil
}
