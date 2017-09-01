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

type CT_WordprocessingCanvasChoice struct {
	Wsp          []*Wsp
	Pic          []*picture.Pic
	ContentPart  []*CT_WordprocessingContentPart
	Wgp          []*Wgp
	GraphicFrame []*CT_GraphicFrame
}

func NewCT_WordprocessingCanvasChoice() *CT_WordprocessingCanvasChoice {
	ret := &CT_WordprocessingCanvasChoice{}
	return ret
}

func (m *CT_WordprocessingCanvasChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.Wsp != nil {
		sewsp := xml.StartElement{Name: xml.Name{Local: "wp:wsp"}}
		e.EncodeElement(m.Wsp, sewsp)
	}
	if m.Pic != nil {
		sepic := xml.StartElement{Name: xml.Name{Local: "pic:pic"}}
		e.EncodeElement(m.Pic, sepic)
	}
	if m.ContentPart != nil {
		secontentPart := xml.StartElement{Name: xml.Name{Local: "wp:contentPart"}}
		e.EncodeElement(m.ContentPart, secontentPart)
	}
	if m.Wgp != nil {
		sewgp := xml.StartElement{Name: xml.Name{Local: "wp:wgp"}}
		e.EncodeElement(m.Wgp, sewgp)
	}
	if m.GraphicFrame != nil {
		segraphicFrame := xml.StartElement{Name: xml.Name{Local: "wp:graphicFrame"}}
		e.EncodeElement(m.GraphicFrame, segraphicFrame)
	}
	return nil
}

func (m *CT_WordprocessingCanvasChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_WordprocessingCanvasChoice:
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
			case "wgp":
				tmp := NewWgp()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Wgp = append(m.Wgp, tmp)
			case "graphicFrame":
				tmp := NewCT_GraphicFrame()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.GraphicFrame = append(m.GraphicFrame, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_WordprocessingCanvasChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_WordprocessingCanvasChoice and its children
func (m *CT_WordprocessingCanvasChoice) Validate() error {
	return m.ValidateWithPath("CT_WordprocessingCanvasChoice")
}

// ValidateWithPath validates the CT_WordprocessingCanvasChoice and its children, prefixing error messages with path
func (m *CT_WordprocessingCanvasChoice) ValidateWithPath(path string) error {
	for i, v := range m.Wsp {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Wsp[%d]", path, i)); err != nil {
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
	for i, v := range m.Wgp {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Wgp[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.GraphicFrame {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/GraphicFrame[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
