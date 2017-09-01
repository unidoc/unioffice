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

	"baliance.com/gooxml"
)

type CT_Cell struct {
	// Reference
	RAttr *string
	// Style Index
	SAttr *uint32
	// Cell Data Type
	TAttr ST_CellType
	// Cell Metadata Index
	CmAttr *uint32
	// Value Metadata Index
	VmAttr *uint32
	// Show Phonetic
	PhAttr *bool
	// Formula
	F *CT_CellFormula
	// Cell Value
	V *string
	// Rich Text Inline
	Is *CT_Rst
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_Cell() *CT_Cell {
	ret := &CT_Cell{}
	return ret
}

func (m *CT_Cell) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.RAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r"},
			Value: fmt.Sprintf("%v", *m.RAttr)})
	}
	if m.SAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "s"},
			Value: fmt.Sprintf("%v", *m.SAttr)})
	}
	if m.TAttr != ST_CellTypeUnset {
		attr, err := m.TAttr.MarshalXMLAttr(xml.Name{Local: "t"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.CmAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cm"},
			Value: fmt.Sprintf("%v", *m.CmAttr)})
	}
	if m.VmAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "vm"},
			Value: fmt.Sprintf("%v", *m.VmAttr)})
	}
	if m.PhAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ph"},
			Value: fmt.Sprintf("%v", *m.PhAttr)})
	}
	e.EncodeToken(start)
	if m.F != nil {
		sef := xml.StartElement{Name: xml.Name{Local: "x:f"}}
		e.EncodeElement(m.F, sef)
	}
	if m.V != nil {
		sev := xml.StartElement{Name: xml.Name{Local: "x:v"}}
		gooxml.AddPreserveSpaceAttr(&sev, *m.V)
		e.EncodeElement(m.V, sev)
	}
	if m.Is != nil {
		seis := xml.StartElement{Name: xml.Name{Local: "x:is"}}
		e.EncodeElement(m.Is, seis)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Cell) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "r" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RAttr = &parsed
		}
		if attr.Name.Local == "s" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.SAttr = &pt
		}
		if attr.Name.Local == "t" {
			m.TAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "cm" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CmAttr = &pt
		}
		if attr.Name.Local == "vm" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.VmAttr = &pt
		}
		if attr.Name.Local == "ph" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PhAttr = &parsed
		}
	}
lCT_Cell:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "f":
				m.F = NewCT_CellFormula()
				if err := d.DecodeElement(m.F, &el); err != nil {
					return err
				}
			case "v":
				m.V = new(string)
				if err := d.DecodeElement(m.V, &el); err != nil {
					return err
				}
			case "is":
				m.Is = NewCT_Rst()
				if err := d.DecodeElement(m.Is, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Cell %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Cell
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Cell and its children
func (m *CT_Cell) Validate() error {
	return m.ValidateWithPath("CT_Cell")
}

// ValidateWithPath validates the CT_Cell and its children, prefixing error messages with path
func (m *CT_Cell) ValidateWithPath(path string) error {
	if err := m.TAttr.ValidateWithPath(path + "/TAttr"); err != nil {
		return err
	}
	if m.F != nil {
		if err := m.F.ValidateWithPath(path + "/F"); err != nil {
			return err
		}
	}
	if m.Is != nil {
		if err := m.Is.ValidateWithPath(path + "/Is"); err != nil {
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
