// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

import (
	"math"
	"sort"

	"baliance.com/gooxml"
)

func init() {
	RegisterFunction("AVERAGE", Average)
	RegisterFunction("AVERAGEA", Averagea)
	RegisterFunction("COUNT", Count)
	RegisterFunction("COUNTA", Counta)
	RegisterFunction("COUNTBLANK", CountBlank)
	RegisterFunction("MAX", Max)
	RegisterFunction("MIN", Min)
	RegisterFunction("MEDIAN", Median)
}

func sumCount(args []Result, countText bool) (float64, float64) {
	cnt := 0.0
	sum := 0.0
	for _, arg := range args {
		switch arg.Type {
		case ResultTypeNumber:
			sum += arg.ValueNumber
			cnt++
		case ResultTypeList, ResultTypeArray:
			s, c := sumCount(arg.ListValues(), countText)
			sum += s
			cnt += c
		case ResultTypeString:
			if countText {
				cnt++
			}
		case ResultTypeEmpty: // do nothing
		}
	}

	return sum, cnt
}

// Average implements the AVERAGE function. It differs slightly from Excel (and
// agrees with LibreOffice) in that boolean values are counted. As an example,
// AVERAGE of two cells containing TRUE & FALSE is 0.5 in LibreOffice and
// #DIV/0! in Excel. gooxml will return 0.5 in this case.
func Average(args []Result) Result {
	sum, cnt := sumCount(args, false)
	if cnt == 0 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "AVERAGE divide by zero")
	}
	return MakeNumberResult(sum / cnt)
}

// Averagea implements the AVERAGEA function, AVERAGEA counts cells that contain
// text as a zero where AVERAGE ignores them entirely.
func Averagea(args []Result) Result {
	sum, cnt := sumCount(args, true)
	if cnt == 0 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "AVERAGE divide by zero")
	}
	return MakeNumberResult(sum / cnt)
}

type countMode byte

const (
	countNormal countMode = iota
	countText
	countEmpty
)

func count(args []Result, m countMode) float64 {
	cnt := 0.0
	for _, arg := range args {
		switch arg.Type {
		case ResultTypeNumber:
			if m != countEmpty {
				cnt++
			}
		case ResultTypeList, ResultTypeArray:
			cnt += count(arg.ListValues(), m)
		case ResultTypeString:
			if m == countText {
				cnt++
			}
		case ResultTypeEmpty:
			if m == countEmpty {
				cnt++
			}
		}
	}

	return cnt
}

// Count implements the COUNT function.
func Count(args []Result) Result {
	return MakeNumberResult(count(args, countNormal))
}

// Counta implements the COUNTA function.
func Counta(args []Result) Result {
	return MakeNumberResult(count(args, countText))
}

// CountBlank implements the COUNTBLANK function.
func CountBlank(args []Result) Result {
	// COUNT and COUNTA don't require arguments, COUNTBLANK does
	if len(args) == 0 {
		return MakeErrorResult("COUNTBLANK requires an argument")
	}
	return MakeNumberResult(count(args, countEmpty))
}

// Min is an implementation of the Excel MIN() function.
func Min(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("MIN requires at least one argument")
	}

	v := math.MaxFloat64
	for _, a := range args {
		a = a.AsNumber()
		switch a.Type {
		case ResultTypeNumber:
			if a.ValueNumber < v {
				v = a.ValueNumber
			}
		case ResultTypeList, ResultTypeArray:
			subMin := Min(a.ListValues())
			if subMin.ValueNumber < v {
				v = subMin.ValueNumber
			}
		case ResultTypeString:
			// treated as zero by Excel
			if 0 < v {
				v = 0
			}
		case ResultTypeEmpty:
		// skip
		case ResultTypeError:
			return a
		default:
			gooxml.Log("unhandled MIN() argument type %s", a.Type)
		}
	}
	if v == math.MaxFloat64 {
		v = 0
	}
	return MakeNumberResult(v)
}

// Max is an implementation of the Excel MAX() function.
func Max(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("MAX requires at least one argument")
	}
	v := -math.MaxFloat64
	for _, a := range args {
		a = a.AsNumber()
		switch a.Type {
		case ResultTypeNumber:
			if a.ValueNumber > v {
				v = a.ValueNumber
			}
		case ResultTypeList, ResultTypeArray:
			subMax := Max(a.ListValues())
			if subMax.ValueNumber > v {
				v = subMax.ValueNumber
			}
		case ResultTypeEmpty:
			// skip
		case ResultTypeString:
			// treated as zero by Excel
			if 0 > v {
				v = 0
			}
		default:
			gooxml.Log("unhandled MAX() argument type %s", a.Type)
		}
	}
	if v == -math.MaxFloat64 {
		v = 0
	}
	return MakeNumberResult(v)
}

func extractNumbers(args []Result) []float64 {
	values := make([]float64, 0)
	for _, a := range args {
		if a.Type == ResultTypeEmpty {
			continue
		}
		a = a.AsNumber()
		switch a.Type {
		case ResultTypeNumber:
			values = append(values, a.ValueNumber)
		case ResultTypeList, ResultTypeArray:
			values = append(values, extractNumbers(a.ListValues())...)
		case ResultTypeString:
			// treated as zero by Excel
		default:
			gooxml.Log("unhandled extractNumbers argument type %s", a.Type)
		}
	}
	return values
}

// Median implements the MEDIAN function that returns the median of a range of
// values.
func Median(args []Result) Result {
	if len(args) == 0 {
		return MakeErrorResult("MEDIAN requires at least one argument")
	}
	values := extractNumbers(args)
	sort.Float64s(values)
	var v float64
	if len(values)%2 == 0 {
		v = (values[len(values)/2-1] + values[len(values)/2]) / 2
	} else {
		v = values[len(values)/2]
	}
	return MakeNumberResult(v)
}
