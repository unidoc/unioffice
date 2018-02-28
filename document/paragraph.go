// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"baliance.com/gooxml/schema/soo/wml"
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

// RemoveRun removes a child run from a paragraph.
func (p Paragraph) RemoveRun(r Run) {
	for _, c := range p.x.EG_PContent {
		for i, rc := range c.EG_ContentRunContent {
			if rc.R == r.x {
				copy(c.EG_ContentRunContent[i:], c.EG_ContentRunContent[i+1:])
				c.EG_ContentRunContent = c.EG_ContentRunContent[0 : len(c.EG_ContentRunContent)-1]
			}
			if rc.Sdt != nil && rc.Sdt.SdtContent != nil {
				for i, rc2 := range rc.Sdt.SdtContent.EG_ContentRunContent {
					if rc2.R == r.x {
						copy(rc.Sdt.SdtContent.EG_ContentRunContent[i:], rc.Sdt.SdtContent.EG_ContentRunContent[i+1:])
						rc.Sdt.SdtContent.EG_ContentRunContent = rc.Sdt.SdtContent.EG_ContentRunContent[0 : len(rc.Sdt.SdtContent.EG_ContentRunContent)-1]
					}
				}
			}
		}
	}
}

// Properties returns the paragraph properties.
func (p Paragraph) Properties() ParagraphProperties {
	p.ensurePPr()
	return ParagraphProperties{p.d, p.x.PPr}
}

// Style returns the style for a paragraph, or an empty string if it is unset.
func (p Paragraph) Style() string {
	if p.x.PPr != nil && p.x.PPr.PStyle != nil {
		return p.x.PPr.PStyle.ValAttr
	}
	return ""
}

// SetStyle sets the style of a paragraph and is identical to setting it on the
// paragraph's Properties()
func (p Paragraph) SetStyle(s string) {
	p.ensurePPr()
	if s == "" {
		p.x.PPr.PStyle = nil
	} else {
		p.x.PPr.PStyle = wml.NewCT_String()
		p.x.PPr.PStyle.ValAttr = s
	}
}

// AddRun adds a run to a paragraph.
func (p Paragraph) AddRun() Run {
	pc := wml.NewEG_PContent()
	p.x.EG_PContent = append(p.x.EG_PContent, pc)

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
			if rc.Sdt != nil && rc.Sdt.SdtContent != nil {
				for _, rc2 := range rc.Sdt.SdtContent.EG_ContentRunContent {
					if rc2.R != nil {
						ret = append(ret, Run{p.d, rc2.R})
					}
				}
			}
		}
	}
	return ret
}

// InsertRunAfter inserts a run in the paragraph after the relative run.
func (p Paragraph) InsertRunAfter(relativeTo Run) Run {
	return p.insertRun(relativeTo, false)
}

// InsertRunBefore inserts a run in the paragraph before the relative run.
func (p Paragraph) InsertRunBefore(relativeTo Run) Run {
	return p.insertRun(relativeTo, true)
}

func (p Paragraph) insertRun(relativeTo Run, before bool) Run {
	for _, c := range p.x.EG_PContent {
		for i, rc := range c.EG_ContentRunContent {
			if rc.R == relativeTo.X() {
				r := wml.NewCT_R()
				c.EG_ContentRunContent = append(c.EG_ContentRunContent, nil)
				if before {
					copy(c.EG_ContentRunContent[i+1:], c.EG_ContentRunContent[i:])
					c.EG_ContentRunContent[i] = wml.NewEG_ContentRunContent()
					c.EG_ContentRunContent[i].R = r
				} else {
					copy(c.EG_ContentRunContent[i+2:], c.EG_ContentRunContent[i+1:])
					c.EG_ContentRunContent[i+1] = wml.NewEG_ContentRunContent()
					c.EG_ContentRunContent[i+1].R = r
				}
				return Run{p.d, r}

			}
			if rc.Sdt != nil && rc.Sdt.SdtContent != nil {
				for _, rc2 := range rc.Sdt.SdtContent.EG_ContentRunContent {
					if rc2.R == relativeTo.X() {
						r := wml.NewCT_R()
						rc.Sdt.SdtContent.EG_ContentRunContent = append(rc.Sdt.SdtContent.EG_ContentRunContent, nil)
						if before {
							copy(rc.Sdt.SdtContent.EG_ContentRunContent[i+1:], rc.Sdt.SdtContent.EG_ContentRunContent[i:])
							rc.Sdt.SdtContent.EG_ContentRunContent[i] = wml.NewEG_ContentRunContent()
							rc.Sdt.SdtContent.EG_ContentRunContent[i].R = r
						} else {
							copy(rc.Sdt.SdtContent.EG_ContentRunContent[i+2:], rc.Sdt.SdtContent.EG_ContentRunContent[i+1:])
							rc.Sdt.SdtContent.EG_ContentRunContent[i+1] = wml.NewEG_ContentRunContent()
							rc.Sdt.SdtContent.EG_ContentRunContent[i+1].R = r
						}
						return Run{p.d, r}
					}
				}
			}
		}
	}
	return p.AddRun()
}

// AddHyperLink adds a new hyperlink to a parapgraph.
func (p Paragraph) AddHyperLink() HyperLink {
	pc := wml.NewEG_PContent()
	p.x.EG_PContent = append(p.x.EG_PContent, pc)

	pc.Hyperlink = wml.NewCT_Hyperlink()
	return HyperLink{p.d, pc.Hyperlink}
}

// AddBookmark adds a bookmark to a document that can then be used from a hyperlink. Name is a document
// unique name that identifies the bookmark so it can be referenced from hyperlinks.
func (p Paragraph) AddBookmark(name string) Bookmark {
	pc := wml.NewEG_PContent()
	rc := wml.NewEG_ContentRunContent()
	pc.EG_ContentRunContent = append(pc.EG_ContentRunContent, rc)

	relt := wml.NewEG_RunLevelElts()
	rc.EG_RunLevelElts = append(rc.EG_RunLevelElts, relt)

	markEl := wml.NewEG_RangeMarkupElements()
	bmStart := wml.NewCT_Bookmark()
	markEl.BookmarkStart = bmStart
	relt.EG_RangeMarkupElements = append(relt.EG_RangeMarkupElements, markEl)

	markEl = wml.NewEG_RangeMarkupElements()
	markEl.BookmarkEnd = wml.NewCT_MarkupRange()

	relt.EG_RangeMarkupElements = append(relt.EG_RangeMarkupElements, markEl)

	p.x.EG_PContent = append(p.x.EG_PContent, pc)

	bm := Bookmark{bmStart}
	bm.SetName(name)
	return bm
}

// SetNumberingLevel sets the numbering level of a paragraph.  If used, then the
// NumberingDefinition must also be set via SetNumberingDefinition or
// SetNumberingDefinitionByID.
func (p Paragraph) SetNumberingLevel(listLevel int) {
	p.ensurePPr()
	if p.x.PPr.NumPr == nil {
		p.x.PPr.NumPr = wml.NewCT_NumPr()
	}
	lvl := wml.NewCT_DecimalNumber()
	lvl.ValAttr = int64(listLevel)
	p.x.PPr.NumPr.Ilvl = lvl
}

// SetNumberingDefinition sets the numbering definition ID via a NumberingDefinition
// defined in numbering.xml
func (p Paragraph) SetNumberingDefinition(nd NumberingDefinition) {
	p.ensurePPr()
	if p.x.PPr.NumPr == nil {
		p.x.PPr.NumPr = wml.NewCT_NumPr()
	}
	lvl := wml.NewCT_DecimalNumber()

	numID := int64(-1)
	for _, n := range p.d.Numbering.x.Num {
		if n.AbstractNumId != nil && n.AbstractNumId.ValAttr == nd.AbstractNumberID() {
			numID = n.NumIdAttr
		}
	}
	if numID == -1 {
		num := wml.NewCT_Num()
		p.d.Numbering.x.Num = append(p.d.Numbering.x.Num, num)
		num.NumIdAttr = int64(len(p.d.Numbering.x.Num))
		num.AbstractNumId = wml.NewCT_DecimalNumber()
		num.AbstractNumId.ValAttr = nd.AbstractNumberID()
	}

	lvl.ValAttr = numID
	p.x.PPr.NumPr.NumId = lvl
}

// SetNumberingDefinitionByID sets the numbering definition ID directly, which must
// match an ID defined in numbering.xml
func (p Paragraph) SetNumberingDefinitionByID(abstractNumberID int64) {
	p.ensurePPr()
	if p.x.PPr.NumPr == nil {
		p.x.PPr.NumPr = wml.NewCT_NumPr()
	}
	lvl := wml.NewCT_DecimalNumber()
	lvl.ValAttr = int64(abstractNumberID)
	p.x.PPr.NumPr.NumId = lvl
}
