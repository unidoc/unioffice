// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package zippkg

import (
	"bytes"
	"testing"
)

func TestSelfClosing(t *testing.T) {

	td := []struct {
		Input    string
		Expected string
	}{
		{"<test></test>", "<test/>"},
		{"<test> </test>", "<test> </test>"},
		{"<test a=\"123\"></test>", "<test a=\"123\"/>"},
		{`<Default Extension="jpg" ContentType="image/jpg"></Default>`, `<Default Extension="jpg" ContentType="image/jpg"/>`},
		{`<Override ContentType="application/vnd.openxmlformats-officedocument.spreadsheetml.styles+xml" PartName="/xl/styles.xml"></Override>`,
			`<Override ContentType="application/vnd.openxmlformats-officedocument.spreadsheetml.styles+xml" PartName="/xl/styles.xml"/>`},
		{"<TestStruct><Foo>bar</Foo></TestStruct>", "<TestStruct><Foo>bar</Foo></TestStruct>"},
		{"<test></test><a></a><b></b>", "<test/><a/><b/>"},
	}

	for _, tc := range td {
		buf := bytes.Buffer{}
		w := SelfClosingWriter{&buf}
		n, err := w.Write([]byte(tc.Input))
		if err != nil {
			t.Errorf("error writing: %s", err)
		}
		if n != len(tc.Input) {
			t.Errorf("expeced to write %d bytes, wrote %d", len(tc.Input), n)
		}

		got := buf.String()
		if got != tc.Expected {
			t.Errorf("expected write(\"%s\") = \"%s\", got \"%s\"", tc.Input, tc.Expected, got)
		}
	}
}
