// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"log"
)

type CT_Chart struct {
	Title            *CT_Title
	AutoTitleDeleted *CT_Boolean
	PivotFmts        *CT_PivotFmts
	View3D           *CT_View3D
	Floor            *CT_Surface
	SideWall         *CT_Surface
	BackWall         *CT_Surface
	PlotArea         *CT_PlotArea
	Legend           *CT_Legend
	PlotVisOnly      *CT_Boolean
	DispBlanksAs     *CT_DispBlanksAs
	ShowDLblsOverMax *CT_Boolean
	ExtLst           *CT_ExtensionList
}

func NewCT_Chart() *CT_Chart {
	ret := &CT_Chart{}
	ret.PlotArea = NewCT_PlotArea()
	return ret
}

func (m *CT_Chart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Title != nil {
		setitle := xml.StartElement{Name: xml.Name{Local: "title"}}
		e.EncodeElement(m.Title, setitle)
	}
	if m.AutoTitleDeleted != nil {
		seautoTitleDeleted := xml.StartElement{Name: xml.Name{Local: "autoTitleDeleted"}}
		e.EncodeElement(m.AutoTitleDeleted, seautoTitleDeleted)
	}
	if m.PivotFmts != nil {
		sepivotFmts := xml.StartElement{Name: xml.Name{Local: "pivotFmts"}}
		e.EncodeElement(m.PivotFmts, sepivotFmts)
	}
	if m.View3D != nil {
		seview3D := xml.StartElement{Name: xml.Name{Local: "view3D"}}
		e.EncodeElement(m.View3D, seview3D)
	}
	if m.Floor != nil {
		sefloor := xml.StartElement{Name: xml.Name{Local: "floor"}}
		e.EncodeElement(m.Floor, sefloor)
	}
	if m.SideWall != nil {
		sesideWall := xml.StartElement{Name: xml.Name{Local: "sideWall"}}
		e.EncodeElement(m.SideWall, sesideWall)
	}
	if m.BackWall != nil {
		sebackWall := xml.StartElement{Name: xml.Name{Local: "backWall"}}
		e.EncodeElement(m.BackWall, sebackWall)
	}
	seplotArea := xml.StartElement{Name: xml.Name{Local: "plotArea"}}
	e.EncodeElement(m.PlotArea, seplotArea)
	if m.Legend != nil {
		selegend := xml.StartElement{Name: xml.Name{Local: "legend"}}
		e.EncodeElement(m.Legend, selegend)
	}
	if m.PlotVisOnly != nil {
		seplotVisOnly := xml.StartElement{Name: xml.Name{Local: "plotVisOnly"}}
		e.EncodeElement(m.PlotVisOnly, seplotVisOnly)
	}
	if m.DispBlanksAs != nil {
		sedispBlanksAs := xml.StartElement{Name: xml.Name{Local: "dispBlanksAs"}}
		e.EncodeElement(m.DispBlanksAs, sedispBlanksAs)
	}
	if m.ShowDLblsOverMax != nil {
		seshowDLblsOverMax := xml.StartElement{Name: xml.Name{Local: "showDLblsOverMax"}}
		e.EncodeElement(m.ShowDLblsOverMax, seshowDLblsOverMax)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Chart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.PlotArea = NewCT_PlotArea()
lCT_Chart:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "title":
				m.Title = NewCT_Title()
				if err := d.DecodeElement(m.Title, &el); err != nil {
					return err
				}
			case "autoTitleDeleted":
				m.AutoTitleDeleted = NewCT_Boolean()
				if err := d.DecodeElement(m.AutoTitleDeleted, &el); err != nil {
					return err
				}
			case "pivotFmts":
				m.PivotFmts = NewCT_PivotFmts()
				if err := d.DecodeElement(m.PivotFmts, &el); err != nil {
					return err
				}
			case "view3D":
				m.View3D = NewCT_View3D()
				if err := d.DecodeElement(m.View3D, &el); err != nil {
					return err
				}
			case "floor":
				m.Floor = NewCT_Surface()
				if err := d.DecodeElement(m.Floor, &el); err != nil {
					return err
				}
			case "sideWall":
				m.SideWall = NewCT_Surface()
				if err := d.DecodeElement(m.SideWall, &el); err != nil {
					return err
				}
			case "backWall":
				m.BackWall = NewCT_Surface()
				if err := d.DecodeElement(m.BackWall, &el); err != nil {
					return err
				}
			case "plotArea":
				if err := d.DecodeElement(m.PlotArea, &el); err != nil {
					return err
				}
			case "legend":
				m.Legend = NewCT_Legend()
				if err := d.DecodeElement(m.Legend, &el); err != nil {
					return err
				}
			case "plotVisOnly":
				m.PlotVisOnly = NewCT_Boolean()
				if err := d.DecodeElement(m.PlotVisOnly, &el); err != nil {
					return err
				}
			case "dispBlanksAs":
				m.DispBlanksAs = NewCT_DispBlanksAs()
				if err := d.DecodeElement(m.DispBlanksAs, &el); err != nil {
					return err
				}
			case "showDLblsOverMax":
				m.ShowDLblsOverMax = NewCT_Boolean()
				if err := d.DecodeElement(m.ShowDLblsOverMax, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Chart %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Chart
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Chart and its children
func (m *CT_Chart) Validate() error {
	return m.ValidateWithPath("CT_Chart")
}

// ValidateWithPath validates the CT_Chart and its children, prefixing error messages with path
func (m *CT_Chart) ValidateWithPath(path string) error {
	if m.Title != nil {
		if err := m.Title.ValidateWithPath(path + "/Title"); err != nil {
			return err
		}
	}
	if m.AutoTitleDeleted != nil {
		if err := m.AutoTitleDeleted.ValidateWithPath(path + "/AutoTitleDeleted"); err != nil {
			return err
		}
	}
	if m.PivotFmts != nil {
		if err := m.PivotFmts.ValidateWithPath(path + "/PivotFmts"); err != nil {
			return err
		}
	}
	if m.View3D != nil {
		if err := m.View3D.ValidateWithPath(path + "/View3D"); err != nil {
			return err
		}
	}
	if m.Floor != nil {
		if err := m.Floor.ValidateWithPath(path + "/Floor"); err != nil {
			return err
		}
	}
	if m.SideWall != nil {
		if err := m.SideWall.ValidateWithPath(path + "/SideWall"); err != nil {
			return err
		}
	}
	if m.BackWall != nil {
		if err := m.BackWall.ValidateWithPath(path + "/BackWall"); err != nil {
			return err
		}
	}
	if err := m.PlotArea.ValidateWithPath(path + "/PlotArea"); err != nil {
		return err
	}
	if m.Legend != nil {
		if err := m.Legend.ValidateWithPath(path + "/Legend"); err != nil {
			return err
		}
	}
	if m.PlotVisOnly != nil {
		if err := m.PlotVisOnly.ValidateWithPath(path + "/PlotVisOnly"); err != nil {
			return err
		}
	}
	if m.DispBlanksAs != nil {
		if err := m.DispBlanksAs.ValidateWithPath(path + "/DispBlanksAs"); err != nil {
			return err
		}
	}
	if m.ShowDLblsOverMax != nil {
		if err := m.ShowDLblsOverMax.ValidateWithPath(path + "/ShowDLblsOverMax"); err != nil {
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
