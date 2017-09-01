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

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/picture"
)

type CT_WordprocessingGroupChoice struct {
	Wsp          []*Wsp
	GrpSp        []*CT_WordprocessingGroup
	GraphicFrame []*CT_GraphicFrame
	Pic          []*picture.Pic
	ContentPart  []*CT_WordprocessingContentPart
}

func NewCT_WordprocessingGroupChoice() *CT_WordprocessingGroupChoice {
	ret := &CT_WordprocessingGroupChoice{}
	return ret
}

func (m *CT_WordprocessingGroupChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.Wsp != nil {
		sewsp := xml.StartElement{Name: xml.Name{Local: "wp:wsp"}}
		e.EncodeElement(m.Wsp, sewsp)
	}
	if m.GrpSp != nil {
		segrpSp := xml.StartElement{Name: xml.Name{Local: "wp:grpSp"}}
		e.EncodeElement(m.GrpSp, segrpSp)
	}
	if m.GraphicFrame != nil {
		segraphicFrame := xml.StartElement{Name: xml.Name{Local: "wp:graphicFrame"}}
		e.EncodeElement(m.GraphicFrame, segraphicFrame)
	}
	if m.Pic != nil {
		sepic := xml.StartElement{Name: xml.Name{Local: "pic:pic"}}
		e.EncodeElement(m.Pic, sepic)
	}
	if m.ContentPart != nil {
		secontentPart := xml.StartElement{Name: xml.Name{Local: "wp:contentPart"}}
		e.EncodeElement(m.ContentPart, secontentPart)
	}
	return nil
}

func (m *CT_WordprocessingGroupChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_WordprocessingGroupChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "wsp":
				tmp := NewWsp()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Wsp = append(m.Wsp, tmp)
			case "grpSp":
				tmp := NewCT_WordprocessingGroup()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.GrpSp = append(m.GrpSp, tmp)
			case "graphicFrame":
				tmp := NewCT_GraphicFrame()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.GraphicFrame = append(m.GraphicFrame, tmp)
			case "pic":
				tmp := picture.NewPic()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Pic = append(m.Pic, tmp)
			case "contentPart":
				tmp := NewCT_WordprocessingContentPart()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ContentPart = append(m.ContentPart, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_WordprocessingGroupChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_WordprocessingGroupChoice and its children
func (m *CT_WordprocessingGroupChoice) Validate() error {
	return m.ValidateWithPath("CT_WordprocessingGroupChoice")
}

// ValidateWithPath validates the CT_WordprocessingGroupChoice and its children, prefixing error messages with path
func (m *CT_WordprocessingGroupChoice) ValidateWithPath(path string) error {
	for i, v := range m.Wsp {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Wsp[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.GrpSp {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/GrpSp[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.GraphicFrame {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/GraphicFrame[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Pic {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Pic[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.ContentPart {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ContentPart[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
