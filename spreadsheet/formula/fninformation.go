// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/unidoc/unioffice/spreadsheet/reference"
)

func init() {
	initRegexp()
	RegisterFunction("NA", NA)
	RegisterFunctionComplex("CELL", Cell)
}

var bs string = string([]byte{92})
var integer, finance, intSep, intPar, decimal, decJust, decPar, par, percent, intCur, cur, curLabel, mdy, dmy, sci *regexp.Regexp

func initRegexp() {
	integer = regexp.MustCompile(`^0+$`) // 12345
	intSep = regexp.MustCompile("^((#|0)+,)+(#|0)+(;|$)") // 123,456,789
	intPar = regexp.MustCompile("^(#|0|,)*_\\);") // (123,456,789)
	finance = regexp.MustCompile("^0+\\.(0+)$") // 1.23
	decimal = regexp.MustCompile("^((#|0)+,)+(#|0)+\\.(0+).*(;|$)") // 1.234
	decJust = regexp.MustCompile("^(_|-| )+\\* #+,#+0\\.(0+).*;") // 1.234 with justified horizontal alignment
	decPar = regexp.MustCompile("^((#|0)+,)+(#|0)+\\.((#|0)+)_\\).*;")  // (1.234)
	percent = regexp.MustCompile("^(#|0)+\\.((#|0)+)%$") // 12.34%
	intCur = regexp.MustCompile("\\[\\$\\$-.+\\](\\* )?(#|0)+,(#|0)+;") // $1,234
	cur = regexp.MustCompile("\\[\\$\\$-.+\\](\\* )?(#|0)+,(#|0)+\\.((#|0|-)+).*;") // $1,234.56
	curLabel = regexp.MustCompile("^((#|0)+,)+(#|0)+(\\.((#|0|-)+))?.+\\[\\$.+\\].*;") // 1,234.56 USD
	mdy = regexp.MustCompile("^M+(/| |,|\"|" + bs + bs + ")+D+(/| |,|\"|" + bs + bs + ")+Y+$") // 01/21/2019
	dmy = regexp.MustCompile("^D+(/| |\\.|\"|" + bs + bs + ")+M+(/| |\\.|\"|" + bs + bs + ")+Y+$") // 21. Jan. 2019
	sci = regexp.MustCompile("^(#|0)+\\.((#|0)*)E\\+(#|0)+(;|$)") // 1.02E+002
	par = regexp.MustCompile("^.*_\\).*;") // (anything in parentheses)
}

// NA is an implementation of the Excel NA() function that just returns the #N/A! error.
func NA(args []Result) Result {
	if len(args) != 0 {
		MakeErrorResult("NA() accepts no arguments")
	}
	return MakeErrorResultType(ErrorTypeNA, "")
}

// Cell is an implementation of the Excel CELL function that returns information
// about the formatting, location, or contents of a cell.
func Cell(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 1 && len(args) != 2 {
		return MakeErrorResult("CELL requires one or two arguments")
	}
	typ := args[0].AsString()
	if typ.Type != ResultTypeString {
		return MakeErrorResult("CELL requires first argument to be of type string")
	}
	ref := args[1].Ref
	if ref.Type != ReferenceTypeCell {
		return MakeErrorResult("CELL requires second argument to be of type reference")
	}
	refStr := ref.Value

	switch typ.ValueString {
	case "address":
		cr, err := reference.ParseCellReference(refStr)
		if err != nil {
			return MakeErrorResult("Incorrect reference: " + refStr)
		}
		address := "$"+cr.Column+"$"+strconv.Itoa(int(cr.RowIdx))
		return MakeStringResult(address)
	case "col":
		cr, err := reference.ParseCellReference(refStr)
		if err != nil {
			return MakeErrorResult("Incorrect reference: " + refStr)
		}
		return MakeNumberResult(float64(cr.ColumnIdx+1))
	case "color":
		red := strings.Contains(ctx.GetFormat(refStr), "[RED]")
		return MakeBoolResult(red)
	case "contents":
		return args[1]
	case "filename":
		return MakeStringResult(ctx.GetFilename())
	case "format":
		result := "G"
		format := ctx.GetFormat(refStr)
		if format == "General" || integer.MatchString(format) {
			result = "F0"
		} else if format == "0%" {
			result = "P0"
		} else if format == "MMM DD" {
			result = "D2"
		} else if format == "MM/YY" {
			result = "D3"
		} else if format == "MM/DD/YY\\ HH:MM\\ AM/PM" || format == "MM/DD/YYYY\\ HH:MM:SS" {
			result = "D4"
		} else if format == "MM\\-DD" {
			result = "D5"
		} else if format == "HH:MM:SS\\ AM/PM" {
			result = "D6"
		} else if format == "HH:MM\\ AM/PM" {
			result = "D7"
		} else if format == "HH:MM:SS" {
			result = "D8"
		} else if format == "HH:MM" {
			result = "D9"
		} else if intSep.MatchString(format) {
			result = ".0"
		} else if intPar.MatchString(format) {
			result = ".0()"
		} else if intCur.MatchString(format) {
			result = "C0"
		} else if mdy.MatchString(format) || dmy.MatchString(format) {
			result = "D1"
		} else if submatch := finance.FindStringSubmatch(format); len(submatch) > 1 {
			result = "F" + strconv.Itoa(len(submatch[1]))
		} else if submatch := decJust.FindStringSubmatch(format); len(submatch) > 1 {
			result = "." + strconv.Itoa(len(submatch[2]))
		} else if submatch := percent.FindStringSubmatch(format); len(submatch) > 1 {
			result = "P" + strconv.Itoa(len(submatch[2]))
		} else if submatch := cur.FindStringSubmatch(format); len(submatch) > 1 {
			result = "C" + itemFromEndLength(submatch, 1)
		} else if submatch := curLabel.FindStringSubmatch(format); len(submatch) > 1 {
			result = "C" + itemFromEndLength(submatch, 1)
		} else if submatch := decPar.FindStringSubmatch(format); len(submatch) > 1 {
			result = "." + itemFromEndLength(submatch, 1) + "()"
		} else if submatch := decimal.FindStringSubmatch(format); len(submatch) > 1 {
			result = "." + itemFromEndLength(submatch, 1)
		} else if submatch := sci.FindStringSubmatch(format); len(submatch) > 1 {
			result = "S" + itemFromEndLength(submatch, 3)
		}
		if strings.Contains(format, "[RED]") {
			result += "-"
		}
		return MakeStringResult(result)
	case "parentheses":
		format := ctx.GetFormat(refStr)
		if par.MatchString(format) {
			return MakeNumberResult(1)
		} else {
			return MakeNumberResult(0)
		}
	case "prefix":
		return MakeStringResult(ctx.GetLabelPrefix(refStr))
	case "protect":
		result := 0.0
		if ctx.GetLocked(refStr) {
			result = 1.0
		}
		return MakeNumberResult(result)
	case "row":
		cr, err := reference.ParseCellReference(refStr)
		if err != nil {
			return MakeErrorResult("Incorrect reference: " + refStr)
		}
		return MakeNumberResult(float64(cr.RowIdx))
	case "type":
		switch args[1].Type {
		case ResultTypeEmpty:
			return MakeStringResult("b")
		case ResultTypeString:
			return MakeStringResult("l")
		default:
			return MakeStringResult("v")
		}
	case "width":
		cr, err := reference.ParseCellReference(refStr)
		if err != nil {
			return MakeErrorResult("Incorrect reference: " + refStr)
		}
		return MakeNumberResult(ctx.GetWidth(int(cr.ColumnIdx)))
	}
	return MakeErrorResult("Incorrect first argument of CELL: "+typ.ValueString)
}

func itemFromEndLength(submatch []string, additionalShift int) string {
	return strconv.Itoa(len(submatch[len(submatch)-1-additionalShift]))
}
