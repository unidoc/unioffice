// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package chart

// Axis is the interface implemented by different axes when assigning to a
// chart.
type Axis interface {
	AxisID() uint32
}

type nullAxis byte

func (n nullAxis) AxisID() uint32 {
	return 0
}

// NullAxis is a null axis with an ID of zero
var NullAxis Axis = nullAxis(0)
