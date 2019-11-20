// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package color

import "fmt"

func FromHex(s string) Color {
	if len(s) == 0 {
		return Auto
	}
	if s[0] == '#' {
		s = s[1:]
	}
	//func Sscanf(str string, format string, a ...interface{}) (n int, err error) {
	var r, g, b uint8
	n, _ := fmt.Sscanf(s, "%02x%02x%02x", &r, &g, &b)
	if n == 3 {
		return RGB(r, g, b)
	}
	return Auto
}
