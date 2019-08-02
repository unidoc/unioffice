// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

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
