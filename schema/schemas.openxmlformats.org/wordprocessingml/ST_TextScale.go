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

// ST_TextScale is a union type
type ST_TextScale struct {
	ST_TextScalePercent *string
	ST_TextScaleDecimal *int32
}

func (m *ST_TextScale) Validate() error {
	return m.ValidateWithPath("")
}
func (m *ST_TextScale) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_TextScalePercent != nil {
		mems = append(mems, "ST_TextScalePercent")
	}
	if m.ST_TextScaleDecimal != nil {
		mems = append(mems, "ST_TextScaleDecimal")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}
func (m ST_TextScale) String() string {
	if m.ST_TextScalePercent != nil {
		return fmt.Sprintf("%v", *m.ST_TextScalePercent)
	}
	if m.ST_TextScaleDecimal != nil {
		return fmt.Sprintf("%v", *m.ST_TextScaleDecimal)
	}
	return ""
}
