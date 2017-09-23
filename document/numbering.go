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
	"baliance.com/gooxml/schema/soo/ofc/sharedTypes"
	"baliance.com/gooxml/schema/soo/wml"
)

// Numbering is the document wide numbering styles contained in numbering.xml.
type Numbering struct {
	x *wml.Numbering
}

// NewNumbering constructs a new numbering.
func NewNumbering() Numbering {
	n := wml.NewNumbering()
	return Numbering{n}
}

// X returns the inner wrapped XML type.
func (n Numbering) X() *wml.Numbering {
	return n.x
}

// Clear resets the numbering.
func (n Numbering) Clear() {
	n.x.AbstractNum = nil
	n.x.Num = nil
	n.x.NumIdMacAtCleanup = nil
	n.x.NumPicBullet = nil
}

// InitializeDefault constructs a default numbering.
func (n Numbering) InitializeDefault() {
	abs := wml.NewCT_AbstractNum()
	n.x.AbstractNum = append(n.x.AbstractNum, abs)
	abs.AbstractNumIdAttr = 1
	const indentStart = 432
	const indentDelta = 144
	for i := 0; i < 9; i++ {
		lvl := wml.NewCT_Lvl()
		lvl.IlvlAttr = int64(i)
		lvl.Start = wml.NewCT_DecimalNumber()
		lvl.Start.ValAttr = 1

		lvl.PStyle = wml.NewCT_String()
		lvl.PStyle.ValAttr = fmt.Sprintf("Heading%d", i+1)

		lvl.NumFmt = wml.NewCT_NumFmt()
		lvl.NumFmt.ValAttr = wml.ST_NumberFormatNone

		lvl.Suff = wml.NewCT_LevelSuffix()
		lvl.Suff.ValAttr = wml.ST_LevelSuffixNothing

		lvl.LvlText = wml.NewCT_LevelText()
		lvl.LvlText.ValAttr = gooxml.String("")

		lvl.LvlJc = wml.NewCT_Jc()
		lvl.LvlJc.ValAttr = wml.ST_JcLeft

		lvl.PPr = wml.NewCT_PPrGeneral()
		lvl.PPr.Tabs = wml.NewCT_Tabs()
		tab := wml.NewCT_TabStop()
		tab.ValAttr = wml.ST_TabJcNum

		indent := int64(i*indentDelta + indentStart)
		tab.PosAttr.Int64 = gooxml.Int64(indent)
		lvl.PPr.Tabs.Tab = append(lvl.PPr.Tabs.Tab, tab)
		lvl.PPr.Ind = wml.NewCT_Ind()
		lvl.PPr.Ind.LeftAttr = &wml.ST_SignedTwipsMeasure{}
		lvl.PPr.Ind.LeftAttr.Int64 = gooxml.Int64(indent)
		lvl.PPr.Ind.HangingAttr = &sharedTypes.ST_TwipsMeasure{}
		lvl.PPr.Ind.HangingAttr.ST_UnsignedDecimalNumber = gooxml.Uint64(uint64(indent))

		abs.Lvl = append(abs.Lvl, lvl)
	}
	num := wml.NewCT_Num()
	num.NumIdAttr = 1
	num.AbstractNumId = wml.NewCT_DecimalNumber()
	num.AbstractNumId.ValAttr = 1
	n.x.Num = append(n.x.Num, num)
}
