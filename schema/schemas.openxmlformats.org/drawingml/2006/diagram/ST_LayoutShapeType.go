// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"fmt"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
)

// ST_LayoutShapeType is a union type
type ST_LayoutShapeType struct {
	ST_ShapeType       drawingml.ST_ShapeType
	ST_OutputShapeType ST_OutputShapeType
}

func (m *ST_LayoutShapeType) Validate() error {
	return m.ValidateWithPath("")
}

func (m *ST_LayoutShapeType) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_ShapeType != drawingml.ST_ShapeTypeUnset {
		mems = append(mems, "ST_ShapeType")
	}
	if m.ST_OutputShapeType != ST_OutputShapeTypeUnset {
		mems = append(mems, "ST_OutputShapeType")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_LayoutShapeType) String() string {
	if m.ST_ShapeType != drawingml.ST_ShapeTypeUnset {
		return m.ST_ShapeType.String()
	}
	if m.ST_OutputShapeType != ST_OutputShapeTypeUnset {
		return m.ST_OutputShapeType.String()
	}
	return ""
}
