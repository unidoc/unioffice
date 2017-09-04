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

type CT_Number struct {
	// Value
	VAttr float64
	// Unused Item
	UAttr *bool
	// Calculated Item
	FAttr *bool
	// Caption
	CAttr *string
	// Member Property Count
	CpAttr *uint32
	// Format Index
	InAttr *uint32
	// Background Color
	BcAttr *string
	// Foreground Color
	FcAttr *string
	// Italic
	IAttr *bool
	// Underline
	UnAttr *bool
	// Strikethrough
	StAttr *bool
	// Bold
	BAttr *bool
	// OLAP Members
	Tpls []*CT_Tuples
	// Member Property Index
	X []*CT_X
}

func NewCT_Number() *CT_Number {
	ret := &CT_Number{}
	ret.IAttr = gooxml.Bool(false)
	ret.UnAttr = gooxml.Bool(false)
	ret.StAttr = gooxml.Bool(false)
	ret.BAttr = gooxml.Bool(false)
	return ret
}

func (m *CT_Number) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "v"},
		Value: fmt.Sprintf("%v", m.VAttr)})
	if m.UAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "u"},
			Value: fmt.Sprintf("%d", b2i(*m.UAttr))})
	}
	if m.FAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "f"},
			Value: fmt.Sprintf("%d", b2i(*m.FAttr))})
	}
	if m.CAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "c"},
			Value: fmt.Sprintf("%v", *m.CAttr)})
	}
	if m.CpAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cp"},
			Value: fmt.Sprintf("%v", *m.CpAttr)})
	}
	if m.InAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "in"},
			Value: fmt.Sprintf("%v", *m.InAttr)})
	}
	if m.BcAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "bc"},
			Value: fmt.Sprintf("%v", *m.BcAttr)})
	}
	if m.FcAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fc"},
			Value: fmt.Sprintf("%v", *m.FcAttr)})
	}
	if m.IAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "i"},
			Value: fmt.Sprintf("%d", b2i(*m.IAttr))})
	}
	if m.UnAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "un"},
			Value: fmt.Sprintf("%d", b2i(*m.UnAttr))})
	}
	if m.StAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "st"},
			Value: fmt.Sprintf("%d", b2i(*m.StAttr))})
	}
	if m.BAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "b"},
			Value: fmt.Sprintf("%d", b2i(*m.BAttr))})
	}
	e.EncodeToken(start)
	if m.Tpls != nil {
		setpls := xml.StartElement{Name: xml.Name{Local: "x:tpls"}}
		e.EncodeElement(m.Tpls, setpls)
	}
	if m.X != nil {
		sex := xml.StartElement{Name: xml.Name{Local: "x:x"}}
		e.EncodeElement(m.X, sex)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Number) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.IAttr = gooxml.Bool(false)
	m.UnAttr = gooxml.Bool(false)
	m.StAttr = gooxml.Bool(false)
	m.BAttr = gooxml.Bool(false)
	for _, attr := range start.Attr {
		if attr.Name.Local == "v" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.VAttr = parsed
		}
		if attr.Name.Local == "u" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UAttr = &parsed
		}
		if attr.Name.Local == "f" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FAttr = &parsed
		}
		if attr.Name.Local == "c" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CAttr = &parsed
		}
		if attr.Name.Local == "cp" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CpAttr = &pt
		}
		if attr.Name.Local == "in" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.InAttr = &pt
		}
		if attr.Name.Local == "bc" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BcAttr = &parsed
		}
		if attr.Name.Local == "fc" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FcAttr = &parsed
		}
		if attr.Name.Local == "i" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.IAttr = &parsed
		}
		if attr.Name.Local == "un" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UnAttr = &parsed
		}
		if attr.Name.Local == "st" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.StAttr = &parsed
		}
		if attr.Name.Local == "b" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.BAttr = &parsed
		}
	}
lCT_Number:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tpls":
				tmp := NewCT_Tuples()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Tpls = append(m.Tpls, tmp)
			case "x":
				tmp := NewCT_X()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.X = append(m.X, tmp)
			default:
				log.Printf("skipping unsupported element on CT_Number %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Number
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Number and its children
func (m *CT_Number) Validate() error {
	return m.ValidateWithPath("CT_Number")
}

// ValidateWithPath validates the CT_Number and its children, prefixing error messages with path
func (m *CT_Number) ValidateWithPath(path string) error {
	for i, v := range m.Tpls {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Tpls[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.X {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/X[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
