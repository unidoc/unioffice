// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package zippkg_test

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"testing"

	"baliance.com/gooxml/zippkg"
)

type TestStruct struct {
	Foo string
}

func TestMarshal(t *testing.T) {
	buf := &bytes.Buffer{}
	zw := zip.NewWriter(buf)
	f := TestStruct{Foo: "bar"}

	fname := "/test/foo.xml"
	if err := zippkg.MarshalXML(zw, fname, &f); err != nil {
		t.Errorf("expected no error, got %s", err)
	}
	if err := zw.Close(); err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	zr, err := zip.NewReader(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}
	if len(zr.File) != 1 {
		t.Errorf("expected one file in zip, got %d", len(zr.File))
	}
	zf := zr.File[0]
	if zf.Name != fname {
		t.Errorf("expected name = %s, got %s", fname, zf.Name)
	}
	rc, err := zf.Open()
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}
	out, err := ioutil.ReadAll(rc)
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}
	exp := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>` + "\n" + `<TestStruct><Foo>bar</Foo></TestStruct>` + "\r\n"
	if got := string(out); got != exp {
		t.Errorf("expected\n%s\n, got \n%s\n", exp, got)
	}
}
