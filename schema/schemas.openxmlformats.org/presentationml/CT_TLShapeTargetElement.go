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

type CT_TLShapeTargetElement struct {
	// Shape ID
	SpidAttr uint32
	// Background
	Bg *CT_Empty
	// Subshape
	SubSp *CT_TLSubShapeId
	// Embedded Chart Element
	OleChartEl *CT_TLOleChartTargetElement
	// Text Element
	TxEl *CT_TLTextTargetElement
	// Graphic Element
	GraphicEl *drawingml.CT_AnimationElementChoice
}

func NewCT_TLShapeTargetElement() *CT_TLShapeTargetElement {
	ret := &CT_TLShapeTargetElement{}
	return ret
}

func (m *CT_TLShapeTargetElement) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "spid"},
		Value: fmt.Sprintf("%v", m.SpidAttr)})
	e.EncodeToken(start)
	if m.Bg != nil {
		sebg := xml.StartElement{Name: xml.Name{Local: "p:bg"}}
		e.EncodeElement(m.Bg, sebg)
	}
	if m.SubSp != nil {
		sesubSp := xml.StartElement{Name: xml.Name{Local: "p:subSp"}}
		e.EncodeElement(m.SubSp, sesubSp)
	}
	if m.OleChartEl != nil {
		seoleChartEl := xml.StartElement{Name: xml.Name{Local: "p:oleChartEl"}}
		e.EncodeElement(m.OleChartEl, seoleChartEl)
	}
	if m.TxEl != nil {
		setxEl := xml.StartElement{Name: xml.Name{Local: "p:txEl"}}
		e.EncodeElement(m.TxEl, setxEl)
	}
	if m.GraphicEl != nil {
		segraphicEl := xml.StartElement{Name: xml.Name{Local: "p:graphicEl"}}
		e.EncodeElement(m.GraphicEl, segraphicEl)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLShapeTargetElement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "spid" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.SpidAttr = uint32(parsed)
		}
	}
lCT_TLShapeTargetElement:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "bg":
				m.Bg = NewCT_Empty()
				if err := d.DecodeElement(m.Bg, &el); err != nil {
					return err
				}
			case "subSp":
				m.SubSp = NewCT_TLSubShapeId()
				if err := d.DecodeElement(m.SubSp, &el); err != nil {
					return err
				}
			case "oleChartEl":
				m.OleChartEl = NewCT_TLOleChartTargetElement()
				if err := d.DecodeElement(m.OleChartEl, &el); err != nil {
					return err
				}
			case "txEl":
				m.TxEl = NewCT_TLTextTargetElement()
				if err := d.DecodeElement(m.TxEl, &el); err != nil {
					return err
				}
			case "graphicEl":
				m.GraphicEl = drawingml.NewCT_AnimationElementChoice()
				if err := d.DecodeElement(m.GraphicEl, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_TLShapeTargetElement %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLShapeTargetElement
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLShapeTargetElement and its children
func (m *CT_TLShapeTargetElement) Validate() error {
	return m.ValidateWithPath("CT_TLShapeTargetElement")
}

// ValidateWithPath validates the CT_TLShapeTargetElement and its children, prefixing error messages with path
func (m *CT_TLShapeTargetElement) ValidateWithPath(path string) error {
	if m.Bg != nil {
		if err := m.Bg.ValidateWithPath(path + "/Bg"); err != nil {
			return err
		}
	}
	if m.SubSp != nil {
		if err := m.SubSp.ValidateWithPath(path + "/SubSp"); err != nil {
			return err
		}
	}
	if m.OleChartEl != nil {
		if err := m.OleChartEl.ValidateWithPath(path + "/OleChartEl"); err != nil {
			return err
		}
	}
	if m.TxEl != nil {
		if err := m.TxEl.ValidateWithPath(path + "/TxEl"); err != nil {
			return err
		}
	}
	if m.GraphicEl != nil {
		if err := m.GraphicEl.ValidateWithPath(path + "/GraphicEl"); err != nil {
			return err
		}
	}
	return nil
}
