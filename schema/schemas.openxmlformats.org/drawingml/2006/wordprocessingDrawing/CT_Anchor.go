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

type CT_Anchor struct {
	DistTAttr          *uint32
	DistBAttr          *uint32
	DistLAttr          *uint32
	DistRAttr          *uint32
	SimplePosAttr      *bool
	RelativeHeightAttr uint32
	BehindDocAttr      bool
	LockedAttr         bool
	LayoutInCellAttr   bool
	HiddenAttr         *bool
	AllowOverlapAttr   bool
	SimplePos          *drawingml.CT_Point2D
	PositionH          *CT_PosH
	PositionV          *CT_PosV
	Extent             *drawingml.CT_PositiveSize2D
	EffectExtent       *CT_EffectExtent
	Choice             *EG_WrapTypeChoice
	DocPr              *drawingml.CT_NonVisualDrawingProps
	CNvGraphicFramePr  *drawingml.CT_NonVisualGraphicFrameProperties
	Graphic            *drawingml.Graphic
}

func NewCT_Anchor() *CT_Anchor {
	ret := &CT_Anchor{}
	ret.SimplePos = drawingml.NewCT_Point2D()
	ret.PositionH = NewCT_PosH()
	ret.PositionV = NewCT_PosV()
	ret.Extent = drawingml.NewCT_PositiveSize2D()
	ret.DocPr = drawingml.NewCT_NonVisualDrawingProps()
	ret.Graphic = drawingml.NewGraphic()
	return ret
}
func (m *CT_Anchor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.DistTAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distT"},
			Value: fmt.Sprintf("%v", *m.DistTAttr)})
	}
	if m.DistBAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distB"},
			Value: fmt.Sprintf("%v", *m.DistBAttr)})
	}
	if m.DistLAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distL"},
			Value: fmt.Sprintf("%v", *m.DistLAttr)})
	}
	if m.DistRAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distR"},
			Value: fmt.Sprintf("%v", *m.DistRAttr)})
	}
	if m.SimplePosAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "simplePos"},
			Value: fmt.Sprintf("%v", *m.SimplePosAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "relativeHeight"},
		Value: fmt.Sprintf("%v", m.RelativeHeightAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "behindDoc"},
		Value: fmt.Sprintf("%v", m.BehindDocAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "locked"},
		Value: fmt.Sprintf("%v", m.LockedAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "layoutInCell"},
		Value: fmt.Sprintf("%v", m.LayoutInCellAttr)})
	if m.HiddenAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hidden"},
			Value: fmt.Sprintf("%v", *m.HiddenAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "allowOverlap"},
		Value: fmt.Sprintf("%v", m.AllowOverlapAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	sesimplePos := xml.StartElement{Name: xml.Name{Local: "wp:simplePos"}}
	e.EncodeElement(m.SimplePos, sesimplePos)
	sepositionH := xml.StartElement{Name: xml.Name{Local: "wp:positionH"}}
	e.EncodeElement(m.PositionH, sepositionH)
	sepositionV := xml.StartElement{Name: xml.Name{Local: "wp:positionV"}}
	e.EncodeElement(m.PositionV, sepositionV)
	seextent := xml.StartElement{Name: xml.Name{Local: "wp:extent"}}
	e.EncodeElement(m.Extent, seextent)
	if m.EffectExtent != nil {
		seeffectExtent := xml.StartElement{Name: xml.Name{Local: "wp:effectExtent"}}
		e.EncodeElement(m.EffectExtent, seeffectExtent)
	}
	if m.Choice != nil {
		m.Choice.MarshalXML(e, start)
	}
	sedocPr := xml.StartElement{Name: xml.Name{Local: "wp:docPr"}}
	e.EncodeElement(m.DocPr, sedocPr)
	if m.CNvGraphicFramePr != nil {
		secNvGraphicFramePr := xml.StartElement{Name: xml.Name{Local: "wp:cNvGraphicFramePr"}}
		e.EncodeElement(m.CNvGraphicFramePr, secNvGraphicFramePr)
	}
	segraphic := xml.StartElement{Name: xml.Name{Local: "a:graphic"}}
	e.EncodeElement(m.Graphic, segraphic)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Anchor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.SimplePos = drawingml.NewCT_Point2D()
	m.PositionH = NewCT_PosH()
	m.PositionV = NewCT_PosV()
	m.Extent = drawingml.NewCT_PositiveSize2D()
	m.DocPr = drawingml.NewCT_NonVisualDrawingProps()
	m.Graphic = drawingml.NewGraphic()
	for _, attr := range start.Attr {
		if attr.Name.Local == "distT" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistTAttr = &pt
		}
		if attr.Name.Local == "distB" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistBAttr = &pt
		}
		if attr.Name.Local == "distL" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistLAttr = &pt
		}
		if attr.Name.Local == "distR" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistRAttr = &pt
		}
		if attr.Name.Local == "simplePos" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SimplePosAttr = &parsed
		}
		if attr.Name.Local == "relativeHeight" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.RelativeHeightAttr = uint32(parsed)
		}
		if attr.Name.Local == "behindDoc" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.BehindDocAttr = parsed
		}
		if attr.Name.Local == "locked" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LockedAttr = parsed
		}
		if attr.Name.Local == "layoutInCell" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LayoutInCellAttr = parsed
		}
		if attr.Name.Local == "hidden" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HiddenAttr = &parsed
		}
		if attr.Name.Local == "allowOverlap" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AllowOverlapAttr = parsed
		}
	}
lCT_Anchor:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "simplePos":
				if err := d.DecodeElement(m.SimplePos, &el); err != nil {
					return err
				}
			case "positionH":
				if err := d.DecodeElement(m.PositionH, &el); err != nil {
					return err
				}
			case "positionV":
				if err := d.DecodeElement(m.PositionV, &el); err != nil {
					return err
				}
			case "extent":
				if err := d.DecodeElement(m.Extent, &el); err != nil {
					return err
				}
			case "effectExtent":
				m.EffectExtent = NewCT_EffectExtent()
				if err := d.DecodeElement(m.EffectExtent, &el); err != nil {
					return err
				}
			case "wrapNone":
				m.Choice = NewEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapNone, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "wrapSquare":
				m.Choice = NewEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapSquare, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "wrapTight":
				m.Choice = NewEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapTight, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "wrapThrough":
				m.Choice = NewEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapThrough, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "wrapTopAndBottom":
				m.Choice = NewEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapTopAndBottom, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "docPr":
				if err := d.DecodeElement(m.DocPr, &el); err != nil {
					return err
				}
			case "cNvGraphicFramePr":
				m.CNvGraphicFramePr = drawingml.NewCT_NonVisualGraphicFrameProperties()
				if err := d.DecodeElement(m.CNvGraphicFramePr, &el); err != nil {
					return err
				}
			case "graphic":
				if err := d.DecodeElement(m.Graphic, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Anchor
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Anchor) Validate() error {
	return m.ValidateWithPath("CT_Anchor")
}
func (m *CT_Anchor) ValidateWithPath(path string) error {
	if err := m.SimplePos.ValidateWithPath(path + "/SimplePos"); err != nil {
		return err
	}
	if err := m.PositionH.ValidateWithPath(path + "/PositionH"); err != nil {
		return err
	}
	if err := m.PositionV.ValidateWithPath(path + "/PositionV"); err != nil {
		return err
	}
	if err := m.Extent.ValidateWithPath(path + "/Extent"); err != nil {
		return err
	}
	if m.EffectExtent != nil {
		if err := m.EffectExtent.ValidateWithPath(path + "/EffectExtent"); err != nil {
			return err
		}
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	if err := m.DocPr.ValidateWithPath(path + "/DocPr"); err != nil {
		return err
	}
	if m.CNvGraphicFramePr != nil {
		if err := m.CNvGraphicFramePr.ValidateWithPath(path + "/CNvGraphicFramePr"); err != nil {
			return err
		}
	}
	if err := m.Graphic.ValidateWithPath(path + "/Graphic"); err != nil {
		return err
	}
	return nil
}
