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

// ST_ModelId is a union type
type ST_ModelId struct {
	Int32   *int32
	ST_Guid *string
}

func (m *ST_ModelId) Validate() error {
	return m.ValidateWithPath("")
}

func (m *ST_ModelId) ValidateWithPath(path string) error {
	mems := []string{}
	if m.Int32 != nil {
		mems = append(mems, "Int32")
	}
	if m.ST_Guid != nil {
		mems = append(mems, "ST_Guid")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_ModelId) String() string {
	if m.Int32 != nil {
		return fmt.Sprintf("%v", *m.Int32)
	}
	if m.ST_Guid != nil {
		return fmt.Sprintf("%v", *m.ST_Guid)
	}
	return ""
}
