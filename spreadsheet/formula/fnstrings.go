// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

import (
	"bytes"
	"log"
)

func init() {
	RegisterFunction("CONCATENATE", Concatenate)
	RegisterFunction("LEFT", Left)
	RegisterFunction("RIGHT", Right)
}

// Concatenate is an implementation of the Excel CONCATENATE() function.
func Concatenate(args []Result) Result {
	buf := bytes.Buffer{}
	for _, a := range args {
		switch a.Type {
		case ResultTypeString:
			buf.WriteString(a.ValueString)
		case ResultTypeNumber:
			buf.WriteString(a.Value())
		case ResultTypeList:
			// concatenate seems to pull just the last value from a list
			if len(a.ValueList) > 0 {
				res := Concatenate(a.ValueList[len(a.ValueList)-1:])
				buf.WriteString(res.ValueString)
			}
		default:
			log.Printf("unhandled CONCATENATE() argument type %s", a.Type)
		}
	}
	return MakeStringResult(buf.String())
}

// Left implements the Excel LEFT(string,[n]) function which returns the
// leftmost n characters.
func Left(args []Result) Result {
	n := 1
	switch len(args) {
	case 1:
		// no length argument returns the single left-most character
	case 2:
		// second argument must be a number
		if args[1].Type != ResultTypeNumber {
			return MakeErrorResult("LEFT expected number argument")
		}
		// Excel truncates floating points
		n = int(args[1].ValueNumber)
		if n < 0 {
			return MakeErrorResult("LEFT expected number argument >= 0")
		}
		if n == 0 { // empty string
			return MakeStringResult("")
		}
	default:
		return MakeErrorResult("LEFT expected one or two arguments")
	}

	// can't call LEFT on a range
	if args[0].Type == ResultTypeList {
		return MakeErrorResult("LEFT can't be called on a range")
	}
	v := args[0].Value()
	if n > len(v) {
		return MakeStringResult(v)
	}
	return MakeStringResult(v[0:n])

}

// Right implements the Excel RIGHT(string,[n]) function which returns the
// rightmost n characters.
func Right(args []Result) Result {
	n := 1
	switch len(args) {
	case 1:
		// no length argument returns the single right-most character
	case 2:
		// second argument must be a number
		if args[1].Type != ResultTypeNumber {
			return MakeErrorResult("RIGHT expected number argument")
		}
		// Excel truncates floating points
		n = int(args[1].ValueNumber)
		if n < 0 {
			return MakeErrorResult("RIGHT expected number argument >= 0")
		}
		if n == 0 { // empty string
			return MakeStringResult("")
		}
	default:
		return MakeErrorResult("RIGHT accepts one or two arguments")
	}

	// can't call RIGHT on a range
	if args[0].Type == ResultTypeList {
		return MakeErrorResult("RIGHT can't be called on a range")
	}
	v := args[0].Value()
	m := len(v)
	if n > m {
		return MakeStringResult(v)
	}
	return MakeStringResult(v[m-n : m])
}
