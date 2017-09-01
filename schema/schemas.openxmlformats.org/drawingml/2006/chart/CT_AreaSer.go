// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"fmt"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

type CT_AreaSer struct {
	Idx            *CT_UnsignedInt
	Order          *CT_UnsignedInt
	Tx             *CT_SerTx
	SpPr           *drawingml.CT_ShapeProperties
	PictureOptions *CT_PictureOptions
	DPt            []*CT_DPt
	DLbls          *CT_DLbls
	Trendline      []*CT_Trendline
	ErrBars        []*CT_ErrBars
	Cat            *CT_AxDataSource
	Val            *CT_NumDataSource
	ExtLst         *CT_ExtensionList
}

func NewCT_AreaSer() *CT_AreaSer {
	ret := &CT_AreaSer{}
	ret.Idx = NewCT_UnsignedInt()
	ret.Order = NewCT_UnsignedInt()
	return ret
}

func (m *CT_AreaSer) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	seidx := xml.StartElement{Name: xml.Name{Local: "idx"}}
	e.EncodeElement(m.Idx, seidx)
	seorder := xml.StartElement{Name: xml.Name{Local: "order"}}
	e.EncodeElement(m.Order, seorder)
	if m.Tx != nil {
		setx := xml.StartElement{Name: xml.Name{Local: "tx"}}
		e.EncodeElement(m.Tx, setx)
	}
	if m.SpPr != nil {
		sespPr := xml.StartElement{Name: xml.Name{Local: "spPr"}}
		e.EncodeElement(m.SpPr, sespPr)
	}
	if m.PictureOptions != nil {
		sepictureOptions := xml.StartElement{Name: xml.Name{Local: "pictureOptions"}}
		e.EncodeElement(m.PictureOptions, sepictureOptions)
	}
	if m.DPt != nil {
		sedPt := xml.StartElement{Name: xml.Name{Local: "dPt"}}
		e.EncodeElement(m.DPt, sedPt)
	}
	if m.DLbls != nil {
		sedLbls := xml.StartElement{Name: xml.Name{Local: "dLbls"}}
		e.EncodeElement(m.DLbls, sedLbls)
	}
	if m.Trendline != nil {
		setrendline := xml.StartElement{Name: xml.Name{Local: "trendline"}}
		e.EncodeElement(m.Trendline, setrendline)
	}
	if m.ErrBars != nil {
		seerrBars := xml.StartElement{Name: xml.Name{Local: "errBars"}}
		e.EncodeElement(m.ErrBars, seerrBars)
	}
	if m.Cat != nil {
		secat := xml.StartElement{Name: xml.Name{Local: "cat"}}
		e.EncodeElement(m.Cat, secat)
	}
	if m.Val != nil {
		seval := xml.StartElement{Name: xml.Name{Local: "val"}}
		e.EncodeElement(m.Val, seval)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AreaSer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Idx = NewCT_UnsignedInt()
	m.Order = NewCT_UnsignedInt()
lCT_AreaSer:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "idx":
				if err := d.DecodeElement(m.Idx, &el); err != nil {
					return err
				}
			case "order":
				if err := d.DecodeElement(m.Order, &el); err != nil {
					return err
				}
			case "tx":
				m.Tx = NewCT_SerTx()
				if err := d.DecodeElement(m.Tx, &el); err != nil {
					return err
				}
			case "spPr":
				m.SpPr = drawingml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case "pictureOptions":
				m.PictureOptions = NewCT_PictureOptions()
				if err := d.DecodeElement(m.PictureOptions, &el); err != nil {
					return err
				}
			case "dPt":
				tmp := NewCT_DPt()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.DPt = append(m.DPt, tmp)
			case "dLbls":
				m.DLbls = NewCT_DLbls()
				if err := d.DecodeElement(m.DLbls, &el); err != nil {
					return err
				}
			case "trendline":
				tmp := NewCT_Trendline()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Trendline = append(m.Trendline, tmp)
			case "errBars":
				tmp := NewCT_ErrBars()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ErrBars = append(m.ErrBars, tmp)
			case "cat":
				m.Cat = NewCT_AxDataSource()
				if err := d.DecodeElement(m.Cat, &el); err != nil {
					return err
				}
			case "val":
				m.Val = NewCT_NumDataSource()
				if err := d.DecodeElement(m.Val, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_AreaSer %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AreaSer
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_AreaSer and its children
func (m *CT_AreaSer) Validate() error {
	return m.ValidateWithPath("CT_AreaSer")
}

// ValidateWithPath validates the CT_AreaSer and its children, prefixing error messages with path
func (m *CT_AreaSer) ValidateWithPath(path string) error {
	if err := m.Idx.ValidateWithPath(path + "/Idx"); err != nil {
		return err
	}
	if err := m.Order.ValidateWithPath(path + "/Order"); err != nil {
		return err
	}
	if m.Tx != nil {
		if err := m.Tx.ValidateWithPath(path + "/Tx"); err != nil {
			return err
		}
	}
	if m.SpPr != nil {
		if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
			return err
		}
	}
	if m.PictureOptions != nil {
		if err := m.PictureOptions.ValidateWithPath(path + "/PictureOptions"); err != nil {
			return err
		}
	}
	for i, v := range m.DPt {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/DPt[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.DLbls != nil {
		if err := m.DLbls.ValidateWithPath(path + "/DLbls"); err != nil {
			return err
		}
	}
	for i, v := range m.Trendline {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Trendline[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.ErrBars {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ErrBars[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.Cat != nil {
		if err := m.Cat.ValidateWithPath(path + "/Cat"); err != nil {
			return err
		}
	}
	if m.Val != nil {
		if err := m.Val.ValidateWithPath(path + "/Val"); err != nil {
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
