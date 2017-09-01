// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"log"
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type ViewPr struct {
	CT_ViewProperties
}

func NewViewPr() *ViewPr {
	ret := &ViewPr{}
	ret.CT_ViewProperties = *NewCT_ViewProperties()
	return ret
}

func (m *ViewPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/presentationml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:p"}, Value: "http://schemas.openxmlformats.org/presentationml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "p:viewPr"
	return m.CT_ViewProperties.MarshalXML(e, start)
}

func (m *ViewPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_ViewProperties = *NewCT_ViewProperties()
	for _, attr := range start.Attr {
		if attr.Name.Local == "lastView" {
			m.LastViewAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "showComments" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowCommentsAttr = &parsed
		}
	}
lViewPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "normalViewPr":
				m.NormalViewPr = NewCT_NormalViewProperties()
				if err := d.DecodeElement(m.NormalViewPr, &el); err != nil {
					return err
				}
			case "slideViewPr":
				m.SlideViewPr = NewCT_SlideViewProperties()
				if err := d.DecodeElement(m.SlideViewPr, &el); err != nil {
					return err
				}
			case "outlineViewPr":
				m.OutlineViewPr = NewCT_OutlineViewProperties()
				if err := d.DecodeElement(m.OutlineViewPr, &el); err != nil {
					return err
				}
			case "notesTextViewPr":
				m.NotesTextViewPr = NewCT_NotesTextViewProperties()
				if err := d.DecodeElement(m.NotesTextViewPr, &el); err != nil {
					return err
				}
			case "sorterViewPr":
				m.SorterViewPr = NewCT_SlideSorterViewProperties()
				if err := d.DecodeElement(m.SorterViewPr, &el); err != nil {
					return err
				}
			case "notesViewPr":
				m.NotesViewPr = NewCT_NotesViewProperties()
				if err := d.DecodeElement(m.NotesViewPr, &el); err != nil {
					return err
				}
			case "gridSpacing":
				m.GridSpacing = drawingml.NewCT_PositiveSize2D()
				if err := d.DecodeElement(m.GridSpacing, &el); err != nil {
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
			break lViewPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the ViewPr and its children
func (m *ViewPr) Validate() error {
	return m.ValidateWithPath("ViewPr")
}

// ValidateWithPath validates the ViewPr and its children, prefixing error messages with path
func (m *ViewPr) ValidateWithPath(path string) error {
	if err := m.CT_ViewProperties.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
