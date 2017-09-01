// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"log"
)

type EG_ScriptStyle struct {
	Scr *CT_Script
	Sty *CT_Style
}

func NewEG_ScriptStyle() *EG_ScriptStyle {
	ret := &EG_ScriptStyle{}
	return ret
}

func (m *EG_ScriptStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Name.Local = "m:EG_ScriptStyle"
	if m.Scr != nil {
		sescr := xml.StartElement{Name: xml.Name{Local: "m:scr"}}
		e.EncodeElement(m.Scr, sescr)
	}
	if m.Sty != nil {
		sesty := xml.StartElement{Name: xml.Name{Local: "m:sty"}}
		e.EncodeElement(m.Sty, sesty)
	}
	return nil
}

func (m *EG_ScriptStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_ScriptStyle:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "scr":
				m.Scr = NewCT_Script()
				if err := d.DecodeElement(m.Scr, &el); err != nil {
					return err
				}
			case "sty":
				m.Sty = NewCT_Style()
				if err := d.DecodeElement(m.Sty, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on EG_ScriptStyle %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_ScriptStyle
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_ScriptStyle and its children
func (m *EG_ScriptStyle) Validate() error {
	return m.ValidateWithPath("EG_ScriptStyle")
}

// ValidateWithPath validates the EG_ScriptStyle and its children, prefixing error messages with path
func (m *EG_ScriptStyle) ValidateWithPath(path string) error {
	if m.Scr != nil {
		if err := m.Scr.ValidateWithPath(path + "/Scr"); err != nil {
			return err
		}
	}
	if m.Sty != nil {
		if err := m.Sty.ValidateWithPath(path + "/Sty"); err != nil {
			return err
		}
	}
	return nil
}
