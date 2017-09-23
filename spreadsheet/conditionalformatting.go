// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import "baliance.com/gooxml/schema/soo/sml"

// ConditionalFormatting controls the formatting styles and rules for a range of
// cells with the same conditional formatting.
type ConditionalFormatting struct {
	x *sml.CT_ConditionalFormatting
}

// X returns the inner wrapped XML type.
func (c ConditionalFormatting) X() *sml.CT_ConditionalFormatting {
	return c.x
}

// AddRule adds and returns a new rule that can be configured.
func (c ConditionalFormatting) AddRule() ConditionalFormattingRule {
	rule := sml.NewCT_CfRule()
	c.x.CfRule = append(c.x.CfRule, rule)
	r := ConditionalFormattingRule{rule}
	r.InitializeDefaults()
	r.SetPriority(int32(len(c.x.CfRule) + 1))
	return r
}
