// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

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
