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

// ST_HexColor is a union type
type ST_HexColor struct {
	ST_HexColorAuto ST_HexColorAuto
	ST_HexColorRGB  *string
}

func (m *ST_HexColor) Validate() error {
	return m.ValidateWithPath("")
}
func (m *ST_HexColor) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_HexColorAuto != ST_HexColorAutoUnset {
		mems = append(mems, "ST_HexColorAuto")
	}
	if m.ST_HexColorRGB != nil {
		mems = append(mems, "ST_HexColorRGB")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}
func (m ST_HexColor) String() string {
	if m.ST_HexColorAuto != ST_HexColorAutoUnset {
		return m.ST_HexColorAuto.String()
	}
	if m.ST_HexColorRGB != nil {
		return fmt.Sprintf("%v", *m.ST_HexColorRGB)
	}
	return ""
}
