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
)

type CT_BuildList struct {
	// Build Paragraph
	BldP []*CT_TLBuildParagraph
	// Build Diagram
	BldDgm []*CT_TLBuildDiagram
	// Build Embedded Chart
	BldOleChart []*CT_TLOleBuildChart
	// Build Graphics
	BldGraphic []*CT_TLGraphicalObjectBuild
}

func NewCT_BuildList() *CT_BuildList {
	ret := &CT_BuildList{}
	return ret
}
func (m *CT_BuildList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	if m.BldP != nil {
		sebldP := xml.StartElement{Name: xml.Name{Local: "p:bldP"}}
		e.EncodeElement(m.BldP, sebldP)
	}
	if m.BldDgm != nil {
		sebldDgm := xml.StartElement{Name: xml.Name{Local: "p:bldDgm"}}
		e.EncodeElement(m.BldDgm, sebldDgm)
	}
	if m.BldOleChart != nil {
		sebldOleChart := xml.StartElement{Name: xml.Name{Local: "p:bldOleChart"}}
		e.EncodeElement(m.BldOleChart, sebldOleChart)
	}
	if m.BldGraphic != nil {
		sebldGraphic := xml.StartElement{Name: xml.Name{Local: "p:bldGraphic"}}
		e.EncodeElement(m.BldGraphic, sebldGraphic)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_BuildList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_BuildList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "bldP":
				tmp := NewCT_TLBuildParagraph()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.BldP = append(m.BldP, tmp)
			case "bldDgm":
				tmp := NewCT_TLBuildDiagram()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.BldDgm = append(m.BldDgm, tmp)
			case "bldOleChart":
				tmp := NewCT_TLOleBuildChart()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.BldOleChart = append(m.BldOleChart, tmp)
			case "bldGraphic":
				tmp := NewCT_TLGraphicalObjectBuild()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.BldGraphic = append(m.BldGraphic, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BuildList
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_BuildList) Validate() error {
	return m.ValidateWithPath("CT_BuildList")
}
func (m *CT_BuildList) ValidateWithPath(path string) error {
	for i, v := range m.BldP {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/BldP[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.BldDgm {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/BldDgm[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.BldOleChart {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/BldOleChart[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.BldGraphic {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/BldGraphic[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
