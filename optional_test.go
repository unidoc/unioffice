// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package gooxml_test

import (
	"testing"

	"baliance.com/gooxml"
)

func TestFloat32(t *testing.T) {
	exp := float32(1.234)
	got := gooxml.Float32(exp)
	if *got != exp {
		t.Errorf("expected %f, got %f", exp, *got)
	}
}

func TestFloat64(t *testing.T) {
	exp := 1.234
	got := gooxml.Float64(exp)
	if *got != exp {
		t.Errorf("expected %f, got %f", exp, *got)
	}
}

func TestUint64(t *testing.T) {
	exp := uint64(123)
	got := gooxml.Uint64(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestUint32(t *testing.T) {
	exp := uint32(123)
	got := gooxml.Uint32(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestInt64(t *testing.T) {
	exp := int64(123)
	got := gooxml.Int64(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestInt32(t *testing.T) {
	exp := int32(123)
	got := gooxml.Int32(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestInt8(t *testing.T) {
	exp := int8(123)
	got := gooxml.Int8(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestBool(t *testing.T) {
	exp := bool(true)
	got := gooxml.Bool(exp)
	if *got != exp {
		t.Errorf("expected %v, got %v", exp, *got)
	}
}

func TestString(t *testing.T) {
	exp := "foo"
	got := gooxml.String(exp)
	if *got != exp {
		t.Errorf("expected %s, got %s", exp, *got)
	}
}

func TestUint8(t *testing.T) {
	exp := uint8(123)
	got := gooxml.Uint8(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestUint16(t *testing.T) {
	exp := uint16(123)
	got := gooxml.Uint16(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestStringf(t *testing.T) {
	exp := "foobar123"
	got := gooxml.Stringf("foo%s%d", "bar", 123)
	if *got != exp {
		t.Errorf("expected %s, got %s", exp, *got)
	}
}
