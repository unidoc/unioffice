// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type ColorsDef struct {
	CT_ColorTransform
}

func NewColorsDef() *ColorsDef {
	ret := &ColorsDef{}
	ret.CT_ColorTransform = *NewCT_ColorTransform()
	return ret
}

func (m *ColorsDef) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/diagram"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:di"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/diagram"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "colorsDef"
	return m.CT_ColorTransform.MarshalXML(e, start)
}

func (m *ColorsDef) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_ColorTransform = *NewCT_ColorTransform()
	for _, attr := range start.Attr {
		if attr.Name.Local == "uniqueId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UniqueIdAttr = &parsed
		}
		if attr.Name.Local == "minVer" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MinVerAttr = &parsed
		}
	}
lColorsDef:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "title":
				tmp := NewCT_CTName()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Title = append(m.Title, tmp)
			case "desc":
				tmp := NewCT_CTDescription()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Desc = append(m.Desc, tmp)
			case "catLst":
				m.CatLst = NewCT_CTCategories()
				if err := d.DecodeElement(m.CatLst, &el); err != nil {
					return err
				}
			case "styleLbl":
				tmp := NewCT_CTStyleLabel()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.StyleLbl = append(m.StyleLbl, tmp)
			case "extLst":
				m.ExtLst = drawingml.NewCT_OfficeArtExtensionList()
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
			break lColorsDef
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the ColorsDef and its children
func (m *ColorsDef) Validate() error {
	return m.ValidateWithPath("ColorsDef")
}

// ValidateWithPath validates the ColorsDef and its children, prefixing error messages with path
func (m *ColorsDef) ValidateWithPath(path string) error {
	if err := m.CT_ColorTransform.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
