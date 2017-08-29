// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sharedTypes

import "fmt"

// ST_OnOff is a union type
type ST_OnOff struct {
	Bool      *bool
	ST_OnOff1 ST_OnOff1
}

func (m *ST_OnOff) Validate() error {
	return m.ValidateWithPath("")
}
func (m *ST_OnOff) ValidateWithPath(path string) error {
	mems := []string{}
	if m.Bool != nil {
		mems = append(mems, "Bool")
	}
	if m.ST_OnOff1 != ST_OnOff1Unset {
		mems = append(mems, "ST_OnOff1")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}
func (m ST_OnOff) String() string {
	if m.Bool != nil {
		return fmt.Sprintf("%v", *m.Bool)
	}
	if m.ST_OnOff1 != ST_OnOff1Unset {
		return m.ST_OnOff1.String()
	}
	return ""
}
