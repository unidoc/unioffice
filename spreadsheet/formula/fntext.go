// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func init() {
	// RegisterFunction("ASC") Need to figure out how to test
	// RegisterFunction("BAHTTEXT")
	RegisterFunction("CHAR", Char)
	RegisterFunction("CLEAN", Clean)
	RegisterFunction("CODE", Code)
	RegisterFunction("CONCATENATE", Concatenate)
	// RegisterFunction("CONCAT", ) Need to test with Excel
	// RegisterFunction("DBCS")
	// RegisterFunction("DOLLAR") Need to test with Excel
	RegisterFunction("EXACT", Exact)
	// RegisterFunction("FIND")
	// RegisterFunction("FINDB")
	RegisterFunction("LEFT", Left)
	RegisterFunction("LEFTB", Left) // for now
	RegisterFunction("LEN", Len)
	RegisterFunction("LENB", Len) // for now
	RegisterFunction("LOWER", Lower)
	// RegisterFunction("MID")
	// RegisterFunction("MIDB")
	// RegisterFunction("NUMBERVALUE")
	// RegisterFunction("PHONETIC")
	RegisterFunction("PROPER", Proper)
	// RegisterFunction("REPLACE")
	// RegisterFunction("REPLACEB")
	RegisterFunction("REPT", Rept)
	RegisterFunction("RIGHT", Right)
	RegisterFunction("RIGHTB", Right) // for now
	//RegisterFunction("SEARCH", )
	//RegisterFunction("SEARCHB", )
	//RegisterFunction("SUBSTITUTE", )
	RegisterFunction("T", T)
	//RegisterFunction("TEXT")
	//RegisterFunction("TEXTJOIN")
	RegisterFunction("TRIM", Trim)
	RegisterFunction("_xlfn.UNICHAR", Char) // for now
	RegisterFunction("_xlfn.UNICODE", Unicode)
	RegisterFunction("UPPER", Upper)
	//RegisterFunction("VALUE", )
}

// Char is an implementation of the Excel CHAR function that takes an integer in
// the range [0,255] and returns the corresponding ASCII character.
func Char(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("CHAR requires a single numeric argument")
	}
	c := args[0].AsNumber()
	if c.Type != ResultTypeNumber {
		return MakeErrorResult("CHAR requires a single numeric argument")
	}
	cv := int(c.ValueNumber)
	if cv < 0 || cv > 255 {
		return MakeErrorResult("CHAR requires arguments in the range [0,255]")
	}
	return MakeStringResult(fmt.Sprintf("%c", cv))
}

// Clean is an implementation of the Excel CLEAN function that removes
// unprintable characters.
func Clean(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("CLEAN requires a single string argument")
	}
	s := args[0].AsString()
	if s.Type != ResultTypeString {
		return MakeErrorResult("CHAR requires a single string argument")
	}
	b := bytes.Buffer{}
	for _, c := range s.ValueString {
		if unicode.IsPrint(c) {
			b.WriteRune(c)
		}
	}
	return MakeStringResult(b.String())
}

// Code is an implementation of the Excel CODE function that returns the first
// character of the string as a number.
func Code(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("CODE requires a single string argument")
	}
	s := args[0].AsString()
	if s.Type != ResultTypeString {
		return MakeErrorResult("CODE requires a single string argument")
	}
	// Zero length string returns a zero
	if len(s.ValueString) == 0 {
		return MakeNumberResult(0)
	}

	return MakeNumberResult(float64(s.ValueString[0]))
}

func Unicode(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("UNICODE requires a single string argument")
	}
	s := args[0].AsString()
	if s.Type != ResultTypeString {
		return MakeErrorResult("UNICODE requires a single string argument")
	}
	// Zero length string returns an error
	if len(s.ValueString) == 0 {
		return MakeErrorResult("UNICODE requires a non-zero length argument")
	}

	return MakeNumberResult(float64(s.ValueString[0]))
}

// Concatenate is an implementation of the Excel CONCATENATE() function.
func Concatenate(args []Result) Result {
	buf := bytes.Buffer{}
	for _, a := range args {
		a = a.AsString()
		switch a.Type {
		case ResultTypeString:
			buf.WriteString(a.ValueString)
		default:
			return MakeErrorResult("CONCATENATE() requires arguments to be strings")
		}
	}
	return MakeStringResult(buf.String())
}

// Exact is an implementation of the Excel EXACT() which compares two strings.
func Exact(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("CONCATENATE() requires two string arguments")
	}
	arg1 := args[0].AsString()
	arg2 := args[1].AsString()
	if arg1.Type != ResultTypeString || arg2.Type != ResultTypeString {
		return MakeErrorResult("CONCATENATE() requires two string arguments")
	}
	return MakeBoolResult(arg1.ValueString == arg2.ValueString)
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

// Len is an implementation of the Excel LEN function that returns length of a string
func Len(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("LEN requires a single string argument")
	}
	s := args[0].AsString()
	if s.Type != ResultTypeString {
		return MakeErrorResult("LEN requires a single string argument")
	}

	return MakeNumberResult(float64(len(s.ValueString)))
}

// Lower is an implementation of the Excel LOWER function that returns a lower
// case version of a string.
func Lower(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("LOWER requires a single string argument")
	}
	s := args[0].AsString()
	if s.Type != ResultTypeString {
		return MakeErrorResult("LOWER requires a single string argument")
	}

	return MakeStringResult(strings.ToLower(s.ValueString))
}

// Proper is an implementation of the Excel PROPER function that returns a copy
// of the string with each word capitalized.
func Proper(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("PROPER requires a single string argument")
	}
	s := args[0].AsString()
	if s.Type != ResultTypeString {
		return MakeErrorResult("PROPER requires a single string argument")
	}

	buf := bytes.Buffer{}
	prevWasLetter := false
	for _, c := range s.ValueString {
		if !prevWasLetter && unicode.IsLetter(c) {
			buf.WriteRune(unicode.ToUpper(c))
		} else {
			// seems odd but matches Excel's behavior
			buf.WriteRune(unicode.ToLower(c))
		}
		prevWasLetter = unicode.IsLetter(c)
	}

	return MakeStringResult(buf.String())
}

// Rept is an implementation of the Excel REPT function that returns n copies of
// a string.
func Rept(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("REPT requires two arguments")
	}
	s := args[0].AsString()
	if s.Type != ResultTypeString {
		return MakeErrorResult("PROPER requires first argument to be a string")
	}

	n := args[1].AsNumber()
	if n.Type != ResultTypeNumber {
		return MakeErrorResult("PROPER requires second argument to be a number")
	}
	if n.ValueNumber < 0 {
		return MakeErrorResult("PROPER requires second argument to be >= 0")
	}
	if n.ValueNumber == 0 {
		return MakeStringResult("")
	}

	buf := bytes.Buffer{}
	for i := 0; i < int(n.ValueNumber); i++ {
		buf.WriteString(s.ValueString)
	}
	return MakeStringResult(buf.String())
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

// T is an implementation of the Excel T function that returns whether the
// argument is text.
func T(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("T requires a single string argument")
	}
	s := args[0]
	if s.Type == ResultTypeError || s.Type == ResultTypeString {
		return s
	}
	return MakeEmptyResult()
}

// Trim is an implementation of the Excel TRIM function that removes leading,
// trailing and consecutive spaces.
func Trim(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("TRIM requires a single string argument")
	}
	s := args[0].AsString()
	if s.Type != ResultTypeString {
		return MakeErrorResult("TRIM requires a single string argument")
	}
	buf := bytes.Buffer{}
	seenLetter := false
	prevWasSpace := false
	trailingSpaces := 0
	for _, c := range s.ValueString {
		isSpace := c == ' '
		if isSpace {
			if !seenLetter {
				continue
			}
			if !prevWasSpace {
				trailingSpaces++
				buf.WriteRune(c)
			}
		} else {
			trailingSpaces = 0
			seenLetter = true
			buf.WriteRune(c)
		}
		prevWasSpace = isSpace
	}
	buf.Truncate(buf.Len() - trailingSpaces)
	return MakeStringResult(buf.String())
}

// Upper is an implementation of the Excel UPPER function that returns a upper
// case version of a string.
func Upper(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("UPPER requires a single string argument")
	}
	s := args[0].AsString()
	if s.Type != ResultTypeString {
		return MakeErrorResult("UPPER requires a single string argument")
	}

	return MakeStringResult(strings.ToUpper(s.ValueString))
}
