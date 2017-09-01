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

// ST_Overlap is a union type
type ST_Overlap struct {
	ST_OverlapPercent *string
	ST_OverlapByte    *int8
}

func (m *ST_Overlap) Validate() error {
	return m.ValidateWithPath("")
}

func (m *ST_Overlap) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_OverlapPercent != nil {
		mems = append(mems, "ST_OverlapPercent")
	}
	if m.ST_OverlapByte != nil {
		mems = append(mems, "ST_OverlapByte")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_Overlap) String() string {
	if m.ST_OverlapPercent != nil {
		return fmt.Sprintf("%v", *m.ST_OverlapPercent)
	}
	if m.ST_OverlapByte != nil {
		return fmt.Sprintf("%v", *m.ST_OverlapByte)
	}
	return ""
}
