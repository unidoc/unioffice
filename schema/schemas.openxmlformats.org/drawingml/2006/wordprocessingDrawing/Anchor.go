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
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type Anchor struct {
	CT_Anchor
}

func NewAnchor() *Anchor {
	ret := &Anchor{}
	ret.CT_Anchor = *NewCT_Anchor()
	return ret
}

func (m *Anchor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:pic"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/picture"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:w"}, Value: "http://schemas.openxmlformats.org/wordprocessingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:wp"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "wp:anchor"
	return m.CT_Anchor.MarshalXML(e, start)
}

func (m *Anchor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Anchor = *NewCT_Anchor()
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
lAnchor:
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
				log.Printf("skipping unsupported element on Anchor %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lAnchor
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Anchor and its children
func (m *Anchor) Validate() error {
	return m.ValidateWithPath("Anchor")
}

// ValidateWithPath validates the Anchor and its children, prefixing error messages with path
func (m *Anchor) ValidateWithPath(path string) error {
	if err := m.CT_Anchor.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
