// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.
package spreadsheet

import (
	"fmt"

	"baliance.com/gooxml"
	"baliance.com/gooxml/schema/soo/sml"
)

// SharedStrings is a shared strings table, where string data can be placed
// outside of the sheet contents and referenced from a sheet.
type SharedStrings struct {
	x         *sml.Sst
	cachedIDs map[string]int
}

// NewSharedStrings constructs a new Shared Strings table.
func NewSharedStrings() SharedStrings {
	return SharedStrings{x: sml.NewSst(),
		cachedIDs: make(map[string]int)}
}

// X returns the inner wrapped XML type.
func (s SharedStrings) X() *sml.Sst {
	return s.x
}

// AddString adds a string to the shared string cache.
func (s SharedStrings) AddString(v string) int {
	if id, ok := s.cachedIDs[v]; ok {
		return id
	}
	rst := sml.NewCT_Rst()
	rst.T = gooxml.String(v)
	s.x.Si = append(s.x.Si, rst)
	id := len(s.x.Si) - 1
	s.cachedIDs[v] = id
	s.x.CountAttr = gooxml.Uint32(uint32(len(s.x.Si)))
	s.x.UniqueCountAttr = s.x.CountAttr
	return id
}

// GetString retrieves a string from the shared strings table by index.
func (s SharedStrings) GetString(id int) (string, error) {
	if id < 0 {
		return "", fmt.Errorf("invalid string index %d, must be > 0", id)
	}
	if id > len(s.x.Si) {
		return "", fmt.Errorf("invalid string index %d, table only has %d values", id, len(s.x.Si))
	}
	si := s.x.Si[id]
	if si.T != nil {
		return *si.T, nil
	}
	return "", nil
}
