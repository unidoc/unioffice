// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_Workbook struct {
	// Document Conformance Class
	ConformanceAttr sharedTypes.ST_ConformanceClass
	// File Version
	FileVersion *CT_FileVersion
	// File Sharing
	FileSharing *CT_FileSharing
	// Workbook Properties
	WorkbookPr *CT_WorkbookPr
	// Workbook Protection
	WorkbookProtection *CT_WorkbookProtection
	// Workbook Views
	BookViews *CT_BookViews
	// Sheets
	Sheets *CT_Sheets
	// Function Groups
	FunctionGroups *CT_FunctionGroups
	// External References
	ExternalReferences *CT_ExternalReferences
	// Defined Names
	DefinedNames *CT_DefinedNames
	// Calculation Properties
	CalcPr *CT_CalcPr
	// Embedded Object Size
	OleSize *CT_OleSize
	// Custom Workbook Views
	CustomWorkbookViews *CT_CustomWorkbookViews
	// PivotCaches
	PivotCaches *CT_PivotCaches
	// Smart Tag Properties
	SmartTagPr *CT_SmartTagPr
	// Smart Tag Types
	SmartTagTypes *CT_SmartTagTypes
	// Web Publishing Properties
	WebPublishing *CT_WebPublishing
	// File Recovery Properties
	FileRecoveryPr []*CT_FileRecoveryPr
	// Web Publish Objects
	WebPublishObjects *CT_WebPublishObjects
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_Workbook() *CT_Workbook {
	ret := &CT_Workbook{}
	ret.Sheets = NewCT_Sheets()
	return ret
}

func (m *CT_Workbook) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.ConformanceAttr != sharedTypes.ST_ConformanceClassUnset {
		attr, err := m.ConformanceAttr.MarshalXMLAttr(xml.Name{Local: "conformance"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.FileVersion != nil {
		sefileVersion := xml.StartElement{Name: xml.Name{Local: "x:fileVersion"}}
		e.EncodeElement(m.FileVersion, sefileVersion)
	}
	if m.FileSharing != nil {
		sefileSharing := xml.StartElement{Name: xml.Name{Local: "x:fileSharing"}}
		e.EncodeElement(m.FileSharing, sefileSharing)
	}
	if m.WorkbookPr != nil {
		seworkbookPr := xml.StartElement{Name: xml.Name{Local: "x:workbookPr"}}
		e.EncodeElement(m.WorkbookPr, seworkbookPr)
	}
	if m.WorkbookProtection != nil {
		seworkbookProtection := xml.StartElement{Name: xml.Name{Local: "x:workbookProtection"}}
		e.EncodeElement(m.WorkbookProtection, seworkbookProtection)
	}
	if m.BookViews != nil {
		sebookViews := xml.StartElement{Name: xml.Name{Local: "x:bookViews"}}
		e.EncodeElement(m.BookViews, sebookViews)
	}
	sesheets := xml.StartElement{Name: xml.Name{Local: "x:sheets"}}
	e.EncodeElement(m.Sheets, sesheets)
	if m.FunctionGroups != nil {
		sefunctionGroups := xml.StartElement{Name: xml.Name{Local: "x:functionGroups"}}
		e.EncodeElement(m.FunctionGroups, sefunctionGroups)
	}
	if m.ExternalReferences != nil {
		seexternalReferences := xml.StartElement{Name: xml.Name{Local: "x:externalReferences"}}
		e.EncodeElement(m.ExternalReferences, seexternalReferences)
	}
	if m.DefinedNames != nil {
		sedefinedNames := xml.StartElement{Name: xml.Name{Local: "x:definedNames"}}
		e.EncodeElement(m.DefinedNames, sedefinedNames)
	}
	if m.CalcPr != nil {
		secalcPr := xml.StartElement{Name: xml.Name{Local: "x:calcPr"}}
		e.EncodeElement(m.CalcPr, secalcPr)
	}
	if m.OleSize != nil {
		seoleSize := xml.StartElement{Name: xml.Name{Local: "x:oleSize"}}
		e.EncodeElement(m.OleSize, seoleSize)
	}
	if m.CustomWorkbookViews != nil {
		secustomWorkbookViews := xml.StartElement{Name: xml.Name{Local: "x:customWorkbookViews"}}
		e.EncodeElement(m.CustomWorkbookViews, secustomWorkbookViews)
	}
	if m.PivotCaches != nil {
		sepivotCaches := xml.StartElement{Name: xml.Name{Local: "x:pivotCaches"}}
		e.EncodeElement(m.PivotCaches, sepivotCaches)
	}
	if m.SmartTagPr != nil {
		sesmartTagPr := xml.StartElement{Name: xml.Name{Local: "x:smartTagPr"}}
		e.EncodeElement(m.SmartTagPr, sesmartTagPr)
	}
	if m.SmartTagTypes != nil {
		sesmartTagTypes := xml.StartElement{Name: xml.Name{Local: "x:smartTagTypes"}}
		e.EncodeElement(m.SmartTagTypes, sesmartTagTypes)
	}
	if m.WebPublishing != nil {
		sewebPublishing := xml.StartElement{Name: xml.Name{Local: "x:webPublishing"}}
		e.EncodeElement(m.WebPublishing, sewebPublishing)
	}
	if m.FileRecoveryPr != nil {
		sefileRecoveryPr := xml.StartElement{Name: xml.Name{Local: "x:fileRecoveryPr"}}
		e.EncodeElement(m.FileRecoveryPr, sefileRecoveryPr)
	}
	if m.WebPublishObjects != nil {
		sewebPublishObjects := xml.StartElement{Name: xml.Name{Local: "x:webPublishObjects"}}
		e.EncodeElement(m.WebPublishObjects, sewebPublishObjects)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Workbook) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Sheets = NewCT_Sheets()
	for _, attr := range start.Attr {
		if attr.Name.Local == "conformance" {
			m.ConformanceAttr.UnmarshalXMLAttr(attr)
		}
	}
lCT_Workbook:
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
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Workbook
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Workbook and its children
func (m *CT_Workbook) Validate() error {
	return m.ValidateWithPath("CT_Workbook")
}

// ValidateWithPath validates the CT_Workbook and its children, prefixing error messages with path
func (m *CT_Workbook) ValidateWithPath(path string) error {
	if err := m.ConformanceAttr.ValidateWithPath(path + "/ConformanceAttr"); err != nil {
		return err
	}
	if m.FileVersion != nil {
		if err := m.FileVersion.ValidateWithPath(path + "/FileVersion"); err != nil {
			return err
		}
	}
	if m.FileSharing != nil {
		if err := m.FileSharing.ValidateWithPath(path + "/FileSharing"); err != nil {
			return err
		}
	}
	if m.WorkbookPr != nil {
		if err := m.WorkbookPr.ValidateWithPath(path + "/WorkbookPr"); err != nil {
			return err
		}
	}
	if m.WorkbookProtection != nil {
		if err := m.WorkbookProtection.ValidateWithPath(path + "/WorkbookProtection"); err != nil {
			return err
		}
	}
	if m.BookViews != nil {
		if err := m.BookViews.ValidateWithPath(path + "/BookViews"); err != nil {
			return err
		}
	}
	if err := m.Sheets.ValidateWithPath(path + "/Sheets"); err != nil {
		return err
	}
	if m.FunctionGroups != nil {
		if err := m.FunctionGroups.ValidateWithPath(path + "/FunctionGroups"); err != nil {
			return err
		}
	}
	if m.ExternalReferences != nil {
		if err := m.ExternalReferences.ValidateWithPath(path + "/ExternalReferences"); err != nil {
			return err
		}
	}
	if m.DefinedNames != nil {
		if err := m.DefinedNames.ValidateWithPath(path + "/DefinedNames"); err != nil {
			return err
		}
	}
	if m.CalcPr != nil {
		if err := m.CalcPr.ValidateWithPath(path + "/CalcPr"); err != nil {
			return err
		}
	}
	if m.OleSize != nil {
		if err := m.OleSize.ValidateWithPath(path + "/OleSize"); err != nil {
			return err
		}
	}
	if m.CustomWorkbookViews != nil {
		if err := m.CustomWorkbookViews.ValidateWithPath(path + "/CustomWorkbookViews"); err != nil {
			return err
		}
	}
	if m.PivotCaches != nil {
		if err := m.PivotCaches.ValidateWithPath(path + "/PivotCaches"); err != nil {
			return err
		}
	}
	if m.SmartTagPr != nil {
		if err := m.SmartTagPr.ValidateWithPath(path + "/SmartTagPr"); err != nil {
			return err
		}
	}
	if m.SmartTagTypes != nil {
		if err := m.SmartTagTypes.ValidateWithPath(path + "/SmartTagTypes"); err != nil {
			return err
		}
	}
	if m.WebPublishing != nil {
		if err := m.WebPublishing.ValidateWithPath(path + "/WebPublishing"); err != nil {
			return err
		}
	}
	for i, v := range m.FileRecoveryPr {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/FileRecoveryPr[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.WebPublishObjects != nil {
		if err := m.WebPublishObjects.ValidateWithPath(path + "/WebPublishObjects"); err != nil {
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
