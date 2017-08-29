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

type CT_PivotAreaReference struct {
	// Field Index
	FieldAttr *uint32
	// Item Index Count
	CountAttr *uint32
	// Selected
	SelectedAttr *bool
	// Positional Reference
	ByPositionAttr *bool
	// Relative Reference
	RelativeAttr *bool
	// Include Default Filter
	DefaultSubtotalAttr *bool
	// Include Sum Filter
	SumSubtotalAttr *bool
	// Include CountA Filter
	CountASubtotalAttr *bool
	// Include Average Filter
	AvgSubtotalAttr *bool
	// Include Maximum Filter
	MaxSubtotalAttr *bool
	// Include Minimum Filter
	MinSubtotalAttr *bool
	// Include Product Filter
	ProductSubtotalAttr *bool
	// Include Count Subtotal
	CountSubtotalAttr *bool
	// Include StdDev Filter
	StdDevSubtotalAttr *bool
	// Include StdDevP Filter
	StdDevPSubtotalAttr *bool
	// Include Var Filter
	VarSubtotalAttr *bool
	// Include VarP Filter
	VarPSubtotalAttr *bool
	// Field Item
	X      []*CT_Index
	ExtLst *CT_ExtensionList
}

func NewCT_PivotAreaReference() *CT_PivotAreaReference {
	ret := &CT_PivotAreaReference{}
	return ret
}
func (m *CT_PivotAreaReference) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.FieldAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "field"},
			Value: fmt.Sprintf("%v", *m.FieldAttr)})
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	if m.SelectedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "selected"},
			Value: fmt.Sprintf("%v", *m.SelectedAttr)})
	}
	if m.ByPositionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "byPosition"},
			Value: fmt.Sprintf("%v", *m.ByPositionAttr)})
	}
	if m.RelativeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "relative"},
			Value: fmt.Sprintf("%v", *m.RelativeAttr)})
	}
	if m.DefaultSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "defaultSubtotal"},
			Value: fmt.Sprintf("%v", *m.DefaultSubtotalAttr)})
	}
	if m.SumSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sumSubtotal"},
			Value: fmt.Sprintf("%v", *m.SumSubtotalAttr)})
	}
	if m.CountASubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "countASubtotal"},
			Value: fmt.Sprintf("%v", *m.CountASubtotalAttr)})
	}
	if m.AvgSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "avgSubtotal"},
			Value: fmt.Sprintf("%v", *m.AvgSubtotalAttr)})
	}
	if m.MaxSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "maxSubtotal"},
			Value: fmt.Sprintf("%v", *m.MaxSubtotalAttr)})
	}
	if m.MinSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minSubtotal"},
			Value: fmt.Sprintf("%v", *m.MinSubtotalAttr)})
	}
	if m.ProductSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "productSubtotal"},
			Value: fmt.Sprintf("%v", *m.ProductSubtotalAttr)})
	}
	if m.CountSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "countSubtotal"},
			Value: fmt.Sprintf("%v", *m.CountSubtotalAttr)})
	}
	if m.StdDevSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "stdDevSubtotal"},
			Value: fmt.Sprintf("%v", *m.StdDevSubtotalAttr)})
	}
	if m.StdDevPSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "stdDevPSubtotal"},
			Value: fmt.Sprintf("%v", *m.StdDevPSubtotalAttr)})
	}
	if m.VarSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "varSubtotal"},
			Value: fmt.Sprintf("%v", *m.VarSubtotalAttr)})
	}
	if m.VarPSubtotalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "varPSubtotal"},
			Value: fmt.Sprintf("%v", *m.VarPSubtotalAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.X != nil {
		sex := xml.StartElement{Name: xml.Name{Local: "x:x"}}
		e.EncodeElement(m.X, sex)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_PivotAreaReference) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "field" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.FieldAttr = &pt
		}
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.CountAttr = &pt
		}
		if attr.Name.Local == "selected" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SelectedAttr = &parsed
		}
		if attr.Name.Local == "byPosition" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ByPositionAttr = &parsed
		}
		if attr.Name.Local == "relative" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RelativeAttr = &parsed
		}
		if attr.Name.Local == "defaultSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DefaultSubtotalAttr = &parsed
		}
		if attr.Name.Local == "sumSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SumSubtotalAttr = &parsed
		}
		if attr.Name.Local == "countASubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CountASubtotalAttr = &parsed
		}
		if attr.Name.Local == "avgSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AvgSubtotalAttr = &parsed
		}
		if attr.Name.Local == "maxSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MaxSubtotalAttr = &parsed
		}
		if attr.Name.Local == "minSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MinSubtotalAttr = &parsed
		}
		if attr.Name.Local == "productSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ProductSubtotalAttr = &parsed
		}
		if attr.Name.Local == "countSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CountSubtotalAttr = &parsed
		}
		if attr.Name.Local == "stdDevSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.StdDevSubtotalAttr = &parsed
		}
		if attr.Name.Local == "stdDevPSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.StdDevPSubtotalAttr = &parsed
		}
		if attr.Name.Local == "varSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.VarSubtotalAttr = &parsed
		}
		if attr.Name.Local == "varPSubtotal" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.VarPSubtotalAttr = &parsed
		}
	}
lCT_PivotAreaReference:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "x":
				tmp := NewCT_Index()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.X = append(m.X, tmp)
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
			break lCT_PivotAreaReference
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_PivotAreaReference) Validate() error {
	return m.ValidateWithPath("CT_PivotAreaReference")
}
func (m *CT_PivotAreaReference) ValidateWithPath(path string) error {
	for i, v := range m.X {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/X[%d]", path, i)); err != nil {
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
