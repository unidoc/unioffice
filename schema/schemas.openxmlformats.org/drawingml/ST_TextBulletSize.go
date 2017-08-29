// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"fmt"
)

// ST_TextBulletSize is a union type
type ST_TextBulletSize struct {
	ST_TextBulletSizePercent *string
	ST_TextBulletSizeDecimal *int32
}

func (m *ST_TextBulletSize) Validate() error {
	return m.ValidateWithPath("")
}
func (m *ST_TextBulletSize) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_TextBulletSizePercent != nil {
		mems = append(mems, "ST_TextBulletSizePercent")
	}
	if m.ST_TextBulletSizeDecimal != nil {
		mems = append(mems, "ST_TextBulletSizeDecimal")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}
func (m ST_TextBulletSize) String() string {
	if m.ST_TextBulletSizePercent != nil {
		return fmt.Sprintf("%v", *m.ST_TextBulletSizePercent)
	}
	if m.ST_TextBulletSizeDecimal != nil {
		return fmt.Sprintf("%v", *m.ST_TextBulletSizeDecimal)
	}
	return ""
}
