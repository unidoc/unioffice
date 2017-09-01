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

type CT_Font struct {
	// Font Name
	Name []*CT_FontName
	// Character Set
	Charset []*CT_IntProperty
	// Font Family
	Family []*CT_FontFamily
	// Bold
	B []*CT_BooleanProperty
	// Italic
	I []*CT_BooleanProperty
	// Strike Through
	Strike []*CT_BooleanProperty
	// Outline
	Outline []*CT_BooleanProperty
	// Shadow
	Shadow []*CT_BooleanProperty
	// Condense
	Condense []*CT_BooleanProperty
	// Extend
	Extend []*CT_BooleanProperty
	// Text Color
	Color []*CT_Color
	// Font Size
	Sz []*CT_FontSize
	// Underline
	U []*CT_UnderlineProperty
	// Text Vertical Alignment
	VertAlign []*CT_VerticalAlignFontProperty
	// Scheme
	Scheme []*CT_FontScheme
}

func NewCT_Font() *CT_Font {
	ret := &CT_Font{}
	return ret
}

func (m *CT_Font) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Name != nil {
		sename := xml.StartElement{Name: xml.Name{Local: "x:name"}}
		e.EncodeElement(m.Name, sename)
	}
	if m.Charset != nil {
		secharset := xml.StartElement{Name: xml.Name{Local: "x:charset"}}
		e.EncodeElement(m.Charset, secharset)
	}
	if m.Family != nil {
		sefamily := xml.StartElement{Name: xml.Name{Local: "x:family"}}
		e.EncodeElement(m.Family, sefamily)
	}
	if m.B != nil {
		seb := xml.StartElement{Name: xml.Name{Local: "x:b"}}
		e.EncodeElement(m.B, seb)
	}
	if m.I != nil {
		sei := xml.StartElement{Name: xml.Name{Local: "x:i"}}
		e.EncodeElement(m.I, sei)
	}
	if m.Strike != nil {
		sestrike := xml.StartElement{Name: xml.Name{Local: "x:strike"}}
		e.EncodeElement(m.Strike, sestrike)
	}
	if m.Outline != nil {
		seoutline := xml.StartElement{Name: xml.Name{Local: "x:outline"}}
		e.EncodeElement(m.Outline, seoutline)
	}
	if m.Shadow != nil {
		seshadow := xml.StartElement{Name: xml.Name{Local: "x:shadow"}}
		e.EncodeElement(m.Shadow, seshadow)
	}
	if m.Condense != nil {
		secondense := xml.StartElement{Name: xml.Name{Local: "x:condense"}}
		e.EncodeElement(m.Condense, secondense)
	}
	if m.Extend != nil {
		seextend := xml.StartElement{Name: xml.Name{Local: "x:extend"}}
		e.EncodeElement(m.Extend, seextend)
	}
	if m.Color != nil {
		secolor := xml.StartElement{Name: xml.Name{Local: "x:color"}}
		e.EncodeElement(m.Color, secolor)
	}
	if m.Sz != nil {
		sesz := xml.StartElement{Name: xml.Name{Local: "x:sz"}}
		e.EncodeElement(m.Sz, sesz)
	}
	if m.U != nil {
		seu := xml.StartElement{Name: xml.Name{Local: "x:u"}}
		e.EncodeElement(m.U, seu)
	}
	if m.VertAlign != nil {
		severtAlign := xml.StartElement{Name: xml.Name{Local: "x:vertAlign"}}
		e.EncodeElement(m.VertAlign, severtAlign)
	}
	if m.Scheme != nil {
		sescheme := xml.StartElement{Name: xml.Name{Local: "x:scheme"}}
		e.EncodeElement(m.Scheme, sescheme)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Font) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Font:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "name":
				tmp := NewCT_FontName()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Name = append(m.Name, tmp)
			case "charset":
				tmp := NewCT_IntProperty()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Charset = append(m.Charset, tmp)
			case "family":
				tmp := NewCT_FontFamily()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Family = append(m.Family, tmp)
			case "b":
				tmp := NewCT_BooleanProperty()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.B = append(m.B, tmp)
			case "i":
				tmp := NewCT_BooleanProperty()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.I = append(m.I, tmp)
			case "strike":
				tmp := NewCT_BooleanProperty()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Strike = append(m.Strike, tmp)
			case "outline":
				tmp := NewCT_BooleanProperty()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Outline = append(m.Outline, tmp)
			case "shadow":
				tmp := NewCT_BooleanProperty()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Shadow = append(m.Shadow, tmp)
			case "condense":
				tmp := NewCT_BooleanProperty()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Condense = append(m.Condense, tmp)
			case "extend":
				tmp := NewCT_BooleanProperty()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Extend = append(m.Extend, tmp)
			case "color":
				tmp := NewCT_Color()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Color = append(m.Color, tmp)
			case "sz":
				tmp := NewCT_FontSize()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Sz = append(m.Sz, tmp)
			case "u":
				tmp := NewCT_UnderlineProperty()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.U = append(m.U, tmp)
			case "vertAlign":
				tmp := NewCT_VerticalAlignFontProperty()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.VertAlign = append(m.VertAlign, tmp)
			case "scheme":
				tmp := NewCT_FontScheme()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Scheme = append(m.Scheme, tmp)
			default:
				log.Printf("skipping unsupported element on CT_Font %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Font
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Font and its children
func (m *CT_Font) Validate() error {
	return m.ValidateWithPath("CT_Font")
}

// ValidateWithPath validates the CT_Font and its children, prefixing error messages with path
func (m *CT_Font) ValidateWithPath(path string) error {
	for i, v := range m.Name {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Name[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Charset {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Charset[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Family {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Family[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.B {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/B[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.I {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/I[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Strike {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Strike[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Outline {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Outline[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Shadow {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Shadow[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Condense {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Condense[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Extend {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Extend[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Color {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Color[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Sz {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Sz[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.U {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/U[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.VertAlign {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/VertAlign[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Scheme {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Scheme[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
