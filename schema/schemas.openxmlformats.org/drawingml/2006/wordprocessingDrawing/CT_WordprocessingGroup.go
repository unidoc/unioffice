// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingDrawing

import (
	"encoding/xml"
	"fmt"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_WordprocessingGroup struct {
	CNvPr      *drawingml.CT_NonVisualDrawingProps
	CNvGrpSpPr *drawingml.CT_NonVisualGroupDrawingShapeProps
	GrpSpPr    *drawingml.CT_GroupShapeProperties
	Choice     []*CT_WordprocessingGroupChoice
	ExtLst     *drawingml.CT_OfficeArtExtensionList
}

func NewCT_WordprocessingGroup() *CT_WordprocessingGroup {
	ret := &CT_WordprocessingGroup{}
	ret.CNvGrpSpPr = drawingml.NewCT_NonVisualGroupDrawingShapeProps()
	ret.GrpSpPr = drawingml.NewCT_GroupShapeProperties()
	return ret
}

func (m *CT_WordprocessingGroup) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.CNvPr != nil {
		secNvPr := xml.StartElement{Name: xml.Name{Local: "wp:cNvPr"}}
		e.EncodeElement(m.CNvPr, secNvPr)
	}
	secNvGrpSpPr := xml.StartElement{Name: xml.Name{Local: "wp:cNvGrpSpPr"}}
	e.EncodeElement(m.CNvGrpSpPr, secNvGrpSpPr)
	segrpSpPr := xml.StartElement{Name: xml.Name{Local: "wp:grpSpPr"}}
	e.EncodeElement(m.GrpSpPr, segrpSpPr)
	if m.Choice != nil {
		for _, c := range m.Choice {
			c.MarshalXML(e, start)
		}
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "wp:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_WordprocessingGroup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CNvGrpSpPr = drawingml.NewCT_NonVisualGroupDrawingShapeProps()
	m.GrpSpPr = drawingml.NewCT_GroupShapeProperties()
lCT_WordprocessingGroup:
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
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_WordprocessingGroup
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_WordprocessingGroup and its children
func (m *CT_WordprocessingGroup) Validate() error {
	return m.ValidateWithPath("CT_WordprocessingGroup")
}

// ValidateWithPath validates the CT_WordprocessingGroup and its children, prefixing error messages with path
func (m *CT_WordprocessingGroup) ValidateWithPath(path string) error {
	if m.CNvPr != nil {
		if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
			return err
		}
	}
	if err := m.CNvGrpSpPr.ValidateWithPath(path + "/CNvGrpSpPr"); err != nil {
		return err
	}
	if err := m.GrpSpPr.ValidateWithPath(path + "/GrpSpPr"); err != nil {
		return err
	}
	for i, v := range m.Choice {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Choice[%d]", path, i)); err != nil {
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
