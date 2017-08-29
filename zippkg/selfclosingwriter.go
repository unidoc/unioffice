// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package zippkg

import (
	"bytes"
	"io"
)

// SelfClosingWriter wraps a writer and replaces XML tags of the
// type <foo></foo> with <foo/>
type SelfClosingWriter struct {
	W io.Writer
}

var closeTag = []byte{'/', '>'}

func (s SelfClosingWriter) Write(b []byte) (int, error) {
	writeStart := 0
	n := 0
	for i := 0; i < len(b)-2; i++ {
		// found an empty tag "></"

		// find the previous tag 'FOO' of '<FOO></FOO>'
		if b[i] == '>' && b[i+1] == '<' && b[i+2] == '/' {
			prevTag := []byte{}
			et := i
			for j := i; j >= 0; j-- {
				if b[j] == ' ' {
					et = j
				} else if b[j] == '<' {
					prevTag = b[j+1 : et]
					break
				}
			}
			nextTag := []byte{}
			for j := i + 3; j < len(b); j++ {
				if b[j] == '>' {
					nextTag = b[i+3 : j]
					break
				}
			}

			// if previous and next tag are equal which catches cases of
			// <a><b>foo</b></a>
			if !bytes.Equal(prevTag, nextTag) {
				continue
			}

			// write up to the start of the tag
			c, err := s.W.Write(b[writeStart:i])
			if err != nil {
				return n + c, err
			}
			n += c

			_, err = s.W.Write(closeTag)
			if err != nil {
				return n, err
			}

			// pretend we wrote close tag
			n += 3

			// skip over the remaining close tag
			for j := i + 2; j < len(b) && b[j] != '>'; j++ {
				n++
				writeStart = j + 2
				i = writeStart
			}
		}
	}
	c, err := s.W.Write(b[writeStart:])
	return c + n, err
}
