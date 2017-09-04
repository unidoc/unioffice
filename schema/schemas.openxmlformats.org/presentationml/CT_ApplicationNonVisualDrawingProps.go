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

type CT_ApplicationNonVisualDrawingProps struct {
	// Is a Photo Album
	IsPhotoAttr *bool
	// Is User Drawn
	UserDrawnAttr *bool
	// Placeholder Shape
	Ph *CT_Placeholder
	// Customer Data List
	CustDataLst *CT_CustomerDataList
	ExtLst      *CT_ExtensionList
}

func NewCT_ApplicationNonVisualDrawingProps() *CT_ApplicationNonVisualDrawingProps {
	ret := &CT_ApplicationNonVisualDrawingProps{}
	return ret
}

func (m *CT_ApplicationNonVisualDrawingProps) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.IsPhotoAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "isPhoto"},
			Value: fmt.Sprintf("%d", b2i(*m.IsPhotoAttr))})
	}
	if m.UserDrawnAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "userDrawn"},
			Value: fmt.Sprintf("%d", b2i(*m.UserDrawnAttr))})
	}
	e.EncodeToken(start)
	if m.Ph != nil {
		seph := xml.StartElement{Name: xml.Name{Local: "p:ph"}}
		e.EncodeElement(m.Ph, seph)
	}
	if m.CustDataLst != nil {
		secustDataLst := xml.StartElement{Name: xml.Name{Local: "p:custDataLst"}}
		e.EncodeElement(m.CustDataLst, secustDataLst)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ApplicationNonVisualDrawingProps) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "isPhoto" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.IsPhotoAttr = &parsed
		}
		if attr.Name.Local == "userDrawn" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UserDrawnAttr = &parsed
		}
	}
lCT_ApplicationNonVisualDrawingProps:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "ph":
				m.Ph = NewCT_Placeholder()
				if err := d.DecodeElement(m.Ph, &el); err != nil {
					return err
				}
			case "custDataLst":
				m.CustDataLst = NewCT_CustomerDataList()
				if err := d.DecodeElement(m.CustDataLst, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_ApplicationNonVisualDrawingProps %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ApplicationNonVisualDrawingProps
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ApplicationNonVisualDrawingProps and its children
func (m *CT_ApplicationNonVisualDrawingProps) Validate() error {
	return m.ValidateWithPath("CT_ApplicationNonVisualDrawingProps")
}

// ValidateWithPath validates the CT_ApplicationNonVisualDrawingProps and its children, prefixing error messages with path
func (m *CT_ApplicationNonVisualDrawingProps) ValidateWithPath(path string) error {
	if m.Ph != nil {
		if err := m.Ph.ValidateWithPath(path + "/Ph"); err != nil {
			return err
		}
	}
	if m.CustDataLst != nil {
		if err := m.CustDataLst.ValidateWithPath(path + "/CustDataLst"); err != nil {
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
