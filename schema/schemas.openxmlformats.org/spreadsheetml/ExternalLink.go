// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"log"
)

type ExternalLink struct {
	CT_ExternalLink
}

func NewExternalLink() *ExternalLink {
	ret := &ExternalLink{}
	ret.CT_ExternalLink = *NewCT_ExternalLink()
	return ret
}

func (m *ExternalLink) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:x"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "x:externalLink"
	return m.CT_ExternalLink.MarshalXML(e, start)
}

func (m *ExternalLink) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_ExternalLink = *NewCT_ExternalLink()
lExternalLink:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "externalBook":
				m.Choice = NewCT_ExternalLinkChoice()
				if err := d.DecodeElement(&m.Choice.ExternalBook, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "ddeLink":
				m.Choice = NewCT_ExternalLinkChoice()
				if err := d.DecodeElement(&m.Choice.DdeLink, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "oleLink":
				m.Choice = NewCT_ExternalLinkChoice()
				if err := d.DecodeElement(&m.Choice.OleLink, &el); err != nil {
					return err
				}
				_ = m.Choice
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on ExternalLink %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lExternalLink
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the ExternalLink and its children
func (m *ExternalLink) Validate() error {
	return m.ValidateWithPath("ExternalLink")
}

// ValidateWithPath validates the ExternalLink and its children, prefixing error messages with path
func (m *ExternalLink) ValidateWithPath(path string) error {
	if err := m.CT_ExternalLink.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
