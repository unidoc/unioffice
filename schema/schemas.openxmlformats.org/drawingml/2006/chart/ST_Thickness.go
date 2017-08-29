// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"fmt"
)

// ST_Thickness is a union type
type ST_Thickness struct {
	ST_ThicknessPercent *string
	Uint32              *uint32
}

func (m *ST_Thickness) Validate() error {
	return m.ValidateWithPath("")
}
func (m *ST_Thickness) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_ThicknessPercent != nil {
		mems = append(mems, "ST_ThicknessPercent")
	}
	if m.Uint32 != nil {
		mems = append(mems, "Uint32")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}
func (m ST_Thickness) String() string {
	if m.ST_ThicknessPercent != nil {
		return fmt.Sprintf("%v", *m.ST_ThicknessPercent)
	}
	if m.Uint32 != nil {
		return fmt.Sprintf("%v", *m.Uint32)
	}
	return ""
}
