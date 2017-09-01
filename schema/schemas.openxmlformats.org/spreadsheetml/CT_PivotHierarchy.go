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
)

type CT_PivotHierarchy struct {
	// Outline New Levels
	OutlineAttr *bool
	// Multiple Field Filters
	MultipleItemSelectionAllowedAttr *bool
	// New Levels Subtotals At Top
	SubtotalTopAttr *bool
	// Show In Field List
	ShowInFieldListAttr *bool
	// Drag To Row
	DragToRowAttr *bool
	// Drag To Column
	DragToColAttr *bool
	// Drag to Page
	DragToPageAttr *bool
	// Drag To Data
	DragToDataAttr *bool
	// Drag Off
	DragOffAttr *bool
	// Inclusive Manual Filter
	IncludeNewItemsInFilterAttr *bool
	// Hierarchy Caption
	CaptionAttr *string
	// OLAP Member Properties
	Mps *CT_MemberProperties
	// Members
	Members []*CT_Members
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_PivotHierarchy() *CT_PivotHierarchy {
	ret := &CT_PivotHierarchy{}
	return ret
}
func (m *CT_PivotHierarchy) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.OutlineAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "outline"},
			Value: fmt.Sprintf("%v", *m.OutlineAttr)})
	}
	if m.MultipleItemSelectionAllowedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "multipleItemSelectionAllowed"},
			Value: fmt.Sprintf("%v", *m.MultipleItemSelectionAllowedAttr)})
	}
	if m.SubtotalTopAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "subtotalTop"},
			Value: fmt.Sprintf("%v", *m.SubtotalTopAttr)})
	}
	if m.ShowInFieldListAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showInFieldList"},
			Value: fmt.Sprintf("%v", *m.ShowInFieldListAttr)})
	}
	if m.DragToRowAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dragToRow"},
			Value: fmt.Sprintf("%v", *m.DragToRowAttr)})
	}
	if m.DragToColAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dragToCol"},
			Value: fmt.Sprintf("%v", *m.DragToColAttr)})
	}
	if m.DragToPageAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dragToPage"},
			Value: fmt.Sprintf("%v", *m.DragToPageAttr)})
	}
	if m.DragToDataAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dragToData"},
			Value: fmt.Sprintf("%v", *m.DragToDataAttr)})
	}
	if m.DragOffAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dragOff"},
			Value: fmt.Sprintf("%v", *m.DragOffAttr)})
	}
	if m.IncludeNewItemsInFilterAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "includeNewItemsInFilter"},
			Value: fmt.Sprintf("%v", *m.IncludeNewItemsInFilterAttr)})
	}
	if m.CaptionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "caption"},
			Value: fmt.Sprintf("%v", *m.CaptionAttr)})
	}
	e.EncodeToken(start)
	if m.Mps != nil {
		semps := xml.StartElement{Name: xml.Name{Local: "x:mps"}}
		e.EncodeElement(m.Mps, semps)
	}
	if m.Members != nil {
		semembers := xml.StartElement{Name: xml.Name{Local: "x:members"}}
		e.EncodeElement(m.Members, semembers)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_PivotHierarchy) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "outline" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.OutlineAttr = &parsed
		}
		if attr.Name.Local == "multipleItemSelectionAllowed" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MultipleItemSelectionAllowedAttr = &parsed
		}
		if attr.Name.Local == "subtotalTop" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SubtotalTopAttr = &parsed
		}
		if attr.Name.Local == "showInFieldList" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowInFieldListAttr = &parsed
		}
		if attr.Name.Local == "dragToRow" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DragToRowAttr = &parsed
		}
		if attr.Name.Local == "dragToCol" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DragToColAttr = &parsed
		}
		if attr.Name.Local == "dragToPage" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DragToPageAttr = &parsed
		}
		if attr.Name.Local == "dragToData" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DragToDataAttr = &parsed
		}
		if attr.Name.Local == "dragOff" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DragOffAttr = &parsed
		}
		if attr.Name.Local == "includeNewItemsInFilter" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.IncludeNewItemsInFilterAttr = &parsed
		}
		if attr.Name.Local == "caption" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CaptionAttr = &parsed
		}
	}
lCT_PivotHierarchy:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "mps":
				m.Mps = NewCT_MemberProperties()
				if err := d.DecodeElement(m.Mps, &el); err != nil {
					return err
				}
			case "members":
				tmp := NewCT_Members()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Members = append(m.Members, tmp)
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
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
			break lCT_PivotHierarchy
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_PivotHierarchy) Validate() error {
	return m.ValidateWithPath("CT_PivotHierarchy")
}
func (m *CT_PivotHierarchy) ValidateWithPath(path string) error {
	if m.Mps != nil {
		if err := m.Mps.ValidateWithPath(path + "/Mps"); err != nil {
			return err
		}
	}
	for i, v := range m.Members {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Members[%d]", path, i)); err != nil {
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
