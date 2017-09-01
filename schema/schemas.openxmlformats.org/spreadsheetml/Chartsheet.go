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

type Chartsheet struct {
	CT_Chartsheet
}

func NewChartsheet() *Chartsheet {
	ret := &Chartsheet{}
	ret.CT_Chartsheet = *NewCT_Chartsheet()
	return ret
}

func (m *Chartsheet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:x"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "x:chartsheet"
	return m.CT_Chartsheet.MarshalXML(e, start)
}

func (m *Chartsheet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Chartsheet = *NewCT_Chartsheet()
lChartsheet:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "sheetPr":
				m.SheetPr = NewCT_ChartsheetPr()
				if err := d.DecodeElement(m.SheetPr, &el); err != nil {
					return err
				}
			case "sheetViews":
				if err := d.DecodeElement(m.SheetViews, &el); err != nil {
					return err
				}
			case "sheetProtection":
				m.SheetProtection = NewCT_ChartsheetProtection()
				if err := d.DecodeElement(m.SheetProtection, &el); err != nil {
					return err
				}
			case "customSheetViews":
				m.CustomSheetViews = NewCT_CustomChartsheetViews()
				if err := d.DecodeElement(m.CustomSheetViews, &el); err != nil {
					return err
				}
			case "pageMargins":
				m.PageMargins = NewCT_PageMargins()
				if err := d.DecodeElement(m.PageMargins, &el); err != nil {
					return err
				}
			case "pageSetup":
				m.PageSetup = NewCT_CsPageSetup()
				if err := d.DecodeElement(m.PageSetup, &el); err != nil {
					return err
				}
			case "headerFooter":
				m.HeaderFooter = NewCT_HeaderFooter()
				if err := d.DecodeElement(m.HeaderFooter, &el); err != nil {
					return err
				}
			case "drawing":
				if err := d.DecodeElement(m.Drawing, &el); err != nil {
					return err
				}
			case "legacyDrawing":
				m.LegacyDrawing = NewCT_LegacyDrawing()
				if err := d.DecodeElement(m.LegacyDrawing, &el); err != nil {
					return err
				}
			case "legacyDrawingHF":
				m.LegacyDrawingHF = NewCT_LegacyDrawing()
				if err := d.DecodeElement(m.LegacyDrawingHF, &el); err != nil {
					return err
				}
			case "drawingHF":
				m.DrawingHF = NewCT_DrawingHF()
				if err := d.DecodeElement(m.DrawingHF, &el); err != nil {
					return err
				}
			case "picture":
				m.Picture = NewCT_SheetBackgroundPicture()
				if err := d.DecodeElement(m.Picture, &el); err != nil {
					return err
				}
			case "webPublishItems":
				m.WebPublishItems = NewCT_WebPublishItems()
				if err := d.DecodeElement(m.WebPublishItems, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on Chartsheet %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lChartsheet
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Chartsheet and its children
func (m *Chartsheet) Validate() error {
	return m.ValidateWithPath("Chartsheet")
}

// ValidateWithPath validates the Chartsheet and its children, prefixing error messages with path
func (m *Chartsheet) ValidateWithPath(path string) error {
	if err := m.CT_Chartsheet.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
