// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"
)

type CT_Drawing struct {
	Anchor []*wordprocessingDrawing.Anchor
	Inline []*wordprocessingDrawing.Inline
}

func NewCT_Drawing() *CT_Drawing {
	ret := &CT_Drawing{}
	return ret
}
func (m *CT_Drawing) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.Anchor != nil {
		seanchor := xml.StartElement{Name: xml.Name{Local: "wp:anchor"}}
		e.EncodeElement(m.Anchor, seanchor)
	}
	if m.Inline != nil {
		seinline := xml.StartElement{Name: xml.Name{Local: "wp:inline"}}
		e.EncodeElement(m.Inline, seinline)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Drawing) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Drawing:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "anchor":
				tmp := wordprocessingDrawing.NewAnchor()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Anchor = append(m.Anchor, tmp)
			case "inline":
				tmp := wordprocessingDrawing.NewInline()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Inline = append(m.Inline, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Drawing
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Drawing) Validate() error {
	return m.ValidateWithPath("CT_Drawing")
}
func (m *CT_Drawing) ValidateWithPath(path string) error {
	for i, v := range m.Anchor {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Anchor[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Inline {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Inline[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
