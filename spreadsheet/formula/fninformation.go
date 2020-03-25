// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import (
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/unidoc/unioffice/spreadsheet/reference"
)

func init() {
	initRegexpInformation()
	RegisterFunction("NA", NA)
	RegisterFunction("ISBLANK", IsBlank)
	RegisterFunction("ISERR", IsErr)
	RegisterFunction("ISERROR", IsError)
	RegisterFunction("ISEVEN", IsEven)
	RegisterFunctionComplex("_xlfn.ISFORMULA", IsFormula)
	RegisterFunctionComplex("ORG.OPENOFFICE.ISLEAPYEAR", IsLeapYear)
	RegisterFunctionComplex("ISLOGICAL", IsLogical)
	RegisterFunction("ISNA", IsNA)
	RegisterFunction("ISNONTEXT", IsNonText)
	RegisterFunction("ISNUMBER", IsNumber)
	RegisterFunction("ISODD", IsOdd)
	RegisterFunctionComplex("ISREF", IsRef)
	RegisterFunction("ISTEXT", IsText)
	RegisterFunctionComplex("CELL", Cell)
}

var bs string = string([]byte{92})
var integer, finance, intSep, intPar, decimal, decJust, decPar, par, percent, intCur, cur, curLabel, mdy, dmy, sci *regexp.Regexp

func initRegexpInformation() {
	integer = regexp.MustCompile(`^0+$`)                                                           // 12345
	intSep = regexp.MustCompile("^((#|0)+,)+(#|0)+(;|$)")                                          // 123,456,789
	intPar = regexp.MustCompile("^(#|0|,)*_\\);")                                                  // (123,456,789)
	finance = regexp.MustCompile("^0+\\.(0+)$")                                                    // 1.23
	decimal = regexp.MustCompile("^((#|0)+,)+(#|0)+\\.(0+).*(;|$)")                                // 1.234
	decJust = regexp.MustCompile("^(_|-| )+\\* #+,#+0\\.(0+).*;")                                  // 1.234 with justified horizontal alignment
	decPar = regexp.MustCompile("^((#|0)+,)+(#|0)+\\.((#|0)+)_\\).*;")                             // (1.234)
	percent = regexp.MustCompile("^(#|0)+\\.((#|0)+)%$")                                           // 12.34%
	intCur = regexp.MustCompile("\\[\\$\\$-.+\\](\\* )?(#|0)+,(#|0)+;")                            // $1,234
	cur = regexp.MustCompile("\\[\\$\\$-.+\\](\\* )?(#|0)+,(#|0)+\\.((#|0|-)+).*;")                // $1,234.56
	curLabel = regexp.MustCompile("^((#|0)+,)+(#|0)+(\\.((#|0|-)+))?.+\\[\\$.+\\].*;")             // 1,234.56 USD
	mdy = regexp.MustCompile("^M+(/| |,|\"|" + bs + bs + ")+D+(/| |,|\"|" + bs + bs + ")+Y+$")     // 01/21/2019
	dmy = regexp.MustCompile("^D+(/| |\\.|\"|" + bs + bs + ")+M+(/| |\\.|\"|" + bs + bs + ")+Y+$") // 21. Jan. 2019
	sci = regexp.MustCompile("^(#|0)+\\.((#|0)*)E\\+(#|0)+(;|$)")                                  // 1.02E+002
	par = regexp.MustCompile("^.*_\\).*;")                                                         // (anything in parentheses)
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
		address := "$" + cr.Column + "$" + strconv.Itoa(int(cr.RowIdx))
		if cr.SheetName != "" {
			address = cr.SheetName + "!" + address
		}
		return MakeStringResult(address)
	case "col":
		cr, err := reference.ParseCellReference(refStr)
		if err != nil {
			return MakeErrorResult("Incorrect reference: " + refStr)
		}
		return MakeNumberResult(float64(cr.ColumnIdx + 1))
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
		if result != "G" && strings.Contains(format, "[RED]") {
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
		if cr.SheetName == "" {
			return MakeNumberResult(ctx.GetWidth(int(cr.ColumnIdx)))
		} else {
			return MakeNumberResult(ctx.Sheet(cr.SheetName).GetWidth(int(cr.ColumnIdx)))
		}
	}
	return MakeErrorResult("Incorrect first argument of CELL: " + typ.ValueString)
}

func itemFromEndLength(submatch []string, additionalShift int) string {
	return strconv.Itoa(len(submatch[len(submatch)-1-additionalShift]))
}

// ISBLANK is an implementation of the Excel ISBLANK() function.
func IsBlank(args []Result) Result {
	if len(args) != 1 {
		MakeErrorResult("ISBLANK() accepts a single argument")
	}
	return MakeBoolResult(args[0].Type == ResultTypeEmpty)
}

// ISERR is an implementation of the Excel ISERR() function.
func IsErr(args []Result) Result {
	if len(args) != 1 {
		MakeErrorResult("ISERR() accepts a single argument")
	}

	return MakeBoolResult(args[0].Type == ResultTypeError && args[0].ValueString != "#N/A")
}

// ISERROR is an implementation of the Excel ISERROR() function.
func IsError(args []Result) Result {
	if len(args) != 1 {
		MakeErrorResult("ISERROR() accepts a single argument")
	}

	return MakeBoolResult(args[0].Type == ResultTypeError)
}

// ISEVEN is an implementation of the Excel ISEVEN() function.
func IsEven(args []Result) Result {
	if len(args) != 1 {
		MakeErrorResult("ISEVEN() accepts a single argument")
	}

	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("ISEVEN accepts a numeric argument")
	}
	value := int(args[0].ValueNumber)

	return MakeBoolResult(value == value/2*2)
}

// ISFORMULA is an implementation of the Excel ISFORMULA() function.
func IsFormula(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 1 {
		MakeErrorResult("ISFORMULA() accepts a single argument")
	}
	ref := args[0].Ref
	if ref.Type != ReferenceTypeCell {
		return MakeErrorResult("ISFORMULA requires the first argument to be of type reference")
	}

	return MakeBoolResult(ctx.HasFormula(ref.Value))
}

// IsLeapYear is an implementation of the Excel ISLEAPYEAR() function.
func IsLeapYear(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 1 || args[0].Type != ResultTypeNumber {
		return MakeErrorResult("ISLEAPYEAR requires a single number argument")
	}
	epoch := ctx.GetEpoch()
	t, err := getValueAsTime(args[0].Value(), epoch)
	if err != nil {
		return MakeErrorResult("ISLEAPYEAR requires a single number argument")
	}
	year := t.Year()
	return MakeBoolResult(isLeapYear(year))
}

func getValueAsTime(value string, epoch time.Time) (time.Time, error) {
	f, _, err := big.ParseFloat(value, 10, 128, big.ToNearestEven)
	if err != nil {
		return time.Time{}, err
	}

	day := new(big.Float)
	day.SetUint64(uint64(24 * time.Hour))
	f.Mul(f, day)
	ns, _ := f.Uint64()
	t := epoch.Add(time.Duration(ns))
	return asLocal(t), nil
}

func asLocal(d time.Time) time.Time {
	d = d.UTC()
	return time.Date(d.Year(), d.Month(), d.Day(), d.Hour(),
		d.Minute(), d.Second(), d.Nanosecond(), time.Local)
}

// IsLogical is an implementation of the Excel ISLOGICAL() function.
func IsLogical(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("ISLOGICAL requires a single number argument")
	}
	ref := args[0].Ref
	if ref.Type != ReferenceTypeCell {
		return MakeErrorResult("ISLOGICAL requires the first argument to be of type reference")
	}

	return MakeBoolResult(ctx.Cell(ref.Value, ev).IsBoolean)
}

// IsNA is an implementation of the Excel ISNA() function.
func IsNA(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("ISNA requires one argument")
	}

	return MakeBoolResult(args[0].Type == ResultTypeError && args[0].ValueString == "#N/A")
}

// ISNONTEXT is an implementation of the Excel ISNONTEXT() function.
func IsNonText(args []Result) Result {
	if len(args) != 1 {
		MakeErrorResult("ISNONTEXT() accepts a single argument")
	}
	return MakeBoolResult(args[0].Type != ResultTypeString)
}

// ISNUMBER is an implementation of the Excel ISNUMBER() function.
func IsNumber(args []Result) Result {
	if len(args) != 1 {
		MakeErrorResult("ISNUMBER() accepts a single argument")
	}
	return MakeBoolResult(args[0].Type == ResultTypeNumber)
}

// ISODD is an implementation of the Excel ISODD() function.
func IsOdd(args []Result) Result {
	if len(args) != 1 {
		MakeErrorResult("ISODD() accepts a single argument")
	}

	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("ISODD accepts a numeric argument")
	}
	value := int(args[0].ValueNumber)

	return MakeBoolResult(value != value/2*2)
}

// ISREF is an implementation of the Excel ISREF() function.
func IsRef(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 1 {
		MakeErrorResult("ISREF() accepts a single argument")
	}
	return MakeBoolResult(ev.LastEvalIsRef())
}

// ISTEXT is an implementation of the Excel ISTEXT() function.
func IsText(args []Result) Result {
	if len(args) != 1 {
		MakeErrorResult("ISTEXT() accepts a single argument")
	}
	return MakeBoolResult(args[0].Type == ResultTypeString)
}
