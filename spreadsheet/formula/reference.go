// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package formula

// ReferenceType is a type of reference
//go:generate stringer -type=ReferenceType
type ReferenceType byte

const (
	ReferenceTypeInvalid ReferenceType = iota
	ReferenceTypeCell
	ReferenceTypeHorizontalRange
	ReferenceTypeVerticalRange
	ReferenceTypeNamedRange
	ReferenceTypeRange
	ReferenceTypeSheet
)

type Reference struct {
	Type  ReferenceType
	Value string
}

var ReferenceInvalid = Reference{Type: ReferenceTypeInvalid}

func MakeRangeReference(ref string) Reference {
	return Reference{Type: ReferenceTypeRange, Value: ref}
}
