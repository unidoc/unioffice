// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_NotesSlide struct {
	// Common slide data for notes slides
	CSld *CT_CommonSlideData
	// Color Scheme Map Override
	ClrMapOvr            *drawingml.CT_ColorMappingOverride
	ExtLst               *CT_ExtensionListModify
	ShowMasterSpAttr     *bool
	ShowMasterPhAnimAttr *bool
}

func NewCT_NotesSlide() *CT_NotesSlide {
	ret := &CT_NotesSlide{}
	ret.CSld = NewCT_CommonSlideData()
	return ret
}

func (m *CT_NotesSlide) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.ShowMasterSpAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showMasterSp"},
			Value: fmt.Sprintf("%v", *m.ShowMasterSpAttr)})
	}
	if m.ShowMasterPhAnimAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showMasterPhAnim"},
			Value: fmt.Sprintf("%v", *m.ShowMasterPhAnimAttr)})
	}
	e.EncodeToken(start)
	secSld := xml.StartElement{Name: xml.Name{Local: "p:cSld"}}
	e.EncodeElement(m.CSld, secSld)
	if m.ClrMapOvr != nil {
		seclrMapOvr := xml.StartElement{Name: xml.Name{Local: "p:clrMapOvr"}}
		e.EncodeElement(m.ClrMapOvr, seclrMapOvr)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_NotesSlide) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CSld = NewCT_CommonSlideData()
	for _, attr := range start.Attr {
		if attr.Name.Local == "showMasterSp" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowMasterSpAttr = &parsed
		}
		if attr.Name.Local == "showMasterPhAnim" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowMasterPhAnimAttr = &parsed
		}
	}
lCT_NotesSlide:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cSld":
				if err := d.DecodeElement(m.CSld, &el); err != nil {
					return err
				}
			case "clrMapOvr":
				m.ClrMapOvr = drawingml.NewCT_ColorMappingOverride()
				if err := d.DecodeElement(m.ClrMapOvr, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionListModify()
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
			break lCT_NotesSlide
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_NotesSlide and its children
func (m *CT_NotesSlide) Validate() error {
	return m.ValidateWithPath("CT_NotesSlide")
}

// ValidateWithPath validates the CT_NotesSlide and its children, prefixing error messages with path
func (m *CT_NotesSlide) ValidateWithPath(path string) error {
	if err := m.CSld.ValidateWithPath(path + "/CSld"); err != nil {
		return err
	}
	if m.ClrMapOvr != nil {
		if err := m.ClrMapOvr.ValidateWithPath(path + "/ClrMapOvr"); err != nil {
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
