// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package unioffice

import "fmt"

// Float32 returns a copy of v as a pointer.
func Float32(v float32) *float32 {
	x := v
	return &x
}

// Float64 returns a copy of v as a pointer.
func Float64(v float64) *float64 {
	x := v
	return &x
}

// Uint64 returns a copy of v as a pointer.
func Uint64(v uint64) *uint64 {
	x := v
	return &x
}

// Uint32 returns a copy of v as a pointer.
func Uint32(v uint32) *uint32 {
	x := v
	return &x
}

// Uint16 returns a copy of v as a pointer.
func Uint16(v uint16) *uint16 {
	x := v
	return &x
}

// Uint8 returns a copy of v as a pointer.
func Uint8(v uint8) *uint8 {
	x := v
	return &x
}

// Int64 returns a copy of v as a pointer.
func Int64(v int64) *int64 {
	x := v
	return &x
}

// Int32 returns a copy of v as a pointer.
func Int32(v int32) *int32 {
	x := v
	return &x
}

// Int8 returns a copy of v as a pointer.
func Int8(v int8) *int8 {
	x := v
	return &x
}

// Bool returns a copy of v as a pointer.
func Bool(v bool) *bool {
	x := v
	return &x
}

// String returns a copy of v as a pointer.
func String(v string) *string {
	x := v
	return &x
}

// Stringf formats according to a format specifier and returns a pointer to the
// resulting string.
func Stringf(f string, args ...interface{}) *string {
	x := fmt.Sprintf(f, args...)
	return &x
}
