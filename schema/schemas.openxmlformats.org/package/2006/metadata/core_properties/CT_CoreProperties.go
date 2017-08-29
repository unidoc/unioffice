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

	"baliance.com/gooxml"
)

type CT_CoreProperties struct {
	Category       *string
	ContentStatus  *string
	Created        *string
	Creator        *string
	Description    *string
	Identifier     *string
	Keywords       *CT_Keywords
	Language       *string
	LastModifiedBy *string
	LastPrinted    *time.Time
	Modified       *string
	Revision       *string
	Subject        *string
	Title          *string
	Version        *string
}

func NewCT_CoreProperties() *CT_CoreProperties {
	ret := &CT_CoreProperties{}
	return ret
}
func (m *CT_CoreProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Category != nil {
		secategory := xml.StartElement{Name: xml.Name{Local: "category"}}
		gooxml.AddPreserveSpaceAttr(&secategory, *m.Category)
		e.EncodeElement(m.Category, secategory)
	}
	if m.ContentStatus != nil {
		secontentStatus := xml.StartElement{Name: xml.Name{Local: "contentStatus"}}
		gooxml.AddPreserveSpaceAttr(&secontentStatus, *m.ContentStatus)
		e.EncodeElement(m.ContentStatus, secontentStatus)
	}
	if m.Created != nil {
		secreated := xml.StartElement{Name: xml.Name{Local: "dcterms:created"}}
		gooxml.AddPreserveSpaceAttr(&secreated, *m.Created)
		e.EncodeElement(m.Created, secreated)
	}
	if m.Creator != nil {
		secreator := xml.StartElement{Name: xml.Name{Local: "dc:creator"}}
		gooxml.AddPreserveSpaceAttr(&secreator, *m.Creator)
		e.EncodeElement(m.Creator, secreator)
	}
	if m.Description != nil {
		sedescription := xml.StartElement{Name: xml.Name{Local: "dc:description"}}
		gooxml.AddPreserveSpaceAttr(&sedescription, *m.Description)
		e.EncodeElement(m.Description, sedescription)
	}
	if m.Identifier != nil {
		seidentifier := xml.StartElement{Name: xml.Name{Local: "dc:identifier"}}
		gooxml.AddPreserveSpaceAttr(&seidentifier, *m.Identifier)
		e.EncodeElement(m.Identifier, seidentifier)
	}
	if m.Keywords != nil {
		sekeywords := xml.StartElement{Name: xml.Name{Local: "keywords"}}
		e.EncodeElement(m.Keywords, sekeywords)
	}
	if m.Language != nil {
		selanguage := xml.StartElement{Name: xml.Name{Local: "dc:language"}}
		gooxml.AddPreserveSpaceAttr(&selanguage, *m.Language)
		e.EncodeElement(m.Language, selanguage)
	}
	if m.LastModifiedBy != nil {
		selastModifiedBy := xml.StartElement{Name: xml.Name{Local: "lastModifiedBy"}}
		gooxml.AddPreserveSpaceAttr(&selastModifiedBy, *m.LastModifiedBy)
		e.EncodeElement(m.LastModifiedBy, selastModifiedBy)
	}
	if m.LastPrinted != nil {
		selastPrinted := xml.StartElement{Name: xml.Name{Local: "lastPrinted"}}
		e.EncodeElement(m.LastPrinted, selastPrinted)
	}
	if m.Modified != nil {
		semodified := xml.StartElement{Name: xml.Name{Local: "dcterms:modified"}}
		gooxml.AddPreserveSpaceAttr(&semodified, *m.Modified)
		e.EncodeElement(m.Modified, semodified)
	}
	if m.Revision != nil {
		serevision := xml.StartElement{Name: xml.Name{Local: "revision"}}
		gooxml.AddPreserveSpaceAttr(&serevision, *m.Revision)
		e.EncodeElement(m.Revision, serevision)
	}
	if m.Subject != nil {
		sesubject := xml.StartElement{Name: xml.Name{Local: "dc:subject"}}
		gooxml.AddPreserveSpaceAttr(&sesubject, *m.Subject)
		e.EncodeElement(m.Subject, sesubject)
	}
	if m.Title != nil {
		setitle := xml.StartElement{Name: xml.Name{Local: "dc:title"}}
		gooxml.AddPreserveSpaceAttr(&setitle, *m.Title)
		e.EncodeElement(m.Title, setitle)
	}
	if m.Version != nil {
		seversion := xml.StartElement{Name: xml.Name{Local: "version"}}
		gooxml.AddPreserveSpaceAttr(&seversion, *m.Version)
		e.EncodeElement(m.Version, seversion)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_CoreProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_CoreProperties:
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
			break lCT_CoreProperties
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_CoreProperties) Validate() error {
	return m.ValidateWithPath("CT_CoreProperties")
}
func (m *CT_CoreProperties) ValidateWithPath(path string) error {
	if m.Keywords != nil {
		if err := m.Keywords.ValidateWithPath(path + "/Keywords"); err != nil {
			return err
		}
	}
	return nil
}
