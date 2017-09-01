// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_R struct {
	// Revision Identifier for Run Properties
	RsidRPrAttr *string
	// Revision Identifier for Run Deletion
	RsidDelAttr *string
	// Revision Identifier for Run
	RsidRAttr *string
	// Run Properties
	RPr                *CT_RPr
	EG_RunInnerContent []*EG_RunInnerContent
}

func NewCT_R() *CT_R {
	ret := &CT_R{}
	return ret
}
func (m *CT_R) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.RsidRPrAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidRPr"},
			Value: fmt.Sprintf("%v", *m.RsidRPrAttr)})
	}
	if m.RsidDelAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidDel"},
			Value: fmt.Sprintf("%v", *m.RsidDelAttr)})
	}
	if m.RsidRAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:rsidR"},
			Value: fmt.Sprintf("%v", *m.RsidRAttr)})
	}
	e.EncodeToken(start)
	if m.RPr != nil {
		serPr := xml.StartElement{Name: xml.Name{Local: "w:rPr"}}
		e.EncodeElement(m.RPr, serPr)
	}
	if m.EG_RunInnerContent != nil {
		for _, c := range m.EG_RunInnerContent {
			c.MarshalXML(e, start)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_R) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "rsidRPr" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RsidRPrAttr = &parsed
		}
		if attr.Name.Local == "rsidDel" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RsidDelAttr = &parsed
		}
		if attr.Name.Local == "rsidR" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RsidRAttr = &parsed
		}
	}
lCT_R:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "rPr":
				m.RPr = NewCT_RPr()
				if err := d.DecodeElement(m.RPr, &el); err != nil {
					return err
				}
			case "br":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.Br = NewCT_Br()
				if err := d.DecodeElement(tmpruninnercontent.Br, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "t":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.T = NewCT_Text()
				if err := d.DecodeElement(tmpruninnercontent.T, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "contentPart":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.ContentPart = NewCT_Rel()
				if err := d.DecodeElement(tmpruninnercontent.ContentPart, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "delText":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.DelText = NewCT_Text()
				if err := d.DecodeElement(tmpruninnercontent.DelText, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "instrText":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.InstrText = NewCT_Text()
				if err := d.DecodeElement(tmpruninnercontent.InstrText, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "delInstrText":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.DelInstrText = NewCT_Text()
				if err := d.DecodeElement(tmpruninnercontent.DelInstrText, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "noBreakHyphen":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.NoBreakHyphen = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.NoBreakHyphen, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "softHyphen":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.SoftHyphen = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.SoftHyphen, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "dayShort":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.DayShort = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.DayShort, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "monthShort":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.MonthShort = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.MonthShort, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "yearShort":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.YearShort = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.YearShort, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "dayLong":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.DayLong = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.DayLong, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "monthLong":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.MonthLong = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.MonthLong, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "yearLong":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.YearLong = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.YearLong, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "annotationRef":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.AnnotationRef = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.AnnotationRef, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "footnoteRef":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.FootnoteRef = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.FootnoteRef, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "endnoteRef":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.EndnoteRef = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.EndnoteRef, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "separator":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.Separator = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.Separator, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "continuationSeparator":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.ContinuationSeparator = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.ContinuationSeparator, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "sym":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.Sym = NewCT_Sym()
				if err := d.DecodeElement(tmpruninnercontent.Sym, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "pgNum":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.PgNum = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.PgNum, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "cr":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.Cr = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.Cr, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "tab":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.Tab = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.Tab, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "object":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.Object = NewCT_Object()
				if err := d.DecodeElement(tmpruninnercontent.Object, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "pict":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.Pict = NewCT_Picture()
				if err := d.DecodeElement(tmpruninnercontent.Pict, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "fldChar":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.FldChar = NewCT_FldChar()
				if err := d.DecodeElement(tmpruninnercontent.FldChar, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "ruby":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.Ruby = NewCT_Ruby()
				if err := d.DecodeElement(tmpruninnercontent.Ruby, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "footnoteReference":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.FootnoteReference = NewCT_FtnEdnRef()
				if err := d.DecodeElement(tmpruninnercontent.FootnoteReference, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "endnoteReference":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.EndnoteReference = NewCT_FtnEdnRef()
				if err := d.DecodeElement(tmpruninnercontent.EndnoteReference, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "commentReference":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.CommentReference = NewCT_Markup()
				if err := d.DecodeElement(tmpruninnercontent.CommentReference, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "drawing":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.Drawing = NewCT_Drawing()
				if err := d.DecodeElement(tmpruninnercontent.Drawing, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "ptab":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.Ptab = NewCT_PTab()
				if err := d.DecodeElement(tmpruninnercontent.Ptab, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			case "lastRenderedPageBreak":
				tmpruninnercontent := NewEG_RunInnerContent()
				tmpruninnercontent.LastRenderedPageBreak = NewCT_Empty()
				if err := d.DecodeElement(tmpruninnercontent.LastRenderedPageBreak, &el); err != nil {
					return err
				}
				m.EG_RunInnerContent = append(m.EG_RunInnerContent, tmpruninnercontent)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_R
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_R) Validate() error {
	return m.ValidateWithPath("CT_R")
}
func (m *CT_R) ValidateWithPath(path string) error {
	if m.RPr != nil {
		if err := m.RPr.ValidateWithPath(path + "/RPr"); err != nil {
			return err
		}
	}
	for i, v := range m.EG_RunInnerContent {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_RunInnerContent[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
