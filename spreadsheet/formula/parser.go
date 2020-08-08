// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package formula

import (
	"io"
	"strings"

	"github.com/unidoc/unioffice"
)

//go:generate goyacc -l -o grammar.go  grammar.y
type plex struct {
	nodes  chan *node
	result Expression
}

func (f *plex) Lex(lval *yySymType) int {
	//yyDebug = 3
	yyErrorVerbose = true
	n := <-f.nodes
	if n != nil {
		lval.node = n
		return int(lval.node.token)
	}
	return 0
}

func (f *plex) Error(s string) {
	unioffice.Log("parse error: %s", s)
}

func Parse(r io.Reader) Expression {
	p := &plex{LexReader(r), nil}
	yyParse(p)
	return p.result
}

func ParseString(s string) Expression {
	if s == "" {
		return NewEmptyExpr()
	}
	return Parse(strings.NewReader(s))
}
