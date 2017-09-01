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

type CT_Stylesheet struct {
	// Number Formats
	NumFmts *CT_NumFmts
	// Fonts
	Fonts *CT_Fonts
	// Fills
	Fills *CT_Fills
	// Borders
	Borders *CT_Borders
	// Formatting Records
	CellStyleXfs *CT_CellStyleXfs
	// Cell Formats
	CellXfs *CT_CellXfs
	// Cell Styles
	CellStyles *CT_CellStyles
	// Formats
	Dxfs *CT_Dxfs
	// Table Styles
	TableStyles *CT_TableStyles
	// Colors
	Colors *CT_Colors
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_Stylesheet() *CT_Stylesheet {
	ret := &CT_Stylesheet{}
	return ret
}
func (m *CT_Stylesheet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.NumFmts != nil {
		senumFmts := xml.StartElement{Name: xml.Name{Local: "x:numFmts"}}
		e.EncodeElement(m.NumFmts, senumFmts)
	}
	if m.Fonts != nil {
		sefonts := xml.StartElement{Name: xml.Name{Local: "x:fonts"}}
		e.EncodeElement(m.Fonts, sefonts)
	}
	if m.Fills != nil {
		sefills := xml.StartElement{Name: xml.Name{Local: "x:fills"}}
		e.EncodeElement(m.Fills, sefills)
	}
	if m.Borders != nil {
		seborders := xml.StartElement{Name: xml.Name{Local: "x:borders"}}
		e.EncodeElement(m.Borders, seborders)
	}
	if m.CellStyleXfs != nil {
		secellStyleXfs := xml.StartElement{Name: xml.Name{Local: "x:cellStyleXfs"}}
		e.EncodeElement(m.CellStyleXfs, secellStyleXfs)
	}
	if m.CellXfs != nil {
		secellXfs := xml.StartElement{Name: xml.Name{Local: "x:cellXfs"}}
		e.EncodeElement(m.CellXfs, secellXfs)
	}
	if m.CellStyles != nil {
		secellStyles := xml.StartElement{Name: xml.Name{Local: "x:cellStyles"}}
		e.EncodeElement(m.CellStyles, secellStyles)
	}
	if m.Dxfs != nil {
		sedxfs := xml.StartElement{Name: xml.Name{Local: "x:dxfs"}}
		e.EncodeElement(m.Dxfs, sedxfs)
	}
	if m.TableStyles != nil {
		setableStyles := xml.StartElement{Name: xml.Name{Local: "x:tableStyles"}}
		e.EncodeElement(m.TableStyles, setableStyles)
	}
	if m.Colors != nil {
		secolors := xml.StartElement{Name: xml.Name{Local: "x:colors"}}
		e.EncodeElement(m.Colors, secolors)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "x:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Stylesheet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Stylesheet:
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
			break lCT_Stylesheet
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Stylesheet) Validate() error {
	return m.ValidateWithPath("CT_Stylesheet")
}
func (m *CT_Stylesheet) ValidateWithPath(path string) error {
	if m.NumFmts != nil {
		if err := m.NumFmts.ValidateWithPath(path + "/NumFmts"); err != nil {
			return err
		}
	}
	if m.Fonts != nil {
		if err := m.Fonts.ValidateWithPath(path + "/Fonts"); err != nil {
			return err
		}
	}
	if m.Fills != nil {
		if err := m.Fills.ValidateWithPath(path + "/Fills"); err != nil {
			return err
		}
	}
	if m.Borders != nil {
		if err := m.Borders.ValidateWithPath(path + "/Borders"); err != nil {
			return err
		}
	}
	if m.CellStyleXfs != nil {
		if err := m.CellStyleXfs.ValidateWithPath(path + "/CellStyleXfs"); err != nil {
			return err
		}
	}
	if m.CellXfs != nil {
		if err := m.CellXfs.ValidateWithPath(path + "/CellXfs"); err != nil {
			return err
		}
	}
	if m.CellStyles != nil {
		if err := m.CellStyles.ValidateWithPath(path + "/CellStyles"); err != nil {
			return err
		}
	}
	if m.Dxfs != nil {
		if err := m.Dxfs.ValidateWithPath(path + "/Dxfs"); err != nil {
			return err
		}
	}
	if m.TableStyles != nil {
		if err := m.TableStyles.ValidateWithPath(path + "/TableStyles"); err != nil {
			return err
		}
	}
	if m.Colors != nil {
		if err := m.Colors.ValidateWithPath(path + "/Colors"); err != nil {
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
