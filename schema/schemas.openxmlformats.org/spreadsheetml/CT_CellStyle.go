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

type CT_CellStyle struct {
	// User Defined Cell Style
	NameAttr *string
	// Format Id
	XfIdAttr uint32
	// Built-In Style Id
	BuiltinIdAttr *uint32
	// Outline Style
	ILevelAttr *uint32
	// Hidden Style
	HiddenAttr *bool
	// Custom Built In
	CustomBuiltinAttr *bool
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_CellStyle() *CT_CellStyle {
	ret := &CT_CellStyle{}
	return ret
}

func (m *CT_CellStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xfId"},
		Value: fmt.Sprintf("%v", m.XfIdAttr)})
	if m.BuiltinIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "builtinId"},
			Value: fmt.Sprintf("%v", *m.BuiltinIdAttr)})
	}
	if m.ILevelAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "iLevel"},
			Value: fmt.Sprintf("%v", *m.ILevelAttr)})
	}
	if m.HiddenAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hidden"},
			Value: fmt.Sprintf("%v", *m.HiddenAttr)})
	}
	if m.CustomBuiltinAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "customBuiltin"},
			Value: fmt.Sprintf("%v", *m.CustomBuiltinAttr)})
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CellStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
		}
		if attr.Name.Local == "xfId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.XfIdAttr = uint32(parsed)
		}
		if attr.Name.Local == "builtinId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.BuiltinIdAttr = &pt
		}
		if attr.Name.Local == "iLevel" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ILevelAttr = &pt
		}
		if attr.Name.Local == "hidden" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HiddenAttr = &parsed
		}
		if attr.Name.Local == "customBuiltin" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CustomBuiltinAttr = &parsed
		}
	}
lCT_CellStyle:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_CellStyle %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CellStyle
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CellStyle and its children
func (m *CT_CellStyle) Validate() error {
	return m.ValidateWithPath("CT_CellStyle")
}

// ValidateWithPath validates the CT_CellStyle and its children, prefixing error messages with path
func (m *CT_CellStyle) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
