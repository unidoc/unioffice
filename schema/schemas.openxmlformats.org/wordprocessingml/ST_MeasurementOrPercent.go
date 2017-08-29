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

// ST_MeasurementOrPercent is a union type
type ST_MeasurementOrPercent struct {
	ST_DecimalNumberOrPercent *ST_DecimalNumberOrPercent
	ST_UniversalMeasure       *string
}

func (m *ST_MeasurementOrPercent) Validate() error {
	return m.ValidateWithPath("")
}
func (m *ST_MeasurementOrPercent) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_DecimalNumberOrPercent != nil {
		if err := m.ST_DecimalNumberOrPercent.ValidateWithPath(path + "/ST_DecimalNumberOrPercent"); err != nil {
			return err
		}
		mems = append(mems, "ST_DecimalNumberOrPercent")
	}
	if m.ST_UniversalMeasure != nil {
		mems = append(mems, "ST_UniversalMeasure")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}
func (m ST_MeasurementOrPercent) String() string {
	if m.ST_DecimalNumberOrPercent != nil {
		return m.ST_DecimalNumberOrPercent.String()
	}
	if m.ST_UniversalMeasure != nil {
		return fmt.Sprintf("%v", *m.ST_UniversalMeasure)
	}
	return ""
}
