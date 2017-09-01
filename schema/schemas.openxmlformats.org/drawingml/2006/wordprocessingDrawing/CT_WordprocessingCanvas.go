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

type CT_WordprocessingCanvas struct {
	Bg     *drawingml.CT_BackgroundFormatting
	Whole  *drawingml.CT_WholeE2oFormatting
	Choice []*CT_WordprocessingCanvasChoice
	ExtLst *drawingml.CT_OfficeArtExtensionList
}

func NewCT_WordprocessingCanvas() *CT_WordprocessingCanvas {
	ret := &CT_WordprocessingCanvas{}
	return ret
}

func (m *CT_WordprocessingCanvas) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Bg != nil {
		sebg := xml.StartElement{Name: xml.Name{Local: "wp:bg"}}
		e.EncodeElement(m.Bg, sebg)
	}
	if m.Whole != nil {
		sewhole := xml.StartElement{Name: xml.Name{Local: "wp:whole"}}
		e.EncodeElement(m.Whole, sewhole)
	}
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

func (m *CT_WordprocessingCanvas) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_WordprocessingCanvas:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "bg":
				m.Bg = drawingml.NewCT_BackgroundFormatting()
				if err := d.DecodeElement(m.Bg, &el); err != nil {
					return err
				}
			case "whole":
				m.Whole = drawingml.NewCT_WholeE2oFormatting()
				if err := d.DecodeElement(m.Whole, &el); err != nil {
					return err
				}
			case "wsp":
				tmp := NewCT_WordprocessingCanvasChoice()
				if err := d.DecodeElement(&tmp.Wsp, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "pic":
				tmp := NewCT_WordprocessingCanvasChoice()
				if err := d.DecodeElement(&tmp.Pic, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "contentPart":
				tmp := NewCT_WordprocessingCanvasChoice()
				if err := d.DecodeElement(&tmp.ContentPart, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "wgp":
				tmp := NewCT_WordprocessingCanvasChoice()
				if err := d.DecodeElement(&tmp.Wgp, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case "graphicFrame":
				tmp := NewCT_WordprocessingCanvasChoice()
				if err := d.DecodeElement(&tmp.GraphicFrame, &el); err != nil {
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
			break lCT_WordprocessingCanvas
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_WordprocessingCanvas and its children
func (m *CT_WordprocessingCanvas) Validate() error {
	return m.ValidateWithPath("CT_WordprocessingCanvas")
}

// ValidateWithPath validates the CT_WordprocessingCanvas and its children, prefixing error messages with path
func (m *CT_WordprocessingCanvas) ValidateWithPath(path string) error {
	if m.Bg != nil {
		if err := m.Bg.ValidateWithPath(path + "/Bg"); err != nil {
			return err
		}
	}
	if m.Whole != nil {
		if err := m.Whole.ValidateWithPath(path + "/Whole"); err != nil {
			return err
		}
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
