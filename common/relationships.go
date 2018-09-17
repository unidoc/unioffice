// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package common

import (
	"fmt"
	"strings"

	"baliance.com/gooxml"
	"baliance.com/gooxml/schema/soo/pkg/relationships"
)

// Relationships represents a .rels file.
type Relationships struct {
	x *relationships.Relationships
}

// NewRelationships creates a new relationship wrapper.
func NewRelationships() Relationships {
	return Relationships{x: relationships.NewRelationships()}
}

// X returns the underlying raw XML data.
func (r Relationships) X() *relationships.Relationships {
	return r.x
}

// IsEmpty returns true if there are no relationships.
func (r Relationships) IsEmpty() bool {
	return r.x == nil || len(r.x.Relationship) == 0
}

// Clear removes any existing relationships.
func (r Relationships) Clear() {
	r.x.Relationship = nil
}

// FindRIDForN returns the relationship ID for the i'th relationship of type t.
func (r Relationships) FindRIDForN(i int, t string) string {
	for _, rel := range r.x.CT_Relationships.Relationship {
		if rel.TypeAttr == t {
			if i == 0 {
				return rel.IdAttr
			}
			i--
		}
	}
	return ""
}

// AddAutoRelationship adds a relationship with an automatically generated
// filename based off of the type. It should be preferred over AddRelationship
// to ensure consistent filenames are maintained.
func (r Relationships) AddAutoRelationship(dt gooxml.DocType, src string, idx int, ctype string) Relationship {
	return r.AddRelationship(gooxml.RelativeFilename(dt, src, ctype, idx), ctype)
}

// AddRelationship adds a relationship.
func (r Relationships) AddRelationship(target, ctype string) Relationship {
	if !strings.HasPrefix(ctype, "http://") {
		gooxml.Log("relationship type %s should start with 'http://'", ctype)
	}
	rel := relationships.NewRelationship()
	nextID := len(r.x.Relationship) + 1
	used := map[string]struct{}{}

	// identify IDs in  use
	for _, exRel := range r.x.Relationship {
		used[exRel.IdAttr] = struct{}{}
	}
	// find the next ID that is unused
	for _, ok := used[fmt.Sprintf("rId%d", nextID)]; ok; _, ok = used[fmt.Sprintf("rId%d", nextID)] {
		nextID++

	}
	rel.IdAttr = fmt.Sprintf("rId%d", nextID)
	rel.TargetAttr = target
	rel.TypeAttr = ctype
	r.x.Relationship = append(r.x.Relationship, rel)
	return Relationship{rel}
}

// Hyperlink is just an appropriately configured relationship.
type Hyperlink Relationship

// AddHyperlink adds an external hyperlink relationship.
func (r Relationships) AddHyperlink(target string) Hyperlink {
	rel := r.AddRelationship(target, gooxml.HyperLinkType)
	rel.x.TargetModeAttr = relationships.ST_TargetModeExternal
	return Hyperlink(rel)
}

// Relationships returns a slice of all of the relationships.
func (r Relationships) Relationships() []Relationship {
	ret := []Relationship{}
	for _, x := range r.x.Relationship {
		ret = append(ret, Relationship{x})
	}
	return ret
}
