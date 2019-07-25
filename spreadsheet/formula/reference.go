// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

// ReferenceType is a type of reference
//go:generate stringer -type=ReferenceType
type ReferenceType byte

const (
	ReferenceTypeInvalid ReferenceType = iota
	ReferenceTypeCell
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
