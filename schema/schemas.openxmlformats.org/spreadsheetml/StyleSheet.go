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

type StyleSheet struct {
	CT_Stylesheet
}

func NewStyleSheet() *StyleSheet {
	ret := &StyleSheet{}
	ret.CT_Stylesheet = *NewCT_Stylesheet()
	return ret
}

func (m *StyleSheet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:x"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "x:styleSheet"
	return m.CT_Stylesheet.MarshalXML(e, start)
}

func (m *StyleSheet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Stylesheet = *NewCT_Stylesheet()
lStyleSheet:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "numFmts":
				m.NumFmts = NewCT_NumFmts()
				if err := d.DecodeElement(m.NumFmts, &el); err != nil {
					return err
				}
			case "fonts":
				m.Fonts = NewCT_Fonts()
				if err := d.DecodeElement(m.Fonts, &el); err != nil {
					return err
				}
			case "fills":
				m.Fills = NewCT_Fills()
				if err := d.DecodeElement(m.Fills, &el); err != nil {
					return err
				}
			case "borders":
				m.Borders = NewCT_Borders()
				if err := d.DecodeElement(m.Borders, &el); err != nil {
					return err
				}
			case "cellStyleXfs":
				m.CellStyleXfs = NewCT_CellStyleXfs()
				if err := d.DecodeElement(m.CellStyleXfs, &el); err != nil {
					return err
				}
			case "cellXfs":
				m.CellXfs = NewCT_CellXfs()
				if err := d.DecodeElement(m.CellXfs, &el); err != nil {
					return err
				}
			case "cellStyles":
				m.CellStyles = NewCT_CellStyles()
				if err := d.DecodeElement(m.CellStyles, &el); err != nil {
					return err
				}
			case "dxfs":
				m.Dxfs = NewCT_Dxfs()
				if err := d.DecodeElement(m.Dxfs, &el); err != nil {
					return err
				}
			case "tableStyles":
				m.TableStyles = NewCT_TableStyles()
				if err := d.DecodeElement(m.TableStyles, &el); err != nil {
					return err
				}
			case "colors":
				m.Colors = NewCT_Colors()
				if err := d.DecodeElement(m.Colors, &el); err != nil {
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
			break lStyleSheet
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the StyleSheet and its children
func (m *StyleSheet) Validate() error {
	return m.ValidateWithPath("StyleSheet")
}

// ValidateWithPath validates the StyleSheet and its children, prefixing error messages with path
func (m *StyleSheet) ValidateWithPath(path string) error {
	if err := m.CT_Stylesheet.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
