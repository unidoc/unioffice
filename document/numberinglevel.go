// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import "baliance.com/gooxml/schema/soo/wml"

// NumberingLevel is the definition for numbering for a particular level within
// a NumberingDefinition.
type NumberingLevel struct {
	x *wml.CT_NumLvl
}

// X returns the inner wrapped XML type.
func (n NumberingLevel) X() *wml.CT_NumLvl {
	return n.x
}

func (n NumberingLevel) ensureLevel() {
	if n.x.Lvl == nil {
		n.x.Lvl = wml.NewCT_Lvl()
	}
}

// SetStartOverride sets the Numbering Level Starting Value Override.
func (n NumberingLevel) SetStartOverride(v int64) {
	n.ensureLevel()
	n.x.StartOverride = wml.NewCT_DecimalNumber()
	n.x.StartOverride.ValAttr = v
}
