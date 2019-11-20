// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package common_test

import (
	"bytes"
	"encoding/xml"
	"os"
	"testing"

	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/testhelper"
	"github.com/unidoc/unioffice/zippkg"
)

func TestRelationshipsUnmarshal(t *testing.T) {
	f, err := os.Open("testdata/rels.xml")
	if err != nil {
		t.Fatalf("error reading content types file")
	}
	dec := xml.NewDecoder(f)
	r := common.NewRelationships()
	if err := dec.Decode(r.X()); err != nil {
		t.Errorf("error decoding content types: %s", err)
	}
	got := &bytes.Buffer{}
	enc := xml.NewEncoder(zippkg.SelfClosingWriter{W: got})
	if err := enc.Encode(r.X()); err != nil {
		t.Errorf("error encoding content types: %s", err)
	}

	testhelper.CompareGoldenXML(t, "rels.xml", got.Bytes())
}

func TestRelationshipsCreation(t *testing.T) {
	f, err := os.Open("testdata/rels.xml")
	if err != nil {
		t.Fatalf("error reading content types file")
	}
	dec := xml.NewDecoder(f)
	r := common.NewRelationships()
	if err := dec.Decode(r.X()); err != nil {
		t.Errorf("error decoding content types: %s", err)
	}

	rel := r.AddRelationship("foo", "http://bar")

	// testdata/rels.xml contains rels with IDs 2,3,4.  We used to
	// just use the len, but issue #198 pointed out a problem with
	// that.  Now it should check and see that rId4 is in use, and
	// go to the next free ID.
	exp := "rId5"
	if rel.ID() != exp {
		t.Errorf("expected %s, got %s", exp, rel.ID())
	}
	// 6 is unused, so we should get it next
	rel = r.AddRelationship("foo", "http://bar")
	exp = "rId6"
	if rel.ID() != exp {
		t.Errorf("expected %s, got %s", exp, rel.ID())
	}
}

func TestRelationshipsCreationUnordered(t *testing.T) {
	r := common.NewRelationships()
	r.AddRelationship("foo", "http://bar").X().IdAttr = "rId1"
	r.AddRelationship("foo", "http://bar").X().IdAttr = "rId3"
	r.AddRelationship("foo", "http://bar").X().IdAttr = "rId5"
	r.AddRelationship("foo", "http://bar").X().IdAttr = "rId7"

	// len is 4, but id 5 is used, so we should get 6 next
	rel := r.AddRelationship("foo", "http://bar")

	exp := "rId6"
	if rel.ID() != exp {
		t.Errorf("expected %s, got %s", exp, rel.ID())
	}

	// would get 7, but it's in use so we get 8 now
	rel = r.AddRelationship("foo", "http://bar")
	exp = "rId8"
	if rel.ID() != exp {
		t.Errorf("expected %s, got %s", exp, rel.ID())
	}

}

func TestRelationshipsRemoval(t *testing.T) {
	r := common.NewRelationships()
	r.AddRelationship("foo1", "http://bar")
	r.AddRelationship("foo2", "http://bar")
	r.AddRelationship("foo3", "http://bar")

	if len(r.Relationships()) != 3 {
		t.Errorf("expected 3, got %d", len(r.Relationships()))
	}
	r.Remove(r.Relationships()[0])

	if len(r.Relationships()) != 2 {
		t.Errorf("expected 2, got %d", len(r.Relationships()))
	}
	if got := r.Relationships()[0].Target(); got != "foo2" {
		t.Errorf("expected foo2, got %s", got)
	}
	if got := r.Relationships()[1].Target(); got != "foo3" {
		t.Errorf("expected foo3, got %s", got)
	}
	r.Remove(r.Relationships()[1])

	if len(r.Relationships()) != 1 {
		t.Errorf("expected 1, got %d", len(r.Relationships()))
	}
	if got := r.Relationships()[0].Target(); got != "foo2" {
		t.Errorf("expected foo2, got %s", got)
	}

	r.Remove(r.Relationships()[0])
	if len(r.Relationships()) != 0 {
		t.Errorf("expected 0, got %d", len(r.Relationships()))
	}
}

func TestCopyRelationship(t *testing.T) {
	r := common.NewRelationships()
	r.AddRelationship("foo1", "http://bar")
	r.AddRelationship("foo2", "http://bar")
	r.AddRelationship("foo3", "http://bar")

	if len(r.Relationships()) != 3 {
		t.Errorf("expected 3, got %d", len(r.Relationships()))
	}

	copied, ok := r.CopyRelationship(r.Relationships()[1].ID())
	if !ok {
		t.Errorf("expected true, got %v", ok)
	}

	if len(r.Relationships()) != 4 {
		t.Errorf("expected 4, got %d", len(r.Relationships()))
	}

	if got := copied.Target(); got != "foo2" {
		t.Errorf("expected foo2, got %s", got)
	}

	_, ok = r.CopyRelationship("qweqwe")
	if ok {
		t.Errorf("expected false, got %v", ok)
	}
}
