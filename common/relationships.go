// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package common

import (
	"fmt"
	"log"
	"strings"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/package/2006/relationships"
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

// AddRelationship adds a relationship.
func (r Relationships) AddRelationship(target, ctype string) Relationship {
	if !strings.HasPrefix(ctype, "http://") {
		log.Printf("relationship type %s should start with 'http://'", ctype)
	}
	rel := relationships.NewRelationship()
	rel.IdAttr = fmt.Sprintf("rId%d", len(r.x.Relationship)+1)
	rel.TargetAttr = target
	rel.TypeAttr = ctype
	r.x.Relationship = append(r.x.Relationship, rel)
	return Relationship{rel}
}

// Relationships returns a slice of all of the relationships.
func (r Relationships) Relationships() []Relationship {
	ret := []Relationship{}
	for _, x := range r.x.Relationship {
		ret = append(ret, Relationship{x})
	}
	return ret
}
