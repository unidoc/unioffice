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

type CT_CacheHierarchy struct {
	// Hierarchy Unique Name
	UniqueNameAttr string
	// Hierarchy Display Name
	CaptionAttr *string
	// Measure Hierarchy
	MeasureAttr *bool
	// Set
	SetAttr *bool
	// Parent Set
	ParentSetAttr *uint32
	// KPI Icon Set
	IconSetAttr *int32
	// Attribute Hierarchy
	AttributeAttr *bool
	// Time
	TimeAttr *bool
	// Key Attribute Hierarchy
	KeyAttributeAttr *bool
	// Default Member Unique Name
	DefaultMemberUniqueNameAttr *string
	// Unique Name of 'All'
	AllUniqueNameAttr *string
	// Display Name of 'All'
	AllCaptionAttr *string
	// Dimension Unique Name
	DimensionUniqueNameAttr *string
	// Display Folder
	DisplayFolderAttr *string
	// Measure Group Name
	MeasureGroupAttr *string
	// Measures
	MeasuresAttr *bool
	// Levels Count
	CountAttr uint32
	// One Field
	OneFieldAttr *bool
	// Member Value Data Type
	MemberValueDatatypeAttr *uint16
	// Unbalanced
	UnbalancedAttr *bool
	// Unbalanced Group
	UnbalancedGroupAttr *bool
	// Hidden
	HiddenAttr *bool
	// Fields Usage
	FieldsUsage *CT_FieldsUsage
	// OLAP Grouping Levels
	GroupLevels *CT_GroupLevels
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_CacheHierarchy() *CT_CacheHierarchy {
	ret := &CT_CacheHierarchy{}
	return ret
}
func (m *CT_CacheHierarchy) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uniqueName"},
		Value: fmt.Sprintf("%v", m.UniqueNameAttr)})
	if m.CaptionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "caption"},
			Value: fmt.Sprintf("%v", *m.CaptionAttr)})
	}
	if m.MeasureAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "measure"},
			Value: fmt.Sprintf("%v", *m.MeasureAttr)})
	}
	if m.SetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "set"},
			Value: fmt.Sprintf("%v", *m.SetAttr)})
	}
	if m.ParentSetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "parentSet"},
			Value: fmt.Sprintf("%v", *m.ParentSetAttr)})
	}
	if m.IconSetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "iconSet"},
			Value: fmt.Sprintf("%v", *m.IconSetAttr)})
	}
	if m.AttributeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "attribute"},
			Value: fmt.Sprintf("%v", *m.AttributeAttr)})
	}
	if m.TimeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "time"},
			Value: fmt.Sprintf("%v", *m.TimeAttr)})
	}
	if m.KeyAttributeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "keyAttribute"},
			Value: fmt.Sprintf("%v", *m.KeyAttributeAttr)})
	}
	if m.DefaultMemberUniqueNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "defaultMemberUniqueName"},
			Value: fmt.Sprintf("%v", *m.DefaultMemberUniqueNameAttr)})
	}
	if m.AllUniqueNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "allUniqueName"},
			Value: fmt.Sprintf("%v", *m.AllUniqueNameAttr)})
	}
	if m.AllCaptionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "allCaption"},
			Value: fmt.Sprintf("%v", *m.AllCaptionAttr)})
	}
	if m.DimensionUniqueNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dimensionUniqueName"},
			Value: fmt.Sprintf("%v", *m.DimensionUniqueNameAttr)})
	}
	if m.DisplayFolderAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "displayFolder"},
			Value: fmt.Sprintf("%v", *m.DisplayFolderAttr)})
	}
	if m.MeasureGroupAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "measureGroup"},
			Value: fmt.Sprintf("%v", *m.MeasureGroupAttr)})
	}
	if m.MeasuresAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "measures"},
			Value: fmt.Sprintf("%v", *m.MeasuresAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
		Value: fmt.Sprintf("%v", m.CountAttr)})
	if m.OneFieldAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "oneField"},
			Value: fmt.Sprintf("%v", *m.OneFieldAttr)})
	}
	if m.MemberValueDatatypeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "memberValueDatatype"},
			Value: fmt.Sprintf("%v", *m.MemberValueDatatypeAttr)})
	}
	if m.UnbalancedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "unbalanced"},
			Value: fmt.Sprintf("%v", *m.UnbalancedAttr)})
	}
	if m.UnbalancedGroupAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "unbalancedGroup"},
			Value: fmt.Sprintf("%v", *m.UnbalancedGroupAttr)})
	}
	if m.HiddenAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hidden"},
			Value: fmt.Sprintf("%v", *m.HiddenAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.FieldsUsage != nil {
		sefieldsUsage := xml.StartElement{Name: xml.Name{Local: "x:fieldsUsage"}}
		e.EncodeElement(m.FieldsUsage, sefieldsUsage)
	}
	if m.GroupLevels != nil {
		segroupLevels := xml.StartElement{Name: xml.Name{Local: "x:groupLevels"}}
		e.EncodeElement(m.GroupLevels, segroupLevels)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_CacheHierarchy) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "uniqueName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UniqueNameAttr = parsed
		}
		if attr.Name.Local == "caption" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CaptionAttr = &parsed
		}
		if attr.Name.Local == "measure" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MeasureAttr = &parsed
		}
		if attr.Name.Local == "set" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SetAttr = &parsed
		}
		if attr.Name.Local == "parentSet" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.ParentSetAttr = &pt
		}
		if attr.Name.Local == "iconSet" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := int32(parsed)
			m.IconSetAttr = &pt
		}
		if attr.Name.Local == "attribute" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AttributeAttr = &parsed
		}
		if attr.Name.Local == "time" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.TimeAttr = &parsed
		}
		if attr.Name.Local == "keyAttribute" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.KeyAttributeAttr = &parsed
		}
		if attr.Name.Local == "defaultMemberUniqueName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DefaultMemberUniqueNameAttr = &parsed
		}
		if attr.Name.Local == "allUniqueName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AllUniqueNameAttr = &parsed
		}
		if attr.Name.Local == "allCaption" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AllCaptionAttr = &parsed
		}
		if attr.Name.Local == "dimensionUniqueName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DimensionUniqueNameAttr = &parsed
		}
		if attr.Name.Local == "displayFolder" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DisplayFolderAttr = &parsed
		}
		if attr.Name.Local == "measureGroup" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MeasureGroupAttr = &parsed
		}
		if attr.Name.Local == "measures" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MeasuresAttr = &parsed
		}
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.CountAttr = uint32(parsed)
		}
		if attr.Name.Local == "oneField" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.OneFieldAttr = &parsed
		}
		if attr.Name.Local == "memberValueDatatype" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 16)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint16(parsed)
			m.MemberValueDatatypeAttr = &pt
		}
		if attr.Name.Local == "unbalanced" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UnbalancedAttr = &parsed
		}
		if attr.Name.Local == "unbalancedGroup" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UnbalancedGroupAttr = &parsed
		}
		if attr.Name.Local == "hidden" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HiddenAttr = &parsed
		}
	}
lCT_CacheHierarchy:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "fieldsUsage":
				m.FieldsUsage = NewCT_FieldsUsage()
				if err := d.DecodeElement(m.FieldsUsage, &el); err != nil {
					return err
				}
			case "groupLevels":
				m.GroupLevels = NewCT_GroupLevels()
				if err := d.DecodeElement(m.GroupLevels, &el); err != nil {
					return err
				}
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
			break lCT_CacheHierarchy
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_CacheHierarchy) Validate() error {
	return m.ValidateWithPath("CT_CacheHierarchy")
}
func (m *CT_CacheHierarchy) ValidateWithPath(path string) error {
	if m.FieldsUsage != nil {
		if err := m.FieldsUsage.ValidateWithPath(path + "/FieldsUsage"); err != nil {
			return err
		}
	}
	if m.GroupLevels != nil {
		if err := m.GroupLevels.ValidateWithPath(path + "/GroupLevels"); err != nil {
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
