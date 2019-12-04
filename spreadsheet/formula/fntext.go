// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/unidoc/unioffice/internal/wildcard"
)

func init() {
	// RegisterFunction("ASC") Need to figure out how to test
	// RegisterFunction("BAHTTEXT")
	RegisterFunction("CHAR", Char)
	RegisterFunction("CLEAN", Clean)
	RegisterFunction("CODE", Code)
	RegisterFunction("CONCATENATE", Concat)
	RegisterFunction("CONCAT", Concat)
	RegisterFunction("_xlfn.CONCAT", Concat)
	// RegisterFunction("DBCS")
	// RegisterFunction("DOLLAR") Need to test with Excel
	RegisterFunction("EXACT", Exact)
	RegisterFunction("FIND", Find)
	RegisterFunctionComplex("FINDB", Findb)
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
	RegisterFunction("REPLACE", Replace)
	//RegisterFunctionComplex("REPLACEB", Replaceb)
	RegisterFunction("REPT", Rept)
	RegisterFunction("RIGHT", Right)
	RegisterFunction("RIGHTB", Right) // for now
	RegisterFunction("SEARCH", Search)
	RegisterFunctionComplex("SEARCHB", Searchb)
	//RegisterFunction("SUBSTITUTE", )
	RegisterFunction("T", T)
	//RegisterFunction("TEXT")
	RegisterFunction("TEXTJOIN", TextJoin)
	RegisterFunction("_xlfn.TEXTJOIN", TextJoin)
	RegisterFunction("TRIM", Trim)
	RegisterFunction("_xlfn.UNICHAR", Char) // for now
	RegisterFunction("_xlfn.UNICODE", Unicode)
	RegisterFunction("UPPER", Upper)
	RegisterFunction("VALUE", Value)
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

// Concat is an implementation of the Excel CONCAT() and deprecated CONCATENATE() function.
func Concat(args []Result) Result {
	buf := bytes.Buffer{}
	for _, a := range args {
		switch a.Type {
		case ResultTypeString:
			buf.WriteString(a.ValueString)
		case ResultTypeNumber:
			var str string
			if a.IsBoolean {
				if a.ValueNumber == 0 {
					str = "FALSE"
				} else {
					str = "TRUE"
				}
			} else {
				str = a.AsString().ValueString
			}
			buf.WriteString(str)
		default:
			return MakeErrorResult("CONCAT() requires arguments to be strings")
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

type parsedSearchObject struct {
	findText string
	text string
	position int
}

func parseSearchResults(fname string, args []Result) (*parsedSearchObject, Result) {
	if len(args) != 2 && len(args) != 3 {
		return nil, MakeErrorResult(fname + " requires two or three arguments")
	}
	findTextResult := args[0]
	if findTextResult.Type != ResultTypeString {
		return nil, MakeErrorResult("The first argument should be a string")
	}
	textResult := args[1]
	if textResult.Type != ResultTypeString {
		return nil, MakeErrorResult("The second argument should be a string")
	}
	text := textResult.ValueString
	findText := findTextResult.ValueString
	position := 1
	if len(args) == 3 {
		positionResult := args[2]
		if positionResult.Type != ResultTypeNumber {
			return nil, MakeErrorResult("Position should be a number")
		}
		position = int(positionResult.ValueNumber)
		if position < 1 {
			return nil, MakeErrorResultType(ErrorTypeValue, "Position should be a number more than 0")
		}
		if position > len(text) {
			return nil, MakeErrorResultType(ErrorTypeValue, "Position should be a number more than 0")
		}
	}
	return &parsedSearchObject{
		findText,
		text,
		position,
	}, MakeEmptyResult()
}

// Find is an implementation of the Excel FIND().
func Find(args []Result) Result {
	parsed, errResult := parseSearchResults("FIND", args)
	if errResult.Type != ResultTypeEmpty {
		return errResult
	}
	findText := parsed.findText
	if findText == "" {
		return MakeNumberResult(1.0)
	}
	text := parsed.text
	position := parsed.position
	stepsCounter := 1
	for i := range text {
		if stepsCounter < position {
			stepsCounter++
			continue
		}
		index := strings.Index(text[i:], findText)
		if index == 0 {
			return MakeNumberResult(float64(stepsCounter))
		}
		stepsCounter++
	}
	return MakeErrorResultType(ErrorTypeValue, "Not found")
}

// Findb is an implementation of the Excel FINDB().
func Findb(ctx Context, ev Evaluator, args []Result) Result {
	if !ctx.IsDBCS() {
		return Find(args)
	}
	parsed, errResult := parseSearchResults("FIND", args)
	if errResult.Type != ResultTypeEmpty {
		return errResult
	}
	findText := parsed.findText
	if findText == "" {
		return MakeNumberResult(1.0)
	}
	text := parsed.text
	position := parsed.position - 1
	stepsCounter := 1
	lastIndex := 0
	for i := range text {
		if i != 0 {
			add := 1
			if i - lastIndex > 1 {
				add = 2
			}
			stepsCounter += add
		}
		if stepsCounter > position {
			index := strings.Index(text[i:], findText)
			if index == 0 {
				return MakeNumberResult(float64(stepsCounter))
			}
		}
		lastIndex = i
	}
	return MakeErrorResultType(ErrorTypeValue, "Not found")
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

	arg := args[0]
	switch arg.Type {
	case ResultTypeError:
		return arg
	case ResultTypeNumber, ResultTypeString:
		return lower(args[0])
	case ResultTypeList:
		list := arg.ValueList
		resultList := []Result{}
		for _, v := range list {
			vLower := lower(v)
			if vLower.Type == ResultTypeError {
				return vLower
			}
			resultList = append(resultList, vLower)
		}
		return MakeListResult(resultList)
	case ResultTypeArray:
		array := arg.ValueArray
		resultArray := [][]Result{}
		for _, r := range array {
			row := []Result{}
			for _, v := range r {
				vLower := lower(v)
				if vLower.Type == ResultTypeError {
					return vLower
				}
				row = append(row, vLower)
			}
			resultArray = append(resultArray, row)
		}
		return MakeArrayResult(resultArray)

	default:
		return MakeErrorResult("Incorrect argument for LOWER")
	}
	return MakeErrorResult("Incorrect argument for LOWER")
}

func lower(arg Result) Result {
	if arg.Type == ResultTypeEmpty {
		return arg
	}
	s := arg.AsString()
	if s.Type != ResultTypeString {
		return MakeErrorResult("LOWER requires a single string argument")
	}
	if arg.IsBoolean {
		if s.ValueString == "1" {
			return MakeStringResult("true")
		} else if s.ValueString == "0" {
			return MakeStringResult("false")
		} else {
			return MakeErrorResult("Incorrect argument for LOWER")
		}
	} else {
		return MakeStringResult(strings.ToLower(s.ValueString))
	}
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

// Search is an implementation of the Excel SEARCH().
func Search(args []Result) Result {
	parsed, errResult := parseSearchResults("FIND", args)
	if errResult.Type != ResultTypeEmpty {
		return errResult
	}
	findText := strings.ToLower(parsed.findText)
	if findText == "" {
		return MakeNumberResult(1.0)
	}
	text := strings.ToLower(parsed.text)
	position := parsed.position
	stepsCounter := 1
	for i := range text {
		if stepsCounter < position {
			stepsCounter++
			continue
		}
		index := wildcard.Index(findText, text[i:])
		if index == 0 {
			return MakeNumberResult(float64(stepsCounter))
		}
		stepsCounter++
	}
	return MakeErrorResultType(ErrorTypeValue, "Not found")
}

// Searchb is an implementation of the Excel SEARCHB().
func Searchb(ctx Context, ev Evaluator, args []Result) Result {
	if !ctx.IsDBCS() {
		return Search(args)
	}
	parsed, errResult := parseSearchResults("FIND", args)
	if errResult.Type != ResultTypeEmpty {
		return errResult
	}
	findText := strings.ToLower(parsed.findText)
	text := strings.ToLower(parsed.text)
	if findText == "" {
		return MakeNumberResult(1.0)
	}
	position := parsed.position - 1
	stepsCounter := 1
	lastIndex := 0
	for i := range text {
		if i != 0 {
			add := 1
			if i - lastIndex > 1 {
				add = 2
			}
			stepsCounter += add
		}
		if stepsCounter > position {
			index := wildcard.Index(findText, text[i:])
			if index == 0 {
				return MakeNumberResult(float64(stepsCounter))
			}
		}
		lastIndex = i
	}
	return MakeErrorResultType(ErrorTypeValue, "Not found")
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

// Value is an implementation of the Excel VALUE function.
func Value(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("VALUE requires a single argument")
	}

	value := args[0]
	if value.Type == ResultTypeNumber {
		return value
	}

	if value.Type == ResultTypeString {
		result, err := strconv.ParseFloat(value.Value(), 64)
		if err == nil {
			return MakeNumberResult(result)
		}
	}

	return MakeErrorResult("Incorrect argument for VALUE")
}

type parsedReplaceObject struct {
	text string
	startPos int
	length int
	textToReplace string
}

func parseReplaceResults(fname string, args []Result) (*parsedReplaceObject, Result) {
	if len(args) != 4 {
		return nil, MakeErrorResult(fname + " requires four arguments")
	}
	if args[0].Type != ResultTypeString {
		return nil, MakeErrorResult(fname + " requires first argument to be a string")
	}
	text := args[0].ValueString
	if args[1].Type != ResultTypeNumber {
		return nil, MakeErrorResult(fname + " requires second argument to be a number")
	}
	startPos := int(args[1].ValueNumber) - 1
	if args[2].Type != ResultTypeNumber {
		return nil, MakeErrorResult(fname + " requires third argument to be a number")
	}
	length := int(args[2].ValueNumber)
	if args[3].Type != ResultTypeString {
		return nil, MakeErrorResult(fname + " requires fourth argument to be a string")
	}
	textToReplace := args[3].ValueString
	return &parsedReplaceObject{
		text,
		startPos,
		length,
		textToReplace,
	}, MakeEmptyResult()
}

// Replace is an implementation of the Excel REPLACE().
func Replace(args []Result) Result {
	parsed, errResult := parseReplaceResults("REPLACE", args)
	if errResult.Type != ResultTypeEmpty {
		return errResult
	}
	text := parsed.text
	startPos := parsed.startPos
	length := parsed.length
	textToReplace := parsed.textToReplace
	textLen := len(text)
	if startPos > textLen {
		startPos = textLen
	}
	endPos := startPos + length
	if endPos > textLen {
		endPos = textLen
	}
	newText := text[0:startPos] + textToReplace + text[endPos:]
	return MakeStringResult(newText)
}

// TextJoin is an implementation of the Excel TEXTJOIN function.
func TextJoin(args []Result) Result {
	if len(args) < 3 {
		return MakeErrorResult("TEXTJOIN requires three or more arguments")
	}

	if args[0].Type != ResultTypeString {
		return MakeErrorResult("TEXTJOIN requires delimiter to be a string")
	}
	delimiter := args[0].ValueString

	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("TEXTJOIN requires second argument to be a number or boolean")
	}
	ignoreEmpty := args[1].ValueNumber != 0

	arr := collectStrings(args[2:], []string{}, ignoreEmpty)
	return MakeStringResult(strings.Join(arr, delimiter))
}

func collectStrings(args []Result, arr []string, ignoreEmpty bool) []string {
	for _, result := range args {
		switch result.Type {
		case ResultTypeEmpty:
			if !ignoreEmpty {
				arr = append(arr, "")
			}
		case ResultTypeString:
			if result.ValueString != "" || !ignoreEmpty {
				arr = append(arr, result.ValueString)
			}
		case ResultTypeNumber:
			arr = append(arr, result.Value())
		case ResultTypeList:
			arr = appendSlices(arr, collectStrings(result.ValueList, []string{}, ignoreEmpty))
		case ResultTypeArray:
			for _, row := range result.ValueArray {
				arr = appendSlices(arr, collectStrings(row, []string{}, ignoreEmpty))
			}
		}
	}
	return arr
}

func appendSlices(s0, s1 []string) []string {
	for _, item := range s1 {
		s0 = append(s0, item)
	}
	return s0
}
