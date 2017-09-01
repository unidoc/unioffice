// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"fmt"
)

// ST_DecimalNumberOrPercent is a union type
type ST_DecimalNumberOrPercent struct {
	ST_UnqualifiedPercentage *int64
	ST_Percentage            *string
}

func (m *ST_DecimalNumberOrPercent) Validate() error {
	return m.ValidateWithPath("")
}

func (m *ST_DecimalNumberOrPercent) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_UnqualifiedPercentage != nil {
		mems = append(mems, "ST_UnqualifiedPercentage")
	}
	if m.ST_Percentage != nil {
		mems = append(mems, "ST_Percentage")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_DecimalNumberOrPercent) String() string {
	if m.ST_UnqualifiedPercentage != nil {
		return fmt.Sprintf("%v", *m.ST_UnqualifiedPercentage)
	}
	if m.ST_Percentage != nil {
		return fmt.Sprintf("%v", *m.ST_Percentage)
	}
	return ""
}
