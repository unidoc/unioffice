// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"log"
	"strconv"
)

type SldMaster struct {
	CT_SlideMaster
}

func NewSldMaster() *SldMaster {
	ret := &SldMaster{}
	ret.CT_SlideMaster = *NewCT_SlideMaster()
	return ret
}
func (m *SldMaster) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/presentationml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:p"}, Value: "http://schemas.openxmlformats.org/presentationml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "p:sldMaster"
	return m.CT_SlideMaster.MarshalXML(e, start)
}
func (m *SldMaster) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_SlideMaster = *NewCT_SlideMaster()
	for _, attr := range start.Attr {
		if attr.Name.Local == "preserve" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PreserveAttr = &parsed
		}
	}
lSldMaster:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cSld":
				if err := d.DecodeElement(m.CSld, &el); err != nil {
					return err
				}
			case "clrMap":
				if err := d.DecodeElement(m.ClrMap, &el); err != nil {
					return err
				}
			case "sldLayoutIdLst":
				m.SldLayoutIdLst = NewCT_SlideLayoutIdList()
				if err := d.DecodeElement(m.SldLayoutIdLst, &el); err != nil {
					return err
				}
			case "transition":
				m.Transition = NewCT_SlideTransition()
				if err := d.DecodeElement(m.Transition, &el); err != nil {
					return err
				}
			case "timing":
				m.Timing = NewCT_SlideTiming()
				if err := d.DecodeElement(m.Timing, &el); err != nil {
					return err
				}
			case "hf":
				m.Hf = NewCT_HeaderFooter()
				if err := d.DecodeElement(m.Hf, &el); err != nil {
					return err
				}
			case "txStyles":
				m.TxStyles = NewCT_SlideMasterTextStyles()
				if err := d.DecodeElement(m.TxStyles, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionListModify()
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
			break lSldMaster
		case xml.CharData:
		}
	}
	return nil
}
func (m *SldMaster) Validate() error {
	return m.ValidateWithPath("SldMaster")
}
func (m *SldMaster) ValidateWithPath(path string) error {
	if err := m.CT_SlideMaster.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
