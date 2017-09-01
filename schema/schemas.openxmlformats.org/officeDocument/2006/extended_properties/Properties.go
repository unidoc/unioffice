// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package extended_properties

import (
	"encoding/xml"
	"log"
)

type Properties struct {
	CT_Properties
}

func NewProperties() *Properties {
	ret := &Properties{}
	ret.CT_Properties = *NewCT_Properties()
	return ret
}
func (m *Properties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/extended-properties"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:vt"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "Properties"
	return m.CT_Properties.MarshalXML(e, start)
}
func (m *Properties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Properties = *NewCT_Properties()
lProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "Template":
				m.Template = new(string)
				if err := d.DecodeElement(m.Template, &el); err != nil {
					return err
				}
			case "Manager":
				m.Manager = new(string)
				if err := d.DecodeElement(m.Manager, &el); err != nil {
					return err
				}
			case "Company":
				m.Company = new(string)
				if err := d.DecodeElement(m.Company, &el); err != nil {
					return err
				}
			case "Pages":
				m.Pages = new(int32)
				if err := d.DecodeElement(m.Pages, &el); err != nil {
					return err
				}
			case "Words":
				m.Words = new(int32)
				if err := d.DecodeElement(m.Words, &el); err != nil {
					return err
				}
			case "Characters":
				m.Characters = new(int32)
				if err := d.DecodeElement(m.Characters, &el); err != nil {
					return err
				}
			case "PresentationFormat":
				m.PresentationFormat = new(string)
				if err := d.DecodeElement(m.PresentationFormat, &el); err != nil {
					return err
				}
			case "Lines":
				m.Lines = new(int32)
				if err := d.DecodeElement(m.Lines, &el); err != nil {
					return err
				}
			case "Paragraphs":
				m.Paragraphs = new(int32)
				if err := d.DecodeElement(m.Paragraphs, &el); err != nil {
					return err
				}
			case "Slides":
				m.Slides = new(int32)
				if err := d.DecodeElement(m.Slides, &el); err != nil {
					return err
				}
			case "Notes":
				m.Notes = new(int32)
				if err := d.DecodeElement(m.Notes, &el); err != nil {
					return err
				}
			case "TotalTime":
				m.TotalTime = new(int32)
				if err := d.DecodeElement(m.TotalTime, &el); err != nil {
					return err
				}
			case "HiddenSlides":
				m.HiddenSlides = new(int32)
				if err := d.DecodeElement(m.HiddenSlides, &el); err != nil {
					return err
				}
			case "MMClips":
				m.MMClips = new(int32)
				if err := d.DecodeElement(m.MMClips, &el); err != nil {
					return err
				}
			case "ScaleCrop":
				m.ScaleCrop = new(bool)
				if err := d.DecodeElement(m.ScaleCrop, &el); err != nil {
					return err
				}
			case "HeadingPairs":
				m.HeadingPairs = NewCT_VectorVariant()
				if err := d.DecodeElement(m.HeadingPairs, &el); err != nil {
					return err
				}
			case "TitlesOfParts":
				m.TitlesOfParts = NewCT_VectorLpstr()
				if err := d.DecodeElement(m.TitlesOfParts, &el); err != nil {
					return err
				}
			case "LinksUpToDate":
				m.LinksUpToDate = new(bool)
				if err := d.DecodeElement(m.LinksUpToDate, &el); err != nil {
					return err
				}
			case "CharactersWithSpaces":
				m.CharactersWithSpaces = new(int32)
				if err := d.DecodeElement(m.CharactersWithSpaces, &el); err != nil {
					return err
				}
			case "SharedDoc":
				m.SharedDoc = new(bool)
				if err := d.DecodeElement(m.SharedDoc, &el); err != nil {
					return err
				}
			case "HyperlinkBase":
				m.HyperlinkBase = new(string)
				if err := d.DecodeElement(m.HyperlinkBase, &el); err != nil {
					return err
				}
			case "HLinks":
				m.HLinks = NewCT_VectorVariant()
				if err := d.DecodeElement(m.HLinks, &el); err != nil {
					return err
				}
			case "HyperlinksChanged":
				m.HyperlinksChanged = new(bool)
				if err := d.DecodeElement(m.HyperlinksChanged, &el); err != nil {
					return err
				}
			case "DigSig":
				m.DigSig = NewCT_DigSigBlob()
				if err := d.DecodeElement(m.DigSig, &el); err != nil {
					return err
				}
			case "Application":
				m.Application = new(string)
				if err := d.DecodeElement(m.Application, &el); err != nil {
					return err
				}
			case "AppVersion":
				m.AppVersion = new(string)
				if err := d.DecodeElement(m.AppVersion, &el); err != nil {
					return err
				}
			case "DocSecurity":
				m.DocSecurity = new(int32)
				if err := d.DecodeElement(m.DocSecurity, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lProperties
		case xml.CharData:
		}
	}
	return nil
}
func (m *Properties) Validate() error {
	return m.ValidateWithPath("Properties")
}
func (m *Properties) ValidateWithPath(path string) error {
	if err := m.CT_Properties.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
