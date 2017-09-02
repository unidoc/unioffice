// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package gooxml_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml"
)

func TestAddPreserveSpaceAttr(t *testing.T) {
	td := []struct {
		Input   string
		HasAttr bool
	}{
		{"", false},
		{"foo", false},
		{"f o o", false},
		{"foo ", true},
		{" foo ", true},
		{" foo ", true},
		{"\tfoo", true},
		{"\nfoo", true},
	}
	for _, tc := range td {
		se := &xml.StartElement{}
		gooxml.AddPreserveSpaceAttr(se, tc.Input)
		if tc.HasAttr && len(se.Attr) == 0 {
			t.Errorf("expected a preserve space attribute for %s", tc.Input)
		} else if !tc.HasAttr && len(se.Attr) != 0 {
			t.Errorf("expected no preserve space attribute for %s", tc.Input)
		}
		if tc.HasAttr {
			if se.Attr[0].Name.Local != "xml:space" {
				t.Errorf("expected name = xml:space, got %s", se.Attr[0].Name.Local)
			}
			if se.Attr[0].Value != "preserve" {
				t.Errorf("expected name = preserve, got %s", se.Attr[0].Value)
			}
		}
	}
}
