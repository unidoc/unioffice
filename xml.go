// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package unioffice

import "encoding/xml"

// NeedsSpacePreserve returns true if the string has leading or trailing space.
func NeedsSpacePreserve(s string) bool {
	if len(s) == 0 {
		return false
	}
	switch s[0] {
	case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
		return true
	}
	switch s[len(s)-1] {
	case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
		return true
	}
	return false
}

// AddPreserveSpaceAttr adds an xml:space="preserve" attribute to a start
// element if it is required for the string s.
func AddPreserveSpaceAttr(se *xml.StartElement, s string) {
	if NeedsSpacePreserve(s) {
		se.Attr = append(se.Attr, xml.Attr{Name: xml.Name{Local: "xml:space"}, Value: "preserve"})
	}
}
