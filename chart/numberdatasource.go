// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package chart

import (
	"fmt"

	crt "github.com/unidoc/unioffice/schema/soo/dml/chart"
)

type NumberDataSource struct {
	x *crt.CT_NumDataSource
}

func MakeNumberDataSource(x *crt.CT_NumDataSource) NumberDataSource {
	return NumberDataSource{x}
}

func (n NumberDataSource) ensureChoice() {
	if n.x.Choice == nil {
		n.x.Choice = crt.NewCT_NumDataSourceChoice()
	}
}
func (n NumberDataSource) SetReference(s string) {
	n.ensureChoice()
	if n.x.Choice.NumRef == nil {
		n.x.Choice.NumRef = crt.NewCT_NumRef()
	}
	n.x.Choice.NumRef.F = s
}

// CreateEmptyNumberCache creates an empty number cache, which is used sometimes
// to increase file format compatibility.  It should actually contain the
// computed cell data, but just creating an empty one is good enough.
func (n NumberDataSource) CreateEmptyNumberCache() {
	n.ensureChoice()
	if n.x.Choice.NumRef == nil {
		n.x.Choice.NumRef = crt.NewCT_NumRef()
	}
	n.x.Choice.NumLit = nil
	n.x.Choice.NumRef.NumCache = crt.NewCT_NumData()
	n.x.Choice.NumRef.NumCache.PtCount = crt.NewCT_UnsignedInt()
	n.x.Choice.NumRef.NumCache.PtCount.ValAttr = 0
}

// SetValues sets values directly on a source.
func (n NumberDataSource) SetValues(v []float64) {
	n.ensureChoice()
	n.x.Choice.NumRef = nil
	n.x.Choice.NumLit = crt.NewCT_NumData()
	n.x.Choice.NumLit.PtCount = crt.NewCT_UnsignedInt()
	n.x.Choice.NumLit.PtCount.ValAttr = uint32(len(v))

	for i, x := range v {
		n.x.Choice.NumLit.Pt = append(n.x.Choice.NumLit.Pt,
			&crt.CT_NumVal{
				IdxAttr: uint32(i),
				V:       fmt.Sprintf("%g", x)})
	}

}
