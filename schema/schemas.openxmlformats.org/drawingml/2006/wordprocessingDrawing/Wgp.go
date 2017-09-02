// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingDrawing

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type Wgp struct {
	CT_WordprocessingGroup
}

func NewWgp() *Wgp {
	ret := &Wgp{}
	ret.CT_WordprocessingGroup = *NewCT_WordprocessingGroup()
	return ret
}

func (m *Wgp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.CT_WordprocessingGroup.MarshalXML(e, start)
}

func (m *Wgp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_WordprocessingGroup = *NewCT_WordprocessingGroup()
lWgp:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cNvPr":
				m.CNvPr = drawingml.NewCT_NonVisualDrawingProps()
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case "cNvGrpSpPr":
				if err := d.DecodeElement(m.CNvGrpSpPr, &el); err != nil {
					return err
				}
			case "grpSpPr":
				if err := d.DecodeElement(m.GrpSpPr, &el); err != nil {
					return err
				}
			case "wsp":
				tmp := NewCT_WordprocessingGroupChoice()
				if err := d.DecodeElement(&tmp.Wsp, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "grpSp":
				tmp := NewCT_WordprocessingGroupChoice()
				if err := d.DecodeElement(&tmp.GrpSp, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "graphicFrame":
				tmp := NewCT_WordprocessingGroupChoice()
				if err := d.DecodeElement(&tmp.GraphicFrame, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "pic":
				tmp := NewCT_WordprocessingGroupChoice()
				if err := d.DecodeElement(&tmp.Pic, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "contentPart":
				tmp := NewCT_WordprocessingGroupChoice()
				if err := d.DecodeElement(&tmp.ContentPart, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "extLst":
				m.ExtLst = drawingml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on Wgp %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWgp
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Wgp and its children
func (m *Wgp) Validate() error {
	return m.ValidateWithPath("Wgp")
}

// ValidateWithPath validates the Wgp and its children, prefixing error messages with path
func (m *Wgp) ValidateWithPath(path string) error {
	if err := m.CT_WordprocessingGroup.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
