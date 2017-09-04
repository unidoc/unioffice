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

	"baliance.com/gooxml"
)

type CT_HtmlPublishProperties struct {
	// Show Speaker Notes
	ShowSpeakerNotesAttr *bool
	// Target Output Profile
	TargetAttr *string
	// HTML Output Title
	TitleAttr *string
	IdAttr    string
	// All Slides
	SldAll *CT_Empty
	// Slide Range
	SldRg *CT_IndexRange
	// Custom Show
	CustShow *CT_CustomShowId
	ExtLst   *CT_ExtensionList
}

func NewCT_HtmlPublishProperties() *CT_HtmlPublishProperties {
	ret := &CT_HtmlPublishProperties{}
	ret.ShowSpeakerNotesAttr = gooxml.Bool(true)
	return ret
}

func (m *CT_HtmlPublishProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ShowSpeakerNotesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showSpeakerNotes"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowSpeakerNotesAttr))})
	}
	if m.TargetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "target"},
			Value: fmt.Sprintf("%v", *m.TargetAttr)})
	}
	if m.TitleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "title"},
			Value: fmt.Sprintf("%v", *m.TitleAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	if m.SldAll != nil {
		sesldAll := xml.StartElement{Name: xml.Name{Local: "p:sldAll"}}
		e.EncodeElement(m.SldAll, sesldAll)
	}
	if m.SldRg != nil {
		sesldRg := xml.StartElement{Name: xml.Name{Local: "p:sldRg"}}
		e.EncodeElement(m.SldRg, sesldRg)
	}
	if m.CustShow != nil {
		secustShow := xml.StartElement{Name: xml.Name{Local: "p:custShow"}}
		e.EncodeElement(m.CustShow, secustShow)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_HtmlPublishProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ShowSpeakerNotesAttr = gooxml.Bool(true)
	for _, attr := range start.Attr {
		if attr.Name.Local == "showSpeakerNotes" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowSpeakerNotesAttr = &parsed
		}
		if attr.Name.Local == "target" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TargetAttr = &parsed
		}
		if attr.Name.Local == "title" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TitleAttr = &parsed
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
		}
	}
lCT_HtmlPublishProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "sldAll":
				m.SldAll = NewCT_Empty()
				if err := d.DecodeElement(m.SldAll, &el); err != nil {
					return err
				}
			case "sldRg":
				m.SldRg = NewCT_IndexRange()
				if err := d.DecodeElement(m.SldRg, &el); err != nil {
					return err
				}
			case "custShow":
				m.CustShow = NewCT_CustomShowId()
				if err := d.DecodeElement(m.CustShow, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_HtmlPublishProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_HtmlPublishProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_HtmlPublishProperties and its children
func (m *CT_HtmlPublishProperties) Validate() error {
	return m.ValidateWithPath("CT_HtmlPublishProperties")
}

// ValidateWithPath validates the CT_HtmlPublishProperties and its children, prefixing error messages with path
func (m *CT_HtmlPublishProperties) ValidateWithPath(path string) error {
	if m.SldAll != nil {
		if err := m.SldAll.ValidateWithPath(path + "/SldAll"); err != nil {
			return err
		}
	}
	if m.SldRg != nil {
		if err := m.SldRg.ValidateWithPath(path + "/SldRg"); err != nil {
			return err
		}
	}
	if m.CustShow != nil {
		if err := m.CustShow.ValidateWithPath(path + "/CustShow"); err != nil {
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
