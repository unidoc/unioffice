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

type Workbook struct {
	CT_Workbook
}

func NewWorkbook() *Workbook {
	ret := &Workbook{}
	ret.CT_Workbook = *NewCT_Workbook()
	return ret
}

func (m *Workbook) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:x"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "x:workbook"
	return m.CT_Workbook.MarshalXML(e, start)
}

func (m *Workbook) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Workbook = *NewCT_Workbook()
	for _, attr := range start.Attr {
		if attr.Name.Local == "conformance" {
			m.ConformanceAttr.UnmarshalXMLAttr(attr)
		}
	}
lWorkbook:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "fileVersion":
				m.FileVersion = NewCT_FileVersion()
				if err := d.DecodeElement(m.FileVersion, &el); err != nil {
					return err
				}
			case "fileSharing":
				m.FileSharing = NewCT_FileSharing()
				if err := d.DecodeElement(m.FileSharing, &el); err != nil {
					return err
				}
			case "workbookPr":
				m.WorkbookPr = NewCT_WorkbookPr()
				if err := d.DecodeElement(m.WorkbookPr, &el); err != nil {
					return err
				}
			case "workbookProtection":
				m.WorkbookProtection = NewCT_WorkbookProtection()
				if err := d.DecodeElement(m.WorkbookProtection, &el); err != nil {
					return err
				}
			case "bookViews":
				m.BookViews = NewCT_BookViews()
				if err := d.DecodeElement(m.BookViews, &el); err != nil {
					return err
				}
			case "sheets":
				if err := d.DecodeElement(m.Sheets, &el); err != nil {
					return err
				}
			case "functionGroups":
				m.FunctionGroups = NewCT_FunctionGroups()
				if err := d.DecodeElement(m.FunctionGroups, &el); err != nil {
					return err
				}
			case "externalReferences":
				m.ExternalReferences = NewCT_ExternalReferences()
				if err := d.DecodeElement(m.ExternalReferences, &el); err != nil {
					return err
				}
			case "definedNames":
				m.DefinedNames = NewCT_DefinedNames()
				if err := d.DecodeElement(m.DefinedNames, &el); err != nil {
					return err
				}
			case "calcPr":
				m.CalcPr = NewCT_CalcPr()
				if err := d.DecodeElement(m.CalcPr, &el); err != nil {
					return err
				}
			case "oleSize":
				m.OleSize = NewCT_OleSize()
				if err := d.DecodeElement(m.OleSize, &el); err != nil {
					return err
				}
			case "customWorkbookViews":
				m.CustomWorkbookViews = NewCT_CustomWorkbookViews()
				if err := d.DecodeElement(m.CustomWorkbookViews, &el); err != nil {
					return err
				}
			case "pivotCaches":
				m.PivotCaches = NewCT_PivotCaches()
				if err := d.DecodeElement(m.PivotCaches, &el); err != nil {
					return err
				}
			case "smartTagPr":
				m.SmartTagPr = NewCT_SmartTagPr()
				if err := d.DecodeElement(m.SmartTagPr, &el); err != nil {
					return err
				}
			case "smartTagTypes":
				m.SmartTagTypes = NewCT_SmartTagTypes()
				if err := d.DecodeElement(m.SmartTagTypes, &el); err != nil {
					return err
				}
			case "webPublishing":
				m.WebPublishing = NewCT_WebPublishing()
				if err := d.DecodeElement(m.WebPublishing, &el); err != nil {
					return err
				}
			case "fileRecoveryPr":
				tmp := NewCT_FileRecoveryPr()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.FileRecoveryPr = append(m.FileRecoveryPr, tmp)
			case "webPublishObjects":
				m.WebPublishObjects = NewCT_WebPublishObjects()
				if err := d.DecodeElement(m.WebPublishObjects, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on Workbook %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWorkbook
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Workbook and its children
func (m *Workbook) Validate() error {
	return m.ValidateWithPath("Workbook")
}

// ValidateWithPath validates the Workbook and its children, prefixing error messages with path
func (m *Workbook) ValidateWithPath(path string) error {
	if err := m.CT_Workbook.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
