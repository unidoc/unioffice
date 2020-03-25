// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import (
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/internal/wildcard"
)

func init() {
	initRegexpStatistical()
	RegisterFunction("AVERAGE", Average)
	RegisterFunction("AVERAGEA", Averagea)
	RegisterFunction("COUNT", Count)
	RegisterFunction("COUNTA", Counta)
	RegisterFunction("COUNTIF", CountIf)
	RegisterFunction("COUNTIFS", CountIfs)
	RegisterFunction("COUNTBLANK", CountBlank)
	RegisterFunction("MAX", Max)
	RegisterFunction("MAXA", MaxA)
	RegisterFunction("MAXIFS", MaxIfs)
	RegisterFunction("_xlfn.MAXIFS", MaxIfs)
	RegisterFunction("MEDIAN", Median)
	RegisterFunction("MIN", Min)
	RegisterFunction("MINA", MinA)
	RegisterFunction("MINIFS", MinIfs)
	RegisterFunction("_xlfn.MINIFS", MinIfs)
}

var number, eq, g, l, ge, le *regexp.Regexp

func initRegexpStatistical() {
	number = regexp.MustCompile(`^([0-9]+)$`)
	eq = regexp.MustCompile(`^=(.*)$`)  // =-12345.67, =A6
	l = regexp.MustCompile(`^<(.*)$`)   // <-12345.67, <A6
	g = regexp.MustCompile(`^>(.*)$`)   // >-12345.67, >A6
	le = regexp.MustCompile(`^<=(.*)$`) // <=-12345.67, <=A6
	ge = regexp.MustCompile(`^>=(.*)$`) // >=-12345.67, >=A6
}

func sumCount(args []Result, countText bool) (float64, float64) {
	cnt := 0.0
	sum := 0.0
	for _, arg := range args {
		switch arg.Type {
		case ResultTypeNumber:
			if countText || !arg.IsBoolean {
				sum += arg.ValueNumber
				cnt++
			}
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
			if m == countText || (m == countNormal && !arg.IsBoolean) {
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

type criteriaParsed struct {
	isNumber bool
	cNum     float64
	cStr     string
	cRegex   *criteriaRegex
}

const (
	_ byte = iota
	isEq
	isLe
	isGe
	isL
	isG
)

type criteriaRegex struct {
	regexType   byte   // type of condition
	compareWith string // value to apply condition to
}

func parseCriteria(criteria Result) *criteriaParsed {
	isNumber := criteria.Type == ResultTypeNumber
	cNum := criteria.ValueNumber
	cStr := strings.ToLower(criteria.ValueString)
	cRegex := parseCriteriaRegex(cStr)
	return &criteriaParsed{
		isNumber,
		cNum,
		cStr,
		cRegex,
	}
}

func parseCriteriaRegex(cStr string) *criteriaRegex {
	cRegex := &criteriaRegex{}
	if cStr == "" {
		return cRegex
	}
	if submatch := number.FindStringSubmatch(cStr); len(submatch) > 1 {
		cRegex.regexType = isEq
		cRegex.compareWith = submatch[1]
	} else if submatch := eq.FindStringSubmatch(cStr); len(submatch) > 1 {
		cRegex.regexType = isEq
		cRegex.compareWith = submatch[1]
	} else if submatch := le.FindStringSubmatch(cStr); len(submatch) > 1 {
		cRegex.regexType = isLe
		cRegex.compareWith = submatch[1]
	} else if submatch := ge.FindStringSubmatch(cStr); len(submatch) > 1 {
		cRegex.regexType = isGe
		cRegex.compareWith = submatch[1]
	} else if submatch := l.FindStringSubmatch(cStr); len(submatch) > 1 {
		cRegex.regexType = isL
		cRegex.compareWith = submatch[1]
	} else if submatch := g.FindStringSubmatch(cStr); len(submatch) > 1 {
		cRegex.regexType = isG
		cRegex.compareWith = submatch[1]
	}
	return cRegex
}

// CountIf implements the COUNTIF function.
func CountIf(args []Result) Result {
	if len(args) < 2 {
		return MakeErrorResult("COUNTIF requires two argumentss")
	}
	arr := args[0]
	if arr.Type != ResultTypeArray && arr.Type != ResultTypeList {
		return MakeErrorResult("COUNTIF requires first argument of type array")
	}
	criteria := parseCriteria(args[1])
	count := 0
	for _, r := range arrayFromRange(arr) {
		for _, value := range r {
			if compare(value, criteria) {
				count++
			}
		}
	}
	return MakeNumberResult(float64(count))
}

func arrayFromRange(result Result) [][]Result {
	switch result.Type {
	case ResultTypeArray:
		return result.ValueArray
	case ResultTypeList:
		return [][]Result{
			result.ValueList,
		}
	default:
		return [][]Result{}
	}
}

// helper type for storing indexes of found values
type rangeIndex struct {
	rowIndex int
	colIndex int
}

func checkIfsRanges(args []Result, sumRange bool, fnName string) Result {
	// quick check before searching
	var minArgs, oddEven string
	if sumRange {
		minArgs = "three"
		oddEven = "odd"
	} else {
		minArgs = "two"
		oddEven = "even"
	}
	argsNum := len(args)
	if (sumRange && argsNum < 3) || (!sumRange && argsNum < 2) {
		return MakeErrorResult(fnName + " requires at least " + minArgs + " argumentss")
	}
	if (argsNum/2*2 == argsNum) == sumRange {
		return MakeErrorResult(fnName + " requires " + oddEven + " number of arguments")
	}

	rangeWidth := -1
	rangeHeight := -1
	for i := 0; i < argsNum; i += 2 {
		arrResult := args[i]
		if arrResult.Type != ResultTypeArray && arrResult.Type != ResultTypeList {
			return MakeErrorResult(fnName + " requires ranges of type list or array")
		}
		arr := arrayFromRange(arrResult)
		if rangeHeight == -1 {
			rangeHeight = len(arr)
			rangeWidth = len(arr[0])
		} else if len(arr) != rangeHeight || len(arr[0]) != rangeWidth {
			return MakeErrorResult(fnName + " requires all ranges to be of the same size")
		}
		if sumRange && i == 0 {
			i-- // after sumRange should go column 1, not 2
		}
	}
	return empty
}

//getIfsMatch returns an array of indexes of cells which meets all *IFS criterias
func getIfsMatch(args []Result) []rangeIndex {
	toLook := []rangeIndex{}
	argsNum := len(args)
	for i := 0; i < argsNum-1; i += 2 {
		found := []rangeIndex{}
		arr := arrayFromRange(args[i])
		criteria := parseCriteria(args[i+1])
		if i == 0 {
			for rowIndex, row := range arr { // for the first range look in every cell of the range
				for colIndex, value := range row {
					if compare(value, criteria) {
						found = append(found, rangeIndex{rowIndex, colIndex})
					}
				}
			}
		} else {
			for _, index2d := range toLook { // for next ranges look only in cells of the range in which values matched for the previous range
				value := arr[index2d.rowIndex][index2d.colIndex]
				if compare(value, criteria) {
					found = append(found, index2d)
				}
			}
		}
		if len(found) == 0 { // if nothing found at some step no sense to continue
			return []rangeIndex{}
		}
		toLook = found[:] // next time look only in the cells with the same indexes where matches happen in the previous range
	}
	return toLook
}

// CountIfs implements the COUNTIFS function.
func CountIfs(args []Result) Result {
	errorResult := checkIfsRanges(args, false, "COUNTIFS")
	if errorResult.Type != ResultTypeEmpty {
		return errorResult
	}
	match := getIfsMatch(args)
	return MakeNumberResult(float64(len(match)))
}

// MaxIfs implements the MAXIFS function.
func MaxIfs(args []Result) Result {
	errorResult := checkIfsRanges(args, true, "MAXIFS")
	if errorResult.Type != ResultTypeEmpty {
		return errorResult
	}
	match := getIfsMatch(args[1:])
	max := -math.MaxFloat64
	maxArr := arrayFromRange(args[0])
	for _, indexes := range match {
		value := maxArr[indexes.rowIndex][indexes.colIndex].ValueNumber
		if max < value {
			max = value
		}
	}
	if max == -math.MaxFloat64 {
		max = 0
	}
	return MakeNumberResult(float64(max))
}

// MinIfs implements the MINIFS function.
func MinIfs(args []Result) Result {
	errorResult := checkIfsRanges(args, true, "MINIFS")
	if errorResult.Type != ResultTypeEmpty {
		return errorResult
	}
	match := getIfsMatch(args[1:])
	min := math.MaxFloat64
	minArr := arrayFromRange(args[0])
	for _, indexes := range match {
		value := minArr[indexes.rowIndex][indexes.colIndex].ValueNumber
		if min > value {
			min = value
		}
	}
	if min == math.MaxFloat64 {
		min = 0
	}
	return MakeNumberResult(float64(min))
}

func max(args []Result, isMaxA bool) Result {
	fName := "MAX"
	if isMaxA {
		fName = "MAXA"
	}
	if len(args) == 0 {
		return MakeErrorResult(fName + " requires at least one argument")
	}
	v := -math.MaxFloat64
	for _, a := range args {
		switch a.Type {
		case ResultTypeNumber:
			if (isMaxA || !a.IsBoolean) && a.ValueNumber > v {
				v = a.ValueNumber
			}
		case ResultTypeList, ResultTypeArray:
			subMax := max(a.ListValues(), isMaxA)
			if subMax.ValueNumber > v {
				v = subMax.ValueNumber
			}
		case ResultTypeEmpty:
			// skip
		case ResultTypeString:
			crit := 0.0
			if isMaxA {
				crit = a.AsNumber().ValueNumber
			}
			// treated as zero by Excel
			if crit > v {
				v = crit
			}
		default:
			unioffice.Log("unhandled "+fName+"() argument type %s", a.Type)
		}
	}
	if v == -math.MaxFloat64 {
		v = 0
	}
	return MakeNumberResult(v)
}

// Max is an implementation of the Excel MAX() function.
func Max(args []Result) Result {
	return max(args, false)
}

// MaxA is an implementation of the Excel MAXA() function.
func MaxA(args []Result) Result {
	return max(args, true)
}

func min(args []Result, isMinA bool) Result {
	fName := "MIN"
	if isMinA {
		fName = "MINA"
	}
	if len(args) == 0 {
		return MakeErrorResult(fName + " requires at least one argument")
	}
	v := math.MaxFloat64
	for _, a := range args {
		switch a.Type {
		case ResultTypeNumber:
			if (isMinA || !a.IsBoolean) && a.ValueNumber < v {
				v = a.ValueNumber
			}
		case ResultTypeList, ResultTypeArray:
			subMin := min(a.ListValues(), isMinA)
			if subMin.ValueNumber < v {
				v = subMin.ValueNumber
			}
		case ResultTypeEmpty:
			// skip
		case ResultTypeString:
			crit := 0.0
			if isMinA {
				crit = a.AsNumber().ValueNumber
			}
			// treated as zero by Excel
			if crit < v {
				v = crit
			}
		default:
			unioffice.Log("unhandled "+fName+"() argument type %s", a.Type)
		}
	}
	if v == math.MaxFloat64 {
		v = 0
	}
	return MakeNumberResult(v)
}

// Min is an implementation of the Excel MIN() function.
func Min(args []Result) Result {
	return min(args, false)
}

// MinA is an implementation of the Excel MINA() function.
func MinA(args []Result) Result {
	return min(args, true)
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
			if !a.IsBoolean {
				values = append(values, a.ValueNumber)
			}
		case ResultTypeList, ResultTypeArray:
			values = append(values, extractNumbers(a.ListValues())...)
		case ResultTypeString:
			// treated as zero by Excel
		default:
			unioffice.Log("unhandled extractNumbers argument type %s", a.Type)
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

func compare(value Result, criteria *criteriaParsed) bool {
	if value.IsBoolean {
		return false
	}
	t := value.Type
	if criteria.isNumber {
		return t == ResultTypeNumber && value.ValueNumber == criteria.cNum
	} else if t == ResultTypeNumber {
		return compareNumberWithRegex(value.ValueNumber, criteria.cRegex)
	}
	return compareStrings(value, criteria)
}

func compareStrings(valueResult Result, criteria *criteriaParsed) bool {
	value := strings.ToLower(valueResult.ValueString) // Excel compares string case-insensitive

	regexType := criteria.cRegex.regexType
	compareWith := criteria.cRegex.compareWith

	if regexType == isEq {
		return value == compareWith || wildcard.Match(compareWith, value)
	}
	if valueResult.Type != ResultTypeEmpty { // the only case when empty result should be taken into account is 'equal' condition like "="&A1 which is handled above, other cases with empty cells (including just A1) should be skipped
		if value == criteria.cStr || wildcard.Match(criteria.cStr, value) {
			return true
		}
		if _, err := strconv.ParseFloat(compareWith, 64); err == nil { // criteria should't be the number when compared to string (except 'equal' condition which is handled above)
			return false
		}
		switch regexType {
		case isLe:
			return value <= compareWith // office apps use the same string comparison as in Golang
		case isGe:
			return value >= compareWith
		case isL:
			return value < compareWith
		case isG:
			return value > compareWith
		}
	}

	return false
}

func compareNumberWithRegex(value float64, cRegex *criteriaRegex) bool {
	compareWith, err := strconv.ParseFloat(cRegex.compareWith, 64)
	if err != nil {
		return false
	}

	switch cRegex.regexType {
	case isEq:
		return value == compareWith
	case isLe:
		return value <= compareWith
	case isGe:
		return value >= compareWith
	case isL:
		return value < compareWith
	case isG:
		return value > compareWith
	}

	return false
}
