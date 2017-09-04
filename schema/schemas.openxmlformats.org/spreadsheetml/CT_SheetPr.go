// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

	"baliance.com/gooxml"
)

type CT_SheetPr struct {
	// Synch Horizontal
	SyncHorizontalAttr *bool
	// Synch Vertical
	SyncVerticalAttr *bool
	// Synch Reference
	SyncRefAttr *string
	// Transition Formula Evaluation
	TransitionEvaluationAttr *bool
	// Transition Formula Entry
	TransitionEntryAttr *bool
	// Published
	PublishedAttr *bool
	// Code Name
	CodeNameAttr *string
	// Filter Mode
	FilterModeAttr *bool
	// Enable Conditional Formatting Calculations
	EnableFormatConditionsCalculationAttr *bool
	// Sheet Tab Color
	TabColor *CT_Color
	// Outline Properties
	OutlinePr *CT_OutlinePr
	// Page Setup Properties
	PageSetUpPr *CT_PageSetUpPr
}

func NewCT_SheetPr() *CT_SheetPr {
	ret := &CT_SheetPr{}
	ret.SyncHorizontalAttr = gooxml.Bool(false)
	ret.SyncVerticalAttr = gooxml.Bool(false)
	ret.TransitionEvaluationAttr = gooxml.Bool(false)
	ret.TransitionEntryAttr = gooxml.Bool(false)
	ret.PublishedAttr = gooxml.Bool(true)
	ret.FilterModeAttr = gooxml.Bool(false)
	ret.EnableFormatConditionsCalculationAttr = gooxml.Bool(true)
	return ret
}

func (m *CT_SheetPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.SyncHorizontalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "syncHorizontal"},
			Value: fmt.Sprintf("%d", b2i(*m.SyncHorizontalAttr))})
	}
	if m.SyncVerticalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "syncVertical"},
			Value: fmt.Sprintf("%d", b2i(*m.SyncVerticalAttr))})
	}
	if m.SyncRefAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "syncRef"},
			Value: fmt.Sprintf("%v", *m.SyncRefAttr)})
	}
	if m.TransitionEvaluationAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "transitionEvaluation"},
			Value: fmt.Sprintf("%d", b2i(*m.TransitionEvaluationAttr))})
	}
	if m.TransitionEntryAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "transitionEntry"},
			Value: fmt.Sprintf("%d", b2i(*m.TransitionEntryAttr))})
	}
	if m.PublishedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "published"},
			Value: fmt.Sprintf("%d", b2i(*m.PublishedAttr))})
	}
	if m.CodeNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "codeName"},
			Value: fmt.Sprintf("%v", *m.CodeNameAttr)})
	}
	if m.FilterModeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "filterMode"},
			Value: fmt.Sprintf("%d", b2i(*m.FilterModeAttr))})
	}
	if m.EnableFormatConditionsCalculationAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "enableFormatConditionsCalculation"},
			Value: fmt.Sprintf("%d", b2i(*m.EnableFormatConditionsCalculationAttr))})
	}
	e.EncodeToken(start)
	if m.TabColor != nil {
		setabColor := xml.StartElement{Name: xml.Name{Local: "x:tabColor"}}
		e.EncodeElement(m.TabColor, setabColor)
	}
	if m.OutlinePr != nil {
		seoutlinePr := xml.StartElement{Name: xml.Name{Local: "x:outlinePr"}}
		e.EncodeElement(m.OutlinePr, seoutlinePr)
	}
	if m.PageSetUpPr != nil {
		sepageSetUpPr := xml.StartElement{Name: xml.Name{Local: "x:pageSetUpPr"}}
		e.EncodeElement(m.PageSetUpPr, sepageSetUpPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SheetPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.SyncHorizontalAttr = gooxml.Bool(false)
	m.SyncVerticalAttr = gooxml.Bool(false)
	m.TransitionEvaluationAttr = gooxml.Bool(false)
	m.TransitionEntryAttr = gooxml.Bool(false)
	m.PublishedAttr = gooxml.Bool(true)
	m.FilterModeAttr = gooxml.Bool(false)
	m.EnableFormatConditionsCalculationAttr = gooxml.Bool(true)
	for _, attr := range start.Attr {
		if attr.Name.Local == "syncHorizontal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SyncHorizontalAttr = &parsed
		}
		if attr.Name.Local == "syncVertical" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SyncVerticalAttr = &parsed
		}
		if attr.Name.Local == "syncRef" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SyncRefAttr = &parsed
		}
		if attr.Name.Local == "transitionEvaluation" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.TransitionEvaluationAttr = &parsed
		}
		if attr.Name.Local == "transitionEntry" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.TransitionEntryAttr = &parsed
		}
		if attr.Name.Local == "published" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PublishedAttr = &parsed
		}
		if attr.Name.Local == "codeName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CodeNameAttr = &parsed
		}
		if attr.Name.Local == "filterMode" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FilterModeAttr = &parsed
		}
		if attr.Name.Local == "enableFormatConditionsCalculation" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EnableFormatConditionsCalculationAttr = &parsed
		}
	}
lCT_SheetPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tabColor":
				m.TabColor = NewCT_Color()
				if err := d.DecodeElement(m.TabColor, &el); err != nil {
					return err
				}
			case "outlinePr":
				m.OutlinePr = NewCT_OutlinePr()
				if err := d.DecodeElement(m.OutlinePr, &el); err != nil {
					return err
				}
			case "pageSetUpPr":
				m.PageSetUpPr = NewCT_PageSetUpPr()
				if err := d.DecodeElement(m.PageSetUpPr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_SheetPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SheetPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SheetPr and its children
func (m *CT_SheetPr) Validate() error {
	return m.ValidateWithPath("CT_SheetPr")
}

// ValidateWithPath validates the CT_SheetPr and its children, prefixing error messages with path
func (m *CT_SheetPr) ValidateWithPath(path string) error {
	if m.TabColor != nil {
		if err := m.TabColor.ValidateWithPath(path + "/TabColor"); err != nil {
			return err
		}
	}
	if m.OutlinePr != nil {
		if err := m.OutlinePr.ValidateWithPath(path + "/OutlinePr"); err != nil {
			return err
		}
	}
	if m.PageSetUpPr != nil {
		if err := m.PageSetUpPr.ValidateWithPath(path + "/PageSetUpPr"); err != nil {
			return err
		}
	}
	return nil
}
