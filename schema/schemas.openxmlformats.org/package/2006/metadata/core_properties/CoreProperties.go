// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package core_properties

import (
	"encoding/xml"
	"log"
	"time"
)

type CoreProperties struct {
	CT_CoreProperties
}

func NewCoreProperties() *CoreProperties {
	ret := &CoreProperties{}
	ret.CT_CoreProperties = *NewCT_CoreProperties()
	return ret
}
func (m *CoreProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:cp"}, Value: "http://schemas.openxmlformats.org/package/2006/metadata/core-properties"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:dc"}, Value: "http://purl.org/dc/elements/1.1/"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:dcterms"}, Value: "http://purl.org/dc/terms/"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "cp:coreProperties"
	return m.CT_CoreProperties.MarshalXML(e, start)
}
func (m *CoreProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_CoreProperties = *NewCT_CoreProperties()
lCoreProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "category":
				m.Category = new(string)
				if err := d.DecodeElement(m.Category, &el); err != nil {
					return err
				}
			case "contentStatus":
				m.ContentStatus = new(string)
				if err := d.DecodeElement(m.ContentStatus, &el); err != nil {
					return err
				}
			case "created":
				m.Created = new(string)
				if err := d.DecodeElement(m.Created, &el); err != nil {
					return err
				}
			case "creator":
				m.Creator = new(string)
				if err := d.DecodeElement(m.Creator, &el); err != nil {
					return err
				}
			case "description":
				m.Description = new(string)
				if err := d.DecodeElement(m.Description, &el); err != nil {
					return err
				}
			case "identifier":
				m.Identifier = new(string)
				if err := d.DecodeElement(m.Identifier, &el); err != nil {
					return err
				}
			case "keywords":
				m.Keywords = NewCT_Keywords()
				if err := d.DecodeElement(m.Keywords, &el); err != nil {
					return err
				}
			case "language":
				m.Language = new(string)
				if err := d.DecodeElement(m.Language, &el); err != nil {
					return err
				}
			case "lastModifiedBy":
				m.LastModifiedBy = new(string)
				if err := d.DecodeElement(m.LastModifiedBy, &el); err != nil {
					return err
				}
			case "lastPrinted":
				m.LastPrinted = new(time.Time)
				if err := d.DecodeElement(m.LastPrinted, &el); err != nil {
					return err
				}
			case "modified":
				m.Modified = new(string)
				if err := d.DecodeElement(m.Modified, &el); err != nil {
					return err
				}
			case "revision":
				m.Revision = new(string)
				if err := d.DecodeElement(m.Revision, &el); err != nil {
					return err
				}
			case "subject":
				m.Subject = new(string)
				if err := d.DecodeElement(m.Subject, &el); err != nil {
					return err
				}
			case "title":
				m.Title = new(string)
				if err := d.DecodeElement(m.Title, &el); err != nil {
					return err
				}
			case "version":
				m.Version = new(string)
				if err := d.DecodeElement(m.Version, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCoreProperties
		case xml.CharData:
		}
	}
	return nil
}
func (m *CoreProperties) Validate() error {
	return m.ValidateWithPath("CoreProperties")
}
func (m *CoreProperties) ValidateWithPath(path string) error {
	if err := m.CT_CoreProperties.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
