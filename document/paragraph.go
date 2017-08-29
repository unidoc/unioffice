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
	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
	wml "baliance.com/gooxml/schema/schemas.openxmlformats.org/wordprocessingml"
)

// Paragraph is a paragraph within a document.
type Paragraph struct {
	d *Document
	x *wml.CT_P
}

// X returns the inner wrapped XML type.
func (p Paragraph) X() *wml.CT_P {
	return p.x
}

func (p Paragraph) ensurePPr() {
	if p.x.PPr == nil {
		p.x.PPr = wml.NewCT_PPr()
	}
}

// SetSpacing sets the spacing that comes before and after the paragraph.
func (p Paragraph) SetSpacing(before, after measurement.Distance) {
	p.ensurePPr()
	p.x.PPr.Spacing = wml.NewCT_Spacing()
	p.x.PPr.Spacing.BeforeAttr = &sharedTypes.ST_TwipsMeasure{}
	p.x.PPr.Spacing.BeforeAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(before / measurement.Twips))
	p.x.PPr.Spacing.AfterAttr = &sharedTypes.ST_TwipsMeasure{}
	p.x.PPr.Spacing.AfterAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(after / measurement.Twips))
}

// SetStyle sets the style of a paragraph.
func (p Paragraph) SetStyle(s string) {
	p.ensurePPr()
	if s == "" {
		p.x.PPr.PStyle = nil
	} else {
		p.x.PPr.PStyle = wml.NewCT_String()
		p.x.PPr.PStyle.ValAttr = s
	}
}

// AddTabStop adds a tab stop to the paragraph.  It controls the position of text when using Run.AddTab()
func (p Paragraph) AddTabStop(position measurement.Distance, justificaton wml.ST_TabJc, leader wml.ST_TabTlc) {
	p.ensurePPr()
	if p.x.PPr.Tabs == nil {
		p.x.PPr.Tabs = wml.NewCT_Tabs()
	}
	tab := wml.NewCT_TabStop()
	tab.LeaderAttr = leader
	tab.ValAttr = justificaton
	tab.PosAttr.Int32 = gooxml.Int32(int32(position / measurement.Twips))
	p.x.PPr.Tabs.Tab = append(p.x.PPr.Tabs.Tab, tab)
}

// AddRun adds a run to a paragraph.
func (p Paragraph) AddRun() Run {
	var pc *wml.EG_PContent
	// no need to add a new EG_PContent if we already have one
	if len(p.x.EG_PContent) > 0 {
		pc = p.x.EG_PContent[len(p.x.EG_PContent)-1]
	} else {
		pc = wml.NewEG_PContent()
		p.x.EG_PContent = append(p.x.EG_PContent, pc)
	}

	rc := wml.NewEG_ContentRunContent()
	pc.EG_ContentRunContent = append(pc.EG_ContentRunContent, rc)
	r := wml.NewCT_R()
	rc.R = r
	return Run{p.d, r}
}

// Runs returns all of the runs in a paragraph.
func (p Paragraph) Runs() []Run {
	ret := []Run{}
	for _, c := range p.x.EG_PContent {
		for _, rc := range c.EG_ContentRunContent {
			if rc.R != nil {
				ret = append(ret, Run{p.d, rc.R})
			}
		}
	}
	return ret
}

// AddSection adds a new document section with an optional section break.  If t
// is ST_SectionMarkUnset, then no break will be inserted.
func (p Paragraph) AddSection(t wml.ST_SectionMark) Section {
	p.ensurePPr()
	p.x.PPr.SectPr = wml.NewCT_SectPr()
	if t != wml.ST_SectionMarkUnset {
		p.x.PPr.SectPr.Type = wml.NewCT_SectType()
		p.x.PPr.SectPr.Type.ValAttr = t
	}
	return Section{p.d, p.x.PPr.SectPr}
}

// SetHeadingLevel sets a heading level and style based on the level to a
// paragraph.  The default styles for a new gooxml document support headings
// from level 1 to 8.
func (p Paragraph) SetHeadingLevel(idx int) {
	p.ensurePPr()
	p.SetStyle(fmt.Sprintf("Heading%d", idx))
	if p.x.PPr.NumPr == nil {
		p.x.PPr.NumPr = wml.NewCT_NumPr()
	}
	p.x.PPr.NumPr.Ilvl = wml.NewCT_DecimalNumber()
	p.x.PPr.NumPr.Ilvl.ValAttr = int32(idx)
	p.x.PPr.NumPr.NumId = wml.NewCT_DecimalNumber()
	p.x.PPr.NumPr.NumId.ValAttr = int32(1)
}
