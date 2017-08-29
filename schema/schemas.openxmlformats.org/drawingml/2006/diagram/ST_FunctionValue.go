// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"fmt"
)

// ST_FunctionValue is a union type
type ST_FunctionValue struct {
	Int32               *int32
	Bool                *bool
	ST_Direction        ST_Direction
	ST_HierBranchStyle  ST_HierBranchStyle
	ST_AnimOneStr       ST_AnimOneStr
	ST_AnimLvlStr       ST_AnimLvlStr
	ST_ResizeHandlesStr ST_ResizeHandlesStr
}

func (m *ST_FunctionValue) Validate() error {
	return m.ValidateWithPath("")
}
func (m *ST_FunctionValue) ValidateWithPath(path string) error {
	mems := []string{}
	if m.Int32 != nil {
		mems = append(mems, "Int32")
	}
	if m.Bool != nil {
		mems = append(mems, "Bool")
	}
	if m.ST_Direction != ST_DirectionUnset {
		mems = append(mems, "ST_Direction")
	}
	if m.ST_HierBranchStyle != ST_HierBranchStyleUnset {
		mems = append(mems, "ST_HierBranchStyle")
	}
	if m.ST_AnimOneStr != ST_AnimOneStrUnset {
		mems = append(mems, "ST_AnimOneStr")
	}
	if m.ST_AnimLvlStr != ST_AnimLvlStrUnset {
		mems = append(mems, "ST_AnimLvlStr")
	}
	if m.ST_ResizeHandlesStr != ST_ResizeHandlesStrUnset {
		mems = append(mems, "ST_ResizeHandlesStr")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}
func (m ST_FunctionValue) String() string {
	if m.Int32 != nil {
		return fmt.Sprintf("%v", *m.Int32)
	}
	if m.Bool != nil {
		return fmt.Sprintf("%v", *m.Bool)
	}
	if m.ST_Direction != ST_DirectionUnset {
		return m.ST_Direction.String()
	}
	if m.ST_HierBranchStyle != ST_HierBranchStyleUnset {
		return m.ST_HierBranchStyle.String()
	}
	if m.ST_AnimOneStr != ST_AnimOneStrUnset {
		return m.ST_AnimOneStr.String()
	}
	if m.ST_AnimLvlStr != ST_AnimLvlStrUnset {
		return m.ST_AnimLvlStr.String()
	}
	if m.ST_ResizeHandlesStr != ST_ResizeHandlesStrUnset {
		return m.ST_ResizeHandlesStr.String()
	}
	return ""
}
