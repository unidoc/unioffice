// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

// Based off of http://ieeexplore.ieee.org/document/7335408/

import (
	"fmt"
	"io"
	"strings"
	"sync"
)

var debugLex = false

type tokenType int

func (t tokenType) String() string {
	return yyTokname(int(t))
}

type node struct {
	token tokenType
	val   string
}

func printable(s string) string {
	s = strings.Replace(s, "\n", "\\n", -1)
	s = strings.Replace(s, "\r", "\\r", -1)
	s = strings.Replace(s, "\t", "\\t", -1)
	return s
}
func (n node) String() string {
	return fmt.Sprintf("{%s %s}", n.token, printable(string(n.val)))
}

//go:generate ragel -G2 -Z lexer.rl
//go:generate goimports -w lexer.go
type Lexer struct {
	nodes    chan *node
	lock     sync.Mutex
	injected []chan *node
	peeked   []*node
}

func NewLexer() *Lexer {
	return &Lexer{nodes: make(chan *node)}
}

func LexReader(r io.Reader) chan *node {
	l := NewLexer()
	go l.lex(r)
	return l.nodes
}

func (l *Lexer) emit(typ tokenType, val []byte) {
	if debugLex {
		fmt.Println("emit", typ, printable(string(val)))
	}
	l.nodes <- &node{typ, string(val)}
}

func (l *Lexer) nextRaw() *node {
	for len(l.injected) != 0 {
		n := <-l.injected[len(l.injected)-1]
		if n != nil {
			return n
		}
		l.injected = l.injected[0 : len(l.injected)-1]
	}
	return <-l.nodes
}

func (l *Lexer) Next() *node {
	l.lock.Lock()
	defer l.lock.Unlock()
	if len(l.peeked) > 0 {
		n := l.peeked[0]
		l.peeked = l.peeked[1:]
		return n
	}
	return l.nextRaw()
}
