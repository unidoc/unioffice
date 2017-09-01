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
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_WordprocessingShape struct {
	NormalEastAsianFlowAttr *bool
	CNvPr                   *drawingml.CT_NonVisualDrawingProps
	Choice                  *CT_WordprocessingShapeChoice
	SpPr                    *drawingml.CT_ShapeProperties
	Style                   *drawingml.CT_ShapeStyle
	ExtLst                  *drawingml.CT_OfficeArtExtensionList
	CChoice                 *CT_WordprocessingShapeChoice1
	BodyPr                  *drawingml.CT_TextBodyProperties
}

func NewCT_WordprocessingShape() *CT_WordprocessingShape {
	ret := &CT_WordprocessingShape{}
	ret.Choice = NewCT_WordprocessingShapeChoice()
	ret.SpPr = drawingml.NewCT_ShapeProperties()
	ret.BodyPr = drawingml.NewCT_TextBodyProperties()
	return ret
}

func (m *CT_WordprocessingShape) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.NormalEastAsianFlowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "normalEastAsianFlow"},
			Value: fmt.Sprintf("%v", *m.NormalEastAsianFlowAttr)})
	}
	e.EncodeToken(start)
	if m.CNvPr != nil {
		secNvPr := xml.StartElement{Name: xml.Name{Local: "wp:cNvPr"}}
		e.EncodeElement(m.CNvPr, secNvPr)
	}
	m.Choice.MarshalXML(e, start)
	sespPr := xml.StartElement{Name: xml.Name{Local: "wp:spPr"}}
	e.EncodeElement(m.SpPr, sespPr)
	if m.Style != nil {
		sestyle := xml.StartElement{Name: xml.Name{Local: "wp:style"}}
		e.EncodeElement(m.Style, sestyle)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "wp:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	if m.CChoice != nil {
		m.CChoice.MarshalXML(e, start)
	}
	sebodyPr := xml.StartElement{Name: xml.Name{Local: "wp:bodyPr"}}
	e.EncodeElement(m.BodyPr, sebodyPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_WordprocessingShape) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Choice = NewCT_WordprocessingShapeChoice()
	m.SpPr = drawingml.NewCT_ShapeProperties()
	m.BodyPr = drawingml.NewCT_TextBodyProperties()
	for _, attr := range start.Attr {
		if attr.Name.Local == "normalEastAsianFlow" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NormalEastAsianFlowAttr = &parsed
		}
	}
lCT_WordprocessingShape:
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
			case "cNvSpPr":
				m.Choice = NewCT_WordprocessingShapeChoice()
				if err := d.DecodeElement(&m.Choice.CNvSpPr, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "cNvCnPr":
				m.Choice = NewCT_WordprocessingShapeChoice()
				if err := d.DecodeElement(&m.Choice.CNvCnPr, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "spPr":
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case "style":
				m.Style = drawingml.NewCT_ShapeStyle()
				if err := d.DecodeElement(m.Style, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = drawingml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			case "txbx":
				m.CChoice = NewCT_WordprocessingShapeChoice1()
				if err := d.DecodeElement(&m.CChoice.Txbx, &el); err != nil {
					return err
				}
				_ = m.CChoice
			case "linkedTxbx":
				m.CChoice = NewCT_WordprocessingShapeChoice1()
				if err := d.DecodeElement(&m.CChoice.LinkedTxbx, &el); err != nil {
					return err
				}
				_ = m.CChoice
			case "bodyPr":
				if err := d.DecodeElement(m.BodyPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_WordprocessingShape %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_WordprocessingShape
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_WordprocessingShape and its children
func (m *CT_WordprocessingShape) Validate() error {
	return m.ValidateWithPath("CT_WordprocessingShape")
}

// ValidateWithPath validates the CT_WordprocessingShape and its children, prefixing error messages with path
func (m *CT_WordprocessingShape) ValidateWithPath(path string) error {
	if m.CNvPr != nil {
		if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
			return err
		}
	}
	if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
		return err
	}
	if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
		return err
	}
	if m.Style != nil {
		if err := m.Style.ValidateWithPath(path + "/Style"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	if m.CChoice != nil {
		if err := m.CChoice.ValidateWithPath(path + "/CChoice"); err != nil {
			return err
		}
	}
	if err := m.BodyPr.ValidateWithPath(path + "/BodyPr"); err != nil {
		return err
	}
	return nil
}
