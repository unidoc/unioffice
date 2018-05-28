// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"fmt"

	"baliance.com/gooxml"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/schema/soo/ofc/sharedTypes"
	"baliance.com/gooxml/schema/soo/wml"
)

// ParagraphProperties are the properties for a paragraph.
type ParagraphProperties struct {
	d *Document
	x *wml.CT_PPr
}

// X returns the inner wrapped XML type.
func (p ParagraphProperties) X() *wml.CT_PPr {
	return p.x
}

// SetSpacing sets the spacing that comes before and after the paragraph.
// Deprecated: See Spacing() instead which allows finer control.
func (p ParagraphProperties) SetSpacing(before, after measurement.Distance) {
	if p.x.Spacing == nil {
		p.x.Spacing = wml.NewCT_Spacing()
	}
	p.x.Spacing.BeforeAttr = &sharedTypes.ST_TwipsMeasure{}
	p.x.Spacing.BeforeAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(before / measurement.Twips))
	p.x.Spacing.AfterAttr = &sharedTypes.ST_TwipsMeasure{}
	p.x.Spacing.AfterAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(after / measurement.Twips))
}

// Spacing returns the paragraph spacing settings.
func (p ParagraphProperties) Spacing() ParagraphSpacing {
	if p.x.Spacing == nil {
		p.x.Spacing = wml.NewCT_Spacing()
	}
	return ParagraphSpacing{p.x.Spacing}
}

// SetAlignment controls the paragraph alignment
func (p ParagraphProperties) SetAlignment(align wml.ST_Jc) {
	if align == wml.ST_JcUnset {
		p.x.Jc = nil
	} else {
		p.x.Jc = wml.NewCT_Jc()
		p.x.Jc.ValAttr = align
	}
}

// Style returns the style for a paragraph, or an empty string if it is unset.
func (p ParagraphProperties) Style() string {
	if p.x.PStyle != nil {
		return p.x.PStyle.ValAttr
	}
	return ""
}

// SetStyle sets the style of a paragraph.
func (p ParagraphProperties) SetStyle(s string) {
	if s == "" {
		p.x.PStyle = nil
	} else {
		p.x.PStyle = wml.NewCT_String()
		p.x.PStyle.ValAttr = s
	}
}

// AddTabStop adds a tab stop to the paragraph.  It controls the position of text when using Run.AddTab()
func (p ParagraphProperties) AddTabStop(position measurement.Distance, justificaton wml.ST_TabJc, leader wml.ST_TabTlc) {
	if p.x.Tabs == nil {
		p.x.Tabs = wml.NewCT_Tabs()
	}
	tab := wml.NewCT_TabStop()
	tab.LeaderAttr = leader
	tab.ValAttr = justificaton
	tab.PosAttr.Int64 = gooxml.Int64(int64(position / measurement.Twips))
	p.x.Tabs.Tab = append(p.x.Tabs.Tab, tab)
}

// AddSection adds a new document section with an optional section break.  If t
// is ST_SectionMarkUnset, then no break will be inserted.
func (p ParagraphProperties) AddSection(t wml.ST_SectionMark) Section {
	p.x.SectPr = wml.NewCT_SectPr()
	if t != wml.ST_SectionMarkUnset {
		p.x.SectPr.Type = wml.NewCT_SectType()
		p.x.SectPr.Type.ValAttr = t
	}
	return Section{p.d, p.x.SectPr}
}

// SetHeadingLevel sets a heading level and style based on the level to a
// paragraph.  The default styles for a new gooxml document support headings
// from level 1 to 8.
func (p ParagraphProperties) SetHeadingLevel(idx int) {
	p.SetStyle(fmt.Sprintf("Heading%d", idx))
	if p.x.NumPr == nil {
		p.x.NumPr = wml.NewCT_NumPr()
	}
	p.x.NumPr.Ilvl = wml.NewCT_DecimalNumber()
	p.x.NumPr.Ilvl.ValAttr = int64(idx)
}

// SetKeepWithNext controls if this paragraph should be kept with the next.
func (p ParagraphProperties) SetKeepWithNext(b bool) {
	if !b {
		p.x.KeepNext = nil
	} else {
		p.x.KeepNext = wml.NewCT_OnOff()
	}
}

// SetKeepOnOnePage controls if all lines in a paragraph are kept on the same
// page.
func (p ParagraphProperties) SetKeepOnOnePage(b bool) {
	if !b {
		p.x.KeepLines = nil
	} else {
		p.x.KeepLines = wml.NewCT_OnOff()
	}
}

// SetPageBreakBefore controls if there is a page break before this paragraph.
func (p ParagraphProperties) SetPageBreakBefore(b bool) {
	if !b {
		p.x.PageBreakBefore = nil
	} else {
		p.x.PageBreakBefore = wml.NewCT_OnOff()
	}
}

// SetWindowControl controls if the first or last line of the paragraph is
// allowed to dispay on a separate page.
func (p ParagraphProperties) SetWindowControl(b bool) {
	if !b {
		p.x.WidowControl = nil
	} else {
		p.x.WidowControl = wml.NewCT_OnOff()
	}
}

// SetFirstLineIndent controls the indentation of the first line in a paragraph.
func (p ParagraphProperties) SetFirstLineIndent(m measurement.Distance) {
	if p.x.Ind == nil {
		p.x.Ind = wml.NewCT_Ind()
	}
	if m == measurement.Zero {
		p.x.Ind.FirstLineAttr = nil
	} else {
		p.x.Ind.FirstLineAttr = &sharedTypes.ST_TwipsMeasure{}
		p.x.Ind.FirstLineAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(m / measurement.Twips))
	}
}

// SetStartIndent controls the start indentation.
func (p ParagraphProperties) SetStartIndent(m measurement.Distance) {
	if p.x.Ind == nil {
		p.x.Ind = wml.NewCT_Ind()
	}
	if m == measurement.Zero {
		p.x.Ind.StartAttr = nil
	} else {
		p.x.Ind.StartAttr = &wml.ST_SignedTwipsMeasure{}
		p.x.Ind.StartAttr.Int64 = gooxml.Int64(int64(m / measurement.Twips))
	}
}

// SetEndIndent controls the end indentation.
func (p ParagraphProperties) SetEndIndent(m measurement.Distance) {
	if p.x.Ind == nil {
		p.x.Ind = wml.NewCT_Ind()
	}
	if m == measurement.Zero {
		p.x.Ind.EndAttr = nil
	} else {
		p.x.Ind.EndAttr = &wml.ST_SignedTwipsMeasure{}
		p.x.Ind.EndAttr.Int64 = gooxml.Int64(int64(m / measurement.Twips))
	}
}

// SetHangingIndent controls the indentation of the non-first lines in a paragraph.
func (p ParagraphProperties) SetHangingIndent(m measurement.Distance) {
	if p.x.Ind == nil {
		p.x.Ind = wml.NewCT_Ind()
	}
	if m == measurement.Zero {
		p.x.Ind.HangingAttr = nil
	} else {
		p.x.Ind.HangingAttr = &sharedTypes.ST_TwipsMeasure{}
		p.x.Ind.HangingAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(m / measurement.Twips))
	}
}
