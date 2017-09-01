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

// ST_SignedTwipsMeasure is a union type
type ST_SignedTwipsMeasure struct {
	Int64               *int64
	ST_UniversalMeasure *string
}

func (m *ST_SignedTwipsMeasure) Validate() error {
	return m.ValidateWithPath("")
}

func (m *ST_SignedTwipsMeasure) ValidateWithPath(path string) error {
	mems := []string{}
	if m.Int64 != nil {
		mems = append(mems, "Int64")
	}
	if m.ST_UniversalMeasure != nil {
		mems = append(mems, "ST_UniversalMeasure")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_SignedTwipsMeasure) String() string {
	if m.Int64 != nil {
		return fmt.Sprintf("%v", *m.Int64)
	}
	if m.ST_UniversalMeasure != nil {
		return fmt.Sprintf("%v", *m.ST_UniversalMeasure)
	}
	return ""
}
