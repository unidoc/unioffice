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
)

type CT_NormalViewProperties struct {
	// Show Outline Icons in Normal View
	ShowOutlineIconsAttr *bool
	// Snap Vertical Splitter
	SnapVertSplitterAttr *bool
	// State of the Vertical Splitter Bar
	VertBarStateAttr ST_SplitterBarState
	// State of the Horizontal Splitter Bar
	HorzBarStateAttr ST_SplitterBarState
	// Prefer Single View
	PreferSingleViewAttr *bool
	// Normal View Restored Left Properties
	RestoredLeft *CT_NormalViewPortion
	// Normal View Restored Top Properties
	RestoredTop *CT_NormalViewPortion
	ExtLst      *CT_ExtensionList
}

func NewCT_NormalViewProperties() *CT_NormalViewProperties {
	ret := &CT_NormalViewProperties{}
	ret.RestoredLeft = NewCT_NormalViewPortion()
	ret.RestoredTop = NewCT_NormalViewPortion()
	return ret
}
func (m *CT_NormalViewProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.ShowOutlineIconsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showOutlineIcons"},
			Value: fmt.Sprintf("%v", *m.ShowOutlineIconsAttr)})
	}
	if m.SnapVertSplitterAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "snapVertSplitter"},
			Value: fmt.Sprintf("%v", *m.SnapVertSplitterAttr)})
	}
	if m.VertBarStateAttr != ST_SplitterBarStateUnset {
		attr, err := m.VertBarStateAttr.MarshalXMLAttr(xml.Name{Local: "vertBarState"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HorzBarStateAttr != ST_SplitterBarStateUnset {
		attr, err := m.HorzBarStateAttr.MarshalXMLAttr(xml.Name{Local: "horzBarState"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.PreferSingleViewAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "preferSingleView"},
			Value: fmt.Sprintf("%v", *m.PreferSingleViewAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	serestoredLeft := xml.StartElement{Name: xml.Name{Local: "p:restoredLeft"}}
	e.EncodeElement(m.RestoredLeft, serestoredLeft)
	serestoredTop := xml.StartElement{Name: xml.Name{Local: "p:restoredTop"}}
	e.EncodeElement(m.RestoredTop, serestoredTop)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_NormalViewProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.RestoredLeft = NewCT_NormalViewPortion()
	m.RestoredTop = NewCT_NormalViewPortion()
	for _, attr := range start.Attr {
		if attr.Name.Local == "showOutlineIcons" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowOutlineIconsAttr = &parsed
		}
		if attr.Name.Local == "snapVertSplitter" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SnapVertSplitterAttr = &parsed
		}
		if attr.Name.Local == "vertBarState" {
			m.VertBarStateAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "horzBarState" {
			m.HorzBarStateAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "preferSingleView" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PreferSingleViewAttr = &parsed
		}
	}
lCT_NormalViewProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "restoredLeft":
				if err := d.DecodeElement(m.RestoredLeft, &el); err != nil {
					return err
				}
			case "restoredTop":
				if err := d.DecodeElement(m.RestoredTop, &el); err != nil {
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
			break lCT_NormalViewProperties
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_NormalViewProperties) Validate() error {
	return m.ValidateWithPath("CT_NormalViewProperties")
}
func (m *CT_NormalViewProperties) ValidateWithPath(path string) error {
	if err := m.VertBarStateAttr.ValidateWithPath(path + "/VertBarStateAttr"); err != nil {
		return err
	}
	if err := m.HorzBarStateAttr.ValidateWithPath(path + "/HorzBarStateAttr"); err != nil {
		return err
	}
	if err := m.RestoredLeft.ValidateWithPath(path + "/RestoredLeft"); err != nil {
		return err
	}
	if err := m.RestoredTop.ValidateWithPath(path + "/RestoredTop"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
