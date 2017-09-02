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

type Worksheet struct {
	CT_Worksheet
}

func NewWorksheet() *Worksheet {
	ret := &Worksheet{}
	ret.CT_Worksheet = *NewCT_Worksheet()
	return ret
}

func (m *Worksheet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:x"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "x:worksheet"
	return m.CT_Worksheet.MarshalXML(e, start)
}

func (m *Worksheet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Worksheet = *NewCT_Worksheet()
lWorksheet:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "sheetPr":
				m.SheetPr = NewCT_SheetPr()
				if err := d.DecodeElement(m.SheetPr, &el); err != nil {
					return err
				}
			case "dimension":
				m.Dimension = NewCT_SheetDimension()
				if err := d.DecodeElement(m.Dimension, &el); err != nil {
					return err
				}
			case "sheetViews":
				m.SheetViews = NewCT_SheetViews()
				if err := d.DecodeElement(m.SheetViews, &el); err != nil {
					return err
				}
			case "sheetFormatPr":
				m.SheetFormatPr = NewCT_SheetFormatPr()
				if err := d.DecodeElement(m.SheetFormatPr, &el); err != nil {
					return err
				}
			case "cols":
				tmp := NewCT_Cols()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Cols = append(m.Cols, tmp)
			case "sheetData":
				if err := d.DecodeElement(m.SheetData, &el); err != nil {
					return err
				}
			case "sheetCalcPr":
				m.SheetCalcPr = NewCT_SheetCalcPr()
				if err := d.DecodeElement(m.SheetCalcPr, &el); err != nil {
					return err
				}
			case "sheetProtection":
				m.SheetProtection = NewCT_SheetProtection()
				if err := d.DecodeElement(m.SheetProtection, &el); err != nil {
					return err
				}
			case "protectedRanges":
				m.ProtectedRanges = NewCT_ProtectedRanges()
				if err := d.DecodeElement(m.ProtectedRanges, &el); err != nil {
					return err
				}
			case "scenarios":
				m.Scenarios = NewCT_Scenarios()
				if err := d.DecodeElement(m.Scenarios, &el); err != nil {
					return err
				}
			case "autoFilter":
				m.AutoFilter = NewCT_AutoFilter()
				if err := d.DecodeElement(m.AutoFilter, &el); err != nil {
					return err
				}
			case "sortState":
				m.SortState = NewCT_SortState()
				if err := d.DecodeElement(m.SortState, &el); err != nil {
					return err
				}
			case "dataConsolidate":
				m.DataConsolidate = NewCT_DataConsolidate()
				if err := d.DecodeElement(m.DataConsolidate, &el); err != nil {
					return err
				}
			case "customSheetViews":
				m.CustomSheetViews = NewCT_CustomSheetViews()
				if err := d.DecodeElement(m.CustomSheetViews, &el); err != nil {
					return err
				}
			case "mergeCells":
				m.MergeCells = NewCT_MergeCells()
				if err := d.DecodeElement(m.MergeCells, &el); err != nil {
					return err
				}
			case "phoneticPr":
				m.PhoneticPr = NewCT_PhoneticPr()
				if err := d.DecodeElement(m.PhoneticPr, &el); err != nil {
					return err
				}
			case "conditionalFormatting":
				tmp := NewCT_ConditionalFormatting()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ConditionalFormatting = append(m.ConditionalFormatting, tmp)
			case "dataValidations":
				m.DataValidations = NewCT_DataValidations()
				if err := d.DecodeElement(m.DataValidations, &el); err != nil {
					return err
				}
			case "hyperlinks":
				m.Hyperlinks = NewCT_Hyperlinks()
				if err := d.DecodeElement(m.Hyperlinks, &el); err != nil {
					return err
				}
			case "printOptions":
				m.PrintOptions = NewCT_PrintOptions()
				if err := d.DecodeElement(m.PrintOptions, &el); err != nil {
					return err
				}
			case "pageMargins":
				m.PageMargins = NewCT_PageMargins()
				if err := d.DecodeElement(m.PageMargins, &el); err != nil {
					return err
				}
			case "pageSetup":
				m.PageSetup = NewCT_PageSetup()
				if err := d.DecodeElement(m.PageSetup, &el); err != nil {
					return err
				}
			case "headerFooter":
				m.HeaderFooter = NewCT_HeaderFooter()
				if err := d.DecodeElement(m.HeaderFooter, &el); err != nil {
					return err
				}
			case "rowBreaks":
				m.RowBreaks = NewCT_PageBreak()
				if err := d.DecodeElement(m.RowBreaks, &el); err != nil {
					return err
				}
			case "colBreaks":
				m.ColBreaks = NewCT_PageBreak()
				if err := d.DecodeElement(m.ColBreaks, &el); err != nil {
					return err
				}
			case "customProperties":
				m.CustomProperties = NewCT_CustomProperties()
				if err := d.DecodeElement(m.CustomProperties, &el); err != nil {
					return err
				}
			case "cellWatches":
				m.CellWatches = NewCT_CellWatches()
				if err := d.DecodeElement(m.CellWatches, &el); err != nil {
					return err
				}
			case "ignoredErrors":
				m.IgnoredErrors = NewCT_IgnoredErrors()
				if err := d.DecodeElement(m.IgnoredErrors, &el); err != nil {
					return err
				}
			case "smartTags":
				m.SmartTags = NewCT_SmartTags()
				if err := d.DecodeElement(m.SmartTags, &el); err != nil {
					return err
				}
			case "drawing":
				m.Drawing = NewCT_Drawing()
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
			case "oleObjects":
				m.OleObjects = NewCT_OleObjects()
				if err := d.DecodeElement(m.OleObjects, &el); err != nil {
					return err
				}
			case "controls":
				m.Controls = NewCT_Controls()
				if err := d.DecodeElement(m.Controls, &el); err != nil {
					return err
				}
			case "webPublishItems":
				m.WebPublishItems = NewCT_WebPublishItems()
				if err := d.DecodeElement(m.WebPublishItems, &el); err != nil {
					return err
				}
			case "tableParts":
				m.TableParts = NewCT_TableParts()
				if err := d.DecodeElement(m.TableParts, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on Worksheet %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWorksheet
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Worksheet and its children
func (m *Worksheet) Validate() error {
	return m.ValidateWithPath("Worksheet")
}

// ValidateWithPath validates the Worksheet and its children, prefixing error messages with path
func (m *Worksheet) ValidateWithPath(path string) error {
	if err := m.CT_Worksheet.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
