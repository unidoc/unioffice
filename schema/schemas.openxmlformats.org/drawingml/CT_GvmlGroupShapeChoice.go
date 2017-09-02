// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_GvmlGroupShapeChoice struct {
	TxSp         []*CT_GvmlTextShape
	Sp           []*CT_GvmlShape
	CxnSp        []*CT_GvmlConnector
	Pic          []*CT_GvmlPicture
	GraphicFrame []*CT_GvmlGraphicalObjectFrame
	GrpSp        []*CT_GvmlGroupShape
}

func NewCT_GvmlGroupShapeChoice() *CT_GvmlGroupShapeChoice {
	ret := &CT_GvmlGroupShapeChoice{}
	return ret
}

func (m *CT_GvmlGroupShapeChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TxSp != nil {
		setxSp := xml.StartElement{Name: xml.Name{Local: "a:txSp"}}
		e.EncodeElement(m.TxSp, setxSp)
	}
	if m.Sp != nil {
		sesp := xml.StartElement{Name: xml.Name{Local: "a:sp"}}
		e.EncodeElement(m.Sp, sesp)
	}
	if m.CxnSp != nil {
		secxnSp := xml.StartElement{Name: xml.Name{Local: "a:cxnSp"}}
		e.EncodeElement(m.CxnSp, secxnSp)
	}
	if m.Pic != nil {
		sepic := xml.StartElement{Name: xml.Name{Local: "a:pic"}}
		e.EncodeElement(m.Pic, sepic)
	}
	if m.GraphicFrame != nil {
		segraphicFrame := xml.StartElement{Name: xml.Name{Local: "a:graphicFrame"}}
		e.EncodeElement(m.GraphicFrame, segraphicFrame)
	}
	if m.GrpSp != nil {
		segrpSp := xml.StartElement{Name: xml.Name{Local: "a:grpSp"}}
		e.EncodeElement(m.GrpSp, segrpSp)
	}
	return nil
}

func (m *CT_GvmlGroupShapeChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_GvmlGroupShapeChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "txSp":
				tmp := NewCT_GvmlTextShape()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.TxSp = append(m.TxSp, tmp)
			case "sp":
				tmp := NewCT_GvmlShape()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Sp = append(m.Sp, tmp)
			case "cxnSp":
				tmp := NewCT_GvmlConnector()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CxnSp = append(m.CxnSp, tmp)
			case "pic":
				tmp := NewCT_GvmlPicture()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Pic = append(m.Pic, tmp)
			case "graphicFrame":
				tmp := NewCT_GvmlGraphicalObjectFrame()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.GraphicFrame = append(m.GraphicFrame, tmp)
			case "grpSp":
				tmp := NewCT_GvmlGroupShape()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.GrpSp = append(m.GrpSp, tmp)
			default:
				log.Printf("skipping unsupported element on CT_GvmlGroupShapeChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GvmlGroupShapeChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GvmlGroupShapeChoice and its children
func (m *CT_GvmlGroupShapeChoice) Validate() error {
	return m.ValidateWithPath("CT_GvmlGroupShapeChoice")
}

// ValidateWithPath validates the CT_GvmlGroupShapeChoice and its children, prefixing error messages with path
func (m *CT_GvmlGroupShapeChoice) ValidateWithPath(path string) error {
	for i, v := range m.TxSp {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/TxSp[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Sp {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Sp[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.CxnSp {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CxnSp[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Pic {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Pic[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.GraphicFrame {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/GraphicFrame[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.GrpSp {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/GrpSp[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
