// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package unioffice_test

import (
	"testing"

	"github.com/unidoc/unioffice"
)

func TestFloat32(t *testing.T) {
	exp := float32(1.234)
	got := unioffice.Float32(exp)
	if *got != exp {
		t.Errorf("expected %f, got %f", exp, *got)
	}
}

func TestFloat64(t *testing.T) {
	exp := 1.234
	got := unioffice.Float64(exp)
	if *got != exp {
		t.Errorf("expected %f, got %f", exp, *got)
	}
}

func TestUint64(t *testing.T) {
	exp := uint64(123)
	got := unioffice.Uint64(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestUint32(t *testing.T) {
	exp := uint32(123)
	got := unioffice.Uint32(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestInt64(t *testing.T) {
	exp := int64(123)
	got := unioffice.Int64(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestInt32(t *testing.T) {
	exp := int32(123)
	got := unioffice.Int32(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestInt8(t *testing.T) {
	exp := int8(123)
	got := unioffice.Int8(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestBool(t *testing.T) {
	exp := bool(true)
	got := unioffice.Bool(exp)
	if *got != exp {
		t.Errorf("expected %v, got %v", exp, *got)
	}
}

func TestString(t *testing.T) {
	exp := "foo"
	got := unioffice.String(exp)
	if *got != exp {
		t.Errorf("expected %s, got %s", exp, *got)
	}
}

func TestUint8(t *testing.T) {
	exp := uint8(123)
	got := unioffice.Uint8(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestUint16(t *testing.T) {
	exp := uint16(123)
	got := unioffice.Uint16(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestStringf(t *testing.T) {
	exp := "foobar123"
	got := unioffice.Stringf("foo%s%d", "bar", 123)
	if *got != exp {
		t.Errorf("expected %s, got %s", exp, *got)
	}
}
