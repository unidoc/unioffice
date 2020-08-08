// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

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
