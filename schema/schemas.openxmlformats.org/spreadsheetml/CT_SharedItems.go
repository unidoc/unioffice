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
	"time"

	"baliance.com/gooxml"
)

type CT_SharedItems struct {
	// Contains Semi Mixed Data Types
	ContainsSemiMixedTypesAttr *bool
	// Contains Non Date
	ContainsNonDateAttr *bool
	// Contains Date
	ContainsDateAttr *bool
	// Contains String
	ContainsStringAttr *bool
	// Contains Blank
	ContainsBlankAttr *bool
	// Contains Mixed Data Types
	ContainsMixedTypesAttr *bool
	// Contains Numbers
	ContainsNumberAttr *bool
	// Contains Integer
	ContainsIntegerAttr *bool
	// Minimum Numeric Value
	MinValueAttr *float64
	// Maximum Numeric Value
	MaxValueAttr *float64
	// Minimum Date Time
	MinDateAttr *time.Time
	// Maximum Date Time Value
	MaxDateAttr *time.Time
	// Shared Items Count
	CountAttr *uint32
	// Long Text
	LongTextAttr *bool
	// No Value
	M []*CT_Missing
	// Numeric
	N []*CT_Number
	// Boolean
	B []*CT_Boolean
	// Error Value
	E []*CT_Error
	// Character Value
	S []*CT_String
	// Date Time
	D []*CT_DateTime
}

func NewCT_SharedItems() *CT_SharedItems {
	ret := &CT_SharedItems{}
	ret.ContainsSemiMixedTypesAttr = gooxml.Bool(true)
	ret.ContainsNonDateAttr = gooxml.Bool(true)
	ret.ContainsDateAttr = gooxml.Bool(false)
	ret.ContainsStringAttr = gooxml.Bool(true)
	ret.ContainsBlankAttr = gooxml.Bool(false)
	ret.ContainsMixedTypesAttr = gooxml.Bool(false)
	ret.ContainsNumberAttr = gooxml.Bool(false)
	ret.ContainsIntegerAttr = gooxml.Bool(false)
	ret.LongTextAttr = gooxml.Bool(false)
	return ret
}

func (m *CT_SharedItems) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ContainsSemiMixedTypesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "containsSemiMixedTypes"},
			Value: fmt.Sprintf("%d", b2i(*m.ContainsSemiMixedTypesAttr))})
	}
	if m.ContainsNonDateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "containsNonDate"},
			Value: fmt.Sprintf("%d", b2i(*m.ContainsNonDateAttr))})
	}
	if m.ContainsDateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "containsDate"},
			Value: fmt.Sprintf("%d", b2i(*m.ContainsDateAttr))})
	}
	if m.ContainsStringAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "containsString"},
			Value: fmt.Sprintf("%d", b2i(*m.ContainsStringAttr))})
	}
	if m.ContainsBlankAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "containsBlank"},
			Value: fmt.Sprintf("%d", b2i(*m.ContainsBlankAttr))})
	}
	if m.ContainsMixedTypesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "containsMixedTypes"},
			Value: fmt.Sprintf("%d", b2i(*m.ContainsMixedTypesAttr))})
	}
	if m.ContainsNumberAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "containsNumber"},
			Value: fmt.Sprintf("%d", b2i(*m.ContainsNumberAttr))})
	}
	if m.ContainsIntegerAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "containsInteger"},
			Value: fmt.Sprintf("%d", b2i(*m.ContainsIntegerAttr))})
	}
	if m.MinValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minValue"},
			Value: fmt.Sprintf("%v", *m.MinValueAttr)})
	}
	if m.MaxValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "maxValue"},
			Value: fmt.Sprintf("%v", *m.MaxValueAttr)})
	}
	if m.MinDateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minDate"},
			Value: fmt.Sprintf("%v", *m.MinDateAttr)})
	}
	if m.MaxDateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "maxDate"},
			Value: fmt.Sprintf("%v", *m.MaxDateAttr)})
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	if m.LongTextAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "longText"},
			Value: fmt.Sprintf("%d", b2i(*m.LongTextAttr))})
	}
	e.EncodeToken(start)
	if m.M != nil {
		sem := xml.StartElement{Name: xml.Name{Local: "x:m"}}
		e.EncodeElement(m.M, sem)
	}
	if m.N != nil {
		sen := xml.StartElement{Name: xml.Name{Local: "x:n"}}
		e.EncodeElement(m.N, sen)
	}
	if m.B != nil {
		seb := xml.StartElement{Name: xml.Name{Local: "x:b"}}
		e.EncodeElement(m.B, seb)
	}
	if m.E != nil {
		see := xml.StartElement{Name: xml.Name{Local: "x:e"}}
		e.EncodeElement(m.E, see)
	}
	if m.S != nil {
		ses := xml.StartElement{Name: xml.Name{Local: "x:s"}}
		e.EncodeElement(m.S, ses)
	}
	if m.D != nil {
		sed := xml.StartElement{Name: xml.Name{Local: "x:d"}}
		e.EncodeElement(m.D, sed)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SharedItems) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ContainsSemiMixedTypesAttr = gooxml.Bool(true)
	m.ContainsNonDateAttr = gooxml.Bool(true)
	m.ContainsDateAttr = gooxml.Bool(false)
	m.ContainsStringAttr = gooxml.Bool(true)
	m.ContainsBlankAttr = gooxml.Bool(false)
	m.ContainsMixedTypesAttr = gooxml.Bool(false)
	m.ContainsNumberAttr = gooxml.Bool(false)
	m.ContainsIntegerAttr = gooxml.Bool(false)
	m.LongTextAttr = gooxml.Bool(false)
	for _, attr := range start.Attr {
		if attr.Name.Local == "containsSemiMixedTypes" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ContainsSemiMixedTypesAttr = &parsed
		}
		if attr.Name.Local == "containsNonDate" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ContainsNonDateAttr = &parsed
		}
		if attr.Name.Local == "containsDate" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ContainsDateAttr = &parsed
		}
		if attr.Name.Local == "containsString" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ContainsStringAttr = &parsed
		}
		if attr.Name.Local == "containsBlank" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ContainsBlankAttr = &parsed
		}
		if attr.Name.Local == "containsMixedTypes" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ContainsMixedTypesAttr = &parsed
		}
		if attr.Name.Local == "containsNumber" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ContainsNumberAttr = &parsed
		}
		if attr.Name.Local == "containsInteger" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ContainsIntegerAttr = &parsed
		}
		if attr.Name.Local == "minValue" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.MinValueAttr = &parsed
		}
		if attr.Name.Local == "maxValue" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.MaxValueAttr = &parsed
		}
		if attr.Name.Local == "minDate" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.MinDateAttr = &parsed
		}
		if attr.Name.Local == "maxDate" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.MaxDateAttr = &parsed
		}
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
		}
		if attr.Name.Local == "longText" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LongTextAttr = &parsed
		}
	}
lCT_SharedItems:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "m":
				tmp := NewCT_Missing()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.M = append(m.M, tmp)
			case "n":
				tmp := NewCT_Number()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.N = append(m.N, tmp)
			case "b":
				tmp := NewCT_Boolean()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.B = append(m.B, tmp)
			case "e":
				tmp := NewCT_Error()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.E = append(m.E, tmp)
			case "s":
				tmp := NewCT_String()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.S = append(m.S, tmp)
			case "d":
				tmp := NewCT_DateTime()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.D = append(m.D, tmp)
			default:
				log.Printf("skipping unsupported element on CT_SharedItems %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SharedItems
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SharedItems and its children
func (m *CT_SharedItems) Validate() error {
	return m.ValidateWithPath("CT_SharedItems")
}

// ValidateWithPath validates the CT_SharedItems and its children, prefixing error messages with path
func (m *CT_SharedItems) ValidateWithPath(path string) error {
	for i, v := range m.M {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/M[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.N {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/N[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.B {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/B[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.E {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/E[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.S {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/S[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.D {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/D[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
