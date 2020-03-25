// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

func init() {
	initRegexpTime()
	RegisterFunction("DATE", Date)
	RegisterFunction("DATEDIF", DateDif)
	RegisterFunction("DATEVALUE", DateValue)
	RegisterFunction("DAY", Day)
	RegisterFunction("DAYS", Days)
	RegisterFunction("_xlfn.DAYS", Days)
	RegisterFunction("EDATE", Edate)
	RegisterFunction("EOMONTH", Eomonth)
	RegisterFunction("MINUTE", Minute)
	RegisterFunction("MONTH", Month)
	RegisterFunction("NOW", Now)
	RegisterFunction("TIME", Time)
	RegisterFunction("TIMEVALUE", TimeValue)
	RegisterFunction("TODAY", Today)
	RegisterFunctionComplex("YEAR", Year)
	RegisterFunction("YEARFRAC", YearFrac)
}

var date1900 int64 = makeDateS(1900, time.January, 1)
var daysTo1970 float64 = 25569.0

var daysInMonth = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

var month2num = map[string]int{
	"january":   1,
	"february":  2,
	"march":     3,
	"april":     4,
	"may":       5,
	"june":      6,
	"july":      7,
	"august":    8,
	"septemper": 9,
	"october":   10,
	"november":  11,
	"december":  12,
	"jan":       1,
	"feb":       2,
	"mar":       3,
	"apr":       4,
	"jun":       6,
	"jul":       7,
	"aug":       8,
	"sep":       9,
	"oct":       10,
	"nov":       11,
	"dec":       12,
}

var dateFormats = map[string]*regexp.Regexp{}
var timeFormats = map[string]*regexp.Regexp{}
var dateOnlyFormats = []*regexp.Regexp{}
var timeOnlyFormats = []*regexp.Regexp{}

const monthRe = `((jan|january)|(feb|february)|(mar|march)|(apr|april)|(may)|(jun|june)|(jul|july)|(aug|august)|(sep|september)|(oct|october)|(nov|november)|(dec|december))`

const df1 = `(([0-9])+)/(([0-9])+)/(([0-9])+)`
const df2 = monthRe + ` (([0-9])+), (([0-9])+)`
const df3 = `(([0-9])+)-(([0-9])+)-(([0-9])+)`
const df4 = `(([0-9])+)-` + monthRe + `-(([0-9])+)`
const datePrefix = `^((` + df1 + `|` + df2 + `|` + df3 + `|` + df4 + `) )?`

const tfhh = `(([0-9])+) (am|pm)`
const tfhhmm = `(([0-9])+):(([0-9])+)( (am|pm))?`
const tfmmss = `(([0-9])+):(([0-9])+\.([0-9])+)( (am|pm))?`
const tfhhmmss = `(([0-9])+):(([0-9])+):(([0-9])+(\.([0-9])+)?)( (am|pm))?`
const timeSuffix = `( (` + tfhh + `|` + tfhhmm + `|` + tfmmss + `|` + tfhhmmss + `))?$`

func initRegexpTime() {
	dateFormats["mm/dd/yy"] = regexp.MustCompile(`^` + df1 + timeSuffix)
	dateFormats["mm dd, yy"] = regexp.MustCompile(`^` + df2 + timeSuffix)
	dateFormats["yy-mm-dd"] = regexp.MustCompile(`^` + df3 + timeSuffix)
	dateFormats["yy-mmStr-dd"] = regexp.MustCompile(`^` + df4 + timeSuffix)
	timeFormats["hh"] = regexp.MustCompile(datePrefix + tfhh + `$`)
	timeFormats["hh:mm"] = regexp.MustCompile(datePrefix + tfhhmm + `$`)
	timeFormats["mm:ss"] = regexp.MustCompile(datePrefix + tfmmss + `$`)
	timeFormats["hh:mm:ss"] = regexp.MustCompile(datePrefix + tfhhmmss + `$`)
	dateOnlyFormats = []*regexp.Regexp{
		regexp.MustCompile(`^` + df1 + `$`),
		regexp.MustCompile(`^` + df2 + `$`),
		regexp.MustCompile(`^` + df3 + `$`),
		regexp.MustCompile(`^` + df4 + `$`),
	}
	timeOnlyFormats = []*regexp.Regexp{
		regexp.MustCompile(`^` + tfhh + `$`),
		regexp.MustCompile(`^` + tfhhmm + `$`),
		regexp.MustCompile(`^` + tfmmss + `$`),
		regexp.MustCompile(`^` + tfhhmmss + `$`),
	}
}

var empty Result = MakeEmptyResult()

// Day is an implementation of the Excel DAY() function.
func Day(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("DAY requires one argument")
	}
	dateArg := args[0]
	switch dateArg.Type {
	case ResultTypeEmpty:
		return MakeNumberResult(0)
	case ResultTypeNumber:
		date := dateFromDays(dateArg.ValueNumber)
		return MakeNumberResult(float64(date.Day()))
	case ResultTypeString:
		dateString := strings.ToLower(dateArg.ValueString)
		if !checkDateOnly(dateString) { // If time also presents in string, we should validate it first as Excel does
			_, _, _, _, dateIsEmpty, errResult := timeValue(dateString)
			if errResult.Type == ResultTypeError {
				errResult.ErrorMessage = "Incorrect arguments for DAY"
				return errResult
			}
			if dateIsEmpty {
				return MakeNumberResult(0)
			}
		}
		_, _, day, _, errResult := dateValue(dateString)
		if errResult.Type == ResultTypeError {
			return errResult
		}
		return MakeNumberResult(float64(day))
	default:
		return MakeErrorResult("Incorrect argument for DAY")
	}
}

// Days is an implementation of the Excel DAYS() function.
func Days(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("DAYS requires two arguments")
	}
	var sd, ed float64
	switch args[0].Type {
	case ResultTypeNumber:
		ed = args[0].ValueNumber
	case ResultTypeString:
		edResult := DateValue([]Result{args[0]})
		if edResult.Type == ResultTypeError {
			return MakeErrorResult("Incorrect end date for DAYS")
		}
		ed = edResult.ValueNumber
	default:
		return MakeErrorResult("Incorrect argument for DAYS")
	}
	switch args[1].Type {
	case ResultTypeNumber:
		sd = args[1].ValueNumber
		if sd < 62 && ed >= 62 {
			sd--
		}
	case ResultTypeString:
		sdResult := DateValue([]Result{args[1]})
		if sdResult.Type == ResultTypeError {
			return MakeErrorResult("Incorrect start date for DAYS")
		}
		sd = sdResult.ValueNumber
	default:
		return MakeErrorResult("Incorrect argument for DAYS")
	}
	days := float64(int(ed - sd))
	return MakeNumberResult(days)
}

// Date is an implementation of the Excel DATE() function.
func Date(args []Result) Result {
	if len(args) != 3 || args[0].Type != ResultTypeNumber || args[1].Type != ResultTypeNumber || args[2].Type != ResultTypeNumber {
		return MakeErrorResult("DATE requires three number arguments")
	}
	year := int(args[0].ValueNumber)
	if year < 0 || year >= 10000 {
		return MakeErrorResultType(ErrorTypeNum, "Incorrect date")
	} else if year <= 1899 {
		year += 1900
	}
	month := time.Month(args[1].ValueNumber)
	day := int(args[2].ValueNumber)
	dateS := makeDateS(year, month, day)
	days := daysBetween(date1900, dateS) + 1
	if days < 0 {
		return MakeErrorResultType(ErrorTypeNum, "Incorrect date")
	}
	return MakeNumberResult(days)
}

const nsPerDay = 86400000000000

func dateFromDays(days float64) time.Time {
	unix := int64((days - daysTo1970) * nsPerDay)
	return time.Unix(0, unix)
}

func daysFromDate(y, m, d int) float64 {
	return float64(makeDateS(y, time.Month(m), d)/86400) + daysTo1970
}

// DateDif is an implementation of the Excel DATEDIF() function.
func DateDif(args []Result) Result {
	if len(args) != 3 || args[0].Type != ResultTypeNumber || args[1].Type != ResultTypeNumber || args[2].Type != ResultTypeString {
		return MakeErrorResult("DATEDIF requires two number and one string argument")
	}
	startDateDays := args[0].ValueNumber
	endDateDays := args[1].ValueNumber
	if endDateDays < startDateDays {
		return MakeErrorResultType(ErrorTypeNum, "The start date is greater than the end date")
	}
	if endDateDays == startDateDays {
		return MakeNumberResult(0)
	}
	interval := strings.ToLower(args[2].ValueString)
	if interval == "d" {
		return MakeNumberResult(endDateDays - startDateDays)
	}
	startDate := dateFromDays(startDateDays)
	endDate := dateFromDays(endDateDays)
	sy, smm, sd := startDate.Date()
	ey, emm, ed := endDate.Date()
	sm := int(smm)
	em := int(emm)
	var diff float64
	switch interval {
	case "y":
		diff = float64(ey - sy)
		if em < sm || (em == sm && ed < sd) {
			diff--
		}
	case "m":
		ydiff := ey - sy
		mdiff := em - sm
		if ed < sd {
			mdiff--
		}
		if mdiff < 0 {
			ydiff--
			mdiff += 12
		}
		diff = float64(ydiff*12 + mdiff)
	case "md":
		smMD := em
		if ed < sd {
			smMD--
		}
		diff = float64(int(endDateDays - daysFromDate(ey, smMD, sd)))
	case "ym":
		diff = float64(em - sm)
		if ed < sd {
			diff--
		}
		if diff < 0 {
			diff += 12
		}
	case "yd":
		syYD := ey
		if em < sm || (em == sm && ed < sd) {
			syYD--
		}
		diff = float64(int(endDateDays - daysFromDate(syYD, sm, sd)))
	default:
		return MakeErrorResultType(ErrorTypeNum, "Incorrect interval value")
	}
	return MakeNumberResult(diff)
}

const dvErrMsg = "Incorrect argument for DATEVALUE"

// DateValue is an implementation of the Excel DATEVALUE() function.
func DateValue(args []Result) Result {
	if len(args) != 1 || args[0].Type != ResultTypeString {
		return MakeErrorResult("DATEVALUE requires a single string arguments")
	}
	dateString := strings.ToLower(args[0].ValueString)
	if !checkDateOnly(dateString) { // If time also presents in string, we should validate it first as Excel does
		_, _, _, _, dateIsEmpty, errResult := timeValue(dateString)
		if errResult.Type == ResultTypeError {
			errResult.ErrorMessage = "Incorrect arguments for DATEVALUE"
			return errResult
		}
		if dateIsEmpty {
			return MakeNumberResult(0)
		}
	}
	year, month, day, _, errResult := dateValue(dateString)
	if errResult.Type == ResultTypeError {
		return errResult
	}
	return MakeNumberResult(daysFromDate(year, month, day))
}

func checkDateOnly(dateString string) bool {
	for _, df := range dateOnlyFormats {
		submatch := df.FindStringSubmatch(dateString)
		if len(submatch) > 1 {
			return true
		}
	}
	return false
}

// dateValue is a helper for DateValue which is used also by TimeValue to validate date part of the string.
// It returns output in a format of (hours, minutes, seconds, timeIsEmpty, errorResult), where timeIsEmpty is true if the input string contains only date (say, 11/11/2019, not 11/11/2019 12:14:18), errorResult is of ResultTypeError if an error occurs and ResultTypeEmpty if not.
func dateValue(dateString string) (int, int, int, bool, Result) {
	pattern := ""
	submatch := []string{}
	for key, df := range dateFormats {
		submatch = df.FindStringSubmatch(dateString)
		if len(submatch) > 1 {
			pattern = key
			break
		}
	}
	if pattern == "" {
		return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, dvErrMsg)
	}
	timeIsEmpty := false

	var year, month, day int
	var err error

	switch pattern {
	case "mm/dd/yy":
		month, err = strconv.Atoi(submatch[1])
		if err != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, dvErrMsg)
		}
		day, err = strconv.Atoi(submatch[3])
		if err != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, dvErrMsg)
		}
		year, err = strconv.Atoi(submatch[5])
		if err != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, dvErrMsg)
		}
		if year < 0 || year > 9999 || (year > 99 && year < 1900) {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, dvErrMsg)
		}
		year = modifyYear(year)
		timeIsEmpty = submatch[8] == ""
	case "mm dd, yy":
		month = month2num[submatch[1]]
		day, err = strconv.Atoi(submatch[14])
		if err != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, dvErrMsg)
		}
		year, err = strconv.Atoi(submatch[16])
		if err != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, dvErrMsg)
		}
		if year < 0 || year > 9999 || (year > 99 && year < 1900) {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, dvErrMsg)
		}
		year = modifyYear(year)
		timeIsEmpty = submatch[19] == ""
	case "yy-mm-dd":
		v1, err := strconv.Atoi(submatch[1])
		if err != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, dvErrMsg)
		}
		v2, err := strconv.Atoi(submatch[3])
		if err != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, dvErrMsg)
		}
		v3, err := strconv.Atoi(submatch[5])
		if err != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, dvErrMsg)
		}
		if v1 >= 1900 && v1 < 10000 {
			year = v1
			month = v2
			day = v3
		} else if v1 > 0 && v1 < 13 {
			month = v1
			day = v2
			year = v3
		} else {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, dvErrMsg)
		}
		timeIsEmpty = submatch[8] == ""
	case "yy-mmStr-dd":
		year, err = strconv.Atoi(submatch[16])
		if err != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, dvErrMsg)
		}
		month = month2num[submatch[3]]
		day, err = strconv.Atoi(submatch[1])
		if err != nil {
			return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, dvErrMsg)
		}
		timeIsEmpty = submatch[19] == ""
	}
	if !validateDate(year, month, day) {
		return 0, 0, 0, false, MakeErrorResultType(ErrorTypeValue, dvErrMsg)
	}
	return year, month, day, timeIsEmpty, empty
}

func validateDate(year, month, day int) bool {
	if month < 1 || month > 12 {
		return false
	}
	if day < 1 {
		return false
	}
	return day <= getDaysInMonth(year, month)
}

func getDaysInMonth(year, month int) int {
	if month == 2 && isLeapYear(year) {
		return 29
	} else {
		return daysInMonth[month-1]
	}
}

func modifyYear(year int) int {
	if year < 1900 {
		if year < 30 {
			year += 2000
		} else {
			year += 1900
		}
	}
	return year
}

// Edate is an implementation of the Excel EDATE() function.
func Edate(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("EDATE requires two arguments")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("Incorrect argument for EDATE")
	}
	mDelta := args[1].ValueNumber
	initialDateArg := args[0]
	var initialDateDay float64
	switch initialDateArg.Type {
	case ResultTypeEmpty:
		return MakeErrorResultType(ErrorTypeNum, "Incorrect argument for EDATE")
	case ResultTypeNumber:
		initialDateDay = initialDateArg.ValueNumber
	case ResultTypeString:
		initialDateDayResult := DateValue([]Result{args[0]})
		if initialDateDayResult.Type == ResultTypeError {
			return MakeErrorResult("Incorrect argument for EDATE")
		}
		initialDateDay = initialDateDayResult.ValueNumber
	default:
		return MakeErrorResult("Incorrect argument for EDATE")
	}
	initialDate := dateFromDays(initialDateDay)
	newDate := initialDate.AddDate(0, int(mDelta), 0)
	y, m, d := newDate.Date()
	newDays := daysFromDate(y, int(m), d)
	if newDays < 1 {
		return MakeErrorResultType(ErrorTypeNum, "Incorrect argument for EDATE")
	}
	return MakeNumberResult(newDays)
}

// Eomonth is an implementation of the Excel EOMONTH() function.
func Eomonth(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("EOMONTH requires two arguments")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("Incorrect argument for EOMONTH")
	}
	mDelta := args[1].ValueNumber
	initialDateArg := args[0]
	var initialDateDay float64
	switch initialDateArg.Type {
	case ResultTypeEmpty:
		initialDateDay = 0
	case ResultTypeNumber:
		initialDateDay = initialDateArg.ValueNumber
	case ResultTypeString:
		initialDateDayResult := DateValue([]Result{args[0]})
		if initialDateDayResult.Type == ResultTypeError {
			return MakeErrorResult("Incorrect argument for EOMONTH")
		}
		initialDateDay = initialDateDayResult.ValueNumber
	default:
		return MakeErrorResult("Incorrect argument for EOMONTH")
	}
	initialDate := dateFromDays(initialDateDay)
	newDate := initialDate.AddDate(0, int(mDelta+1), 0)
	y, m, _ := newDate.Date()
	eomonth := daysFromDate(y, int(m), 0)
	if eomonth < 1 {
		return MakeErrorResultType(ErrorTypeNum, "Incorrect argument for EOMONTH")
	}
	if y == 1900 && m == 3 {
		eomonth--
	}
	return MakeNumberResult(eomonth)
}

// Minute is an implementation of the Excel MINUTE() function.
func Minute(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("MINUTE requires one argument")
	}
	timeArg := args[0]
	switch timeArg.Type {
	case ResultTypeEmpty:
		return MakeNumberResult(0)
	case ResultTypeNumber:
		date := dateFromDays(timeArg.ValueNumber)
		return MakeNumberResult(float64(date.Minute()))
	case ResultTypeString:
		timeString := strings.ToLower(timeArg.ValueString)
		if !checkTimeOnly(timeString) { // If date also presents in string, we should validate it first as Excel does
			_, _, _, timeIsEmpty, errResult := dateValue(timeString)
			if errResult.Type == ResultTypeError {
				errResult.ErrorMessage = "Incorrect arguments for MINUTE"
				return errResult
			}
			if timeIsEmpty {
				return MakeNumberResult(0)
			}
		}
		_, minute, _, _, _, errResult := timeValue(timeString)
		if errResult.Type == ResultTypeError {
			return errResult
		}
		return MakeNumberResult(float64(minute))
	default:
		return MakeErrorResult("Incorrect argument for MINUTE")
	}
}

// Month is an implementation of the Excel MONTH() function.
func Month(args []Result) Result {
	if len(args) != 1 {
		return MakeErrorResult("MONTH requires one argument")
	}
	dateArg := args[0]
	switch dateArg.Type {
	case ResultTypeEmpty:
		return MakeNumberResult(1)
	case ResultTypeNumber:
		date := dateFromDays(dateArg.ValueNumber)
		return MakeNumberResult(float64(date.Month()))
	case ResultTypeString:
		dateString := strings.ToLower(dateArg.ValueString)
		if !checkDateOnly(dateString) { // If time also presents in string, we should validate it first as Excel does
			_, _, _, _, dateIsEmpty, errResult := timeValue(dateString)
			if errResult.Type == ResultTypeError {
				errResult.ErrorMessage = "Incorrect arguments for MONTH"
				return errResult
			}
			if dateIsEmpty {
				return MakeNumberResult(1)
			}
		}
		_, month, _, _, errResult := dateValue(dateString)
		if errResult.Type == ResultTypeError {
			return errResult
		}
		return MakeNumberResult(float64(month))
	default:
		return MakeErrorResult("Incorrect argument for MONTH")
	}
}

// Now is an implementation of the Excel NOW() function.
func Now(args []Result) Result {
	if len(args) > 0 {
		return MakeErrorResult("NOW doesn't require arguments")
	}
	now := time.Now()
	_, offset := now.Zone()
	nowS := daysTo1970 + float64(now.Unix()+int64(offset))/86400
	return MakeNumberResult(nowS)
}

// Today is an implementation of the Excel TODAY() function.
func Today(args []Result) Result {
	if len(args) > 0 {
		return MakeErrorResult("TODAY doesn't require arguments")
	}
	now := time.Now()
	_, offset := now.Zone()
	nowS := daysBetween(date1900, now.Unix()+int64(offset)) + 1
	return MakeNumberResult(nowS)
}

func daysFromTime(hours, minutes, seconds float64) float64 {
	return (hours*3600 + minutes*60 + seconds) / 86400
}

// Time is an implementation of the Excel TIME() function.
func Time(args []Result) Result {
	if len(args) != 3 || args[0].Type != ResultTypeNumber || args[1].Type != ResultTypeNumber || args[2].Type != ResultTypeNumber {
		return MakeErrorResult("TIME requires three number arguments")
	}
	hours := args[0].ValueNumber
	minutes := args[1].ValueNumber
	seconds := args[2].ValueNumber
	total := daysFromTime(hours, minutes, seconds)
	if total >= 0 {
		return MakeNumberResult(total)
	} else {
		return MakeErrorResultType(ErrorTypeNum, "")
	}
}

const tvErrMsg = "Incorrect argument for TIMEVALUE"

// TimeValue is an implementation of the Excel TIMEVALUE() function.
func TimeValue(args []Result) Result {
	if len(args) != 1 || args[0].Type != ResultTypeString {
		return MakeErrorResult("TIMEVALUE requires a single string arguments")
	}
	timeString := strings.ToLower(args[0].ValueString)
	if !checkTimeOnly(timeString) { // If date also presents in string, we should validate it first as Excel does
		_, _, _, timeIsEmpty, errResult := dateValue(timeString)
		if errResult.Type == ResultTypeError {
			errResult.ErrorMessage = "Incorrect arguments for TIMEVALUE"
			return errResult
		}
		if timeIsEmpty {
			return MakeNumberResult(0)
		}
	}
	hours, minutes, seconds, pm, _, errResult := timeValue(timeString)
	if errResult.Type == ResultTypeError {
		return errResult
	}
	resultValue := daysFromTime(float64(hours), float64(minutes), seconds)
	if pm {
		resultValue += 0.5
	} else if resultValue >= 1 {
		resultValue -= float64(int(resultValue))
	}
	return MakeNumberResult(resultValue)
}

func checkTimeOnly(timeString string) bool {
	for _, tf := range timeOnlyFormats {
		submatch := tf.FindStringSubmatch(timeString)
		if len(submatch) > 1 {
			return true
		}
	}
	return false
}

// timeValue is a helper for TimeValue which is used also by DateValue to validate time part of the string.
// It returns output in a format of (year, month, day, pm, dateIsEmpty, errorResult), where pm is true if the time is marked as PM, dateIsEmpty is true if the input string contains only time (say, 12:14:18, not 11/11/2019 12:14:18), errorResult is of ResultTypeError if an error occurs and ResultTypeEmpty if not.
func timeValue(timeString string) (int, int, float64, bool, bool, Result) {
	pattern := ""
	submatch := []string{}
	for key, tf := range timeFormats {
		submatch = tf.FindStringSubmatch(timeString)
		if len(submatch) > 1 {
			pattern = key
			break
		}
	}
	if pattern == "" {
		return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, tvErrMsg)
	}
	dateIsEmpty := submatch[1] == ""
	submatch = submatch[49:] // cut off date

	l := len(submatch)
	last := submatch[l-1]
	am := last == "am"
	pm := last == "pm"

	var hours, minutes int
	var seconds float64
	var err error

	switch pattern {
	case "hh":
		hours, err = strconv.Atoi(submatch[0])
		if err != nil {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, tvErrMsg)
		}
		minutes = 0
		seconds = 0
	case "hh:mm":
		hours, err = strconv.Atoi(submatch[0])
		if err != nil {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, tvErrMsg)
		}
		minutes, err = strconv.Atoi(submatch[2])
		if err != nil {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, tvErrMsg)
		}
		seconds = 0
	case "mm:ss":
		hours = 0
		minutes, err = strconv.Atoi(submatch[0])
		if err != nil {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, tvErrMsg)
		}
		seconds, err = strconv.ParseFloat(submatch[2], 64)
		if err != nil {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, tvErrMsg)
		}
	case "hh:mm:ss":
		hours, err = strconv.Atoi(submatch[0])
		if err != nil {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, tvErrMsg)
		}
		minutes, err = strconv.Atoi(submatch[2])
		if err != nil {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, tvErrMsg)
		}
		seconds, err = strconv.ParseFloat(submatch[4], 64)
		if err != nil {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, tvErrMsg)
		}
	}
	if minutes >= 60 {
		return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, tvErrMsg)
	}
	if am || pm {
		if hours > 12 || seconds >= 60 {
			return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, tvErrMsg)
		} else if hours == 12 {
			hours = 0
		}
	} else if hours >= 24 || seconds >= 10000 {
		return 0, 0, 0, false, false, MakeErrorResultType(ErrorTypeValue, tvErrMsg)
	}
	return hours, minutes, seconds, pm, dateIsEmpty, empty
}

// Year is an implementation of the Excel YEAR() function.
func Year(ctx Context, ev Evaluator, args []Result) Result {
	if len(args) != 1 || args[0].Type != ResultTypeNumber {
		return MakeErrorResult("YEAR requires a single number argument")
	}
	epoch := ctx.GetEpoch()
	t, err := getValueAsTime(args[0].Value(), epoch)
	if err != nil {
		return MakeErrorResult("YEAR requires a single date argument")
	}
	return MakeNumberResult(float64(t.Year()))
}

// YearFrac is an implementation of the Excel YEARFRAC() function.
func YearFrac(args []Result) Result {
	argsNum := len(args)
	if (argsNum != 2 && argsNum != 3) || args[0].Type != ResultTypeNumber || args[1].Type != ResultTypeNumber {
		return MakeErrorResult("YEARFRAC requires two or three number arguments")
	}

	basis := 0
	if argsNum == 3 && args[2].Type != ResultTypeEmpty {
		if args[2].Type != ResultTypeNumber {
			return MakeErrorResult("YEARFRAC requires basis argument to be a number")
		}
		basis = int(args[2].ValueNumber)
		if !checkBasis(basis) {
			return MakeErrorResultType(ErrorTypeNum, "Incorrect basis argument for YEARFRAC")
		}
	}

	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("YEARFRAC requires start date to be number argument")
	}
	startDate := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("YEARFRAC requires end date to be number argument")
	}
	endDate := args[1].ValueNumber
	yf, errResult := yearFrac(startDate, endDate, basis)
	if errResult.Type == ResultTypeError {
		return errResult
	}
	return MakeNumberResult(yf)
}

// yearFrac returns float64 fraction of the year and Result value which can be of ResultTypeError type if an error occurs or ResultTypeEmpty if doesn't.
func yearFrac(startDateF, endDateF float64, basis int) (float64, Result) {
	startDate, endDate := dateFromDays(startDateF), dateFromDays(endDateF)
	startDateS := startDate.Unix()
	endDateS := endDate.Unix()
	if startDateS == endDateS {
		return 0, empty
	}
	sy, smM, sd := startDate.Date()
	ey, emM, ed := endDate.Date()
	sm, em := int(smM), int(emM)
	var dayDiff, daysInYear float64
	switch basis {
	case 0:
		if sd == 31 {
			sd--
		}
		if sd == 30 && ed == 31 {
			ed--
		} else if leap := isLeapYear(sy); sm == 2 && ((leap && sd == 29) || (!leap && sd == 28)) {
			sd = 30
			if leap := isLeapYear(ey); em == 2 && ((leap && ed == 29) || (!leap && ed == 28)) {
				ed = 30
			}
		}
		dayDiff = float64((ey-sy)*360 + (em-sm)*30 + (ed - sd))
		daysInYear = 360
	case 1:
		dayDiff = endDateF - startDateF
		isYearDifferent := sy != ey
		if isYearDifferent && (ey != sy+1 || sm < em || (sm == em && sd < ed)) {
			dayCount := 0
			for y := sy; y <= ey; y++ {
				dayCount += getDaysInYear(y, 1)
			}
			daysInYear = float64(dayCount) / float64(ey-sy+1)
		} else {
			if !isYearDifferent && isLeapYear(sy) {
				daysInYear = 366
			} else {
				if isYearDifferent && ((isLeapYear(sy) && (sm < 2 || (sm == 2 && sd <= 29))) || (isLeapYear(ey) && (em > 2 || (em == 2 && ed == 29)))) {
					daysInYear = 366
				} else {
					daysInYear = 365
				}
			}
		}
	case 2:
		dayDiff = endDateF - startDateF
		daysInYear = 360
	case 3:
		dayDiff = endDateF - startDateF
		daysInYear = 365
	case 4:
		if sd == 31 {
			sd--
		}
		if ed == 31 {
			ed--
		}
		dayDiff = float64((ey-sy)*360 + (em-sm)*30 + (ed - sd))
		daysInYear = 360
	default:
		return 0, MakeErrorResultType(ErrorTypeNum, "Incorrect basis for YearFrac")
	}
	return dayDiff / daysInYear, empty
}

func getDaysInYear(year, basis int) int {
	switch basis {
	case 1:
		if isLeapYear(year) {
			return 366
		} else {
			return 365
		}
	case 3:
		return 365
	default:
		return 360
	}
}

func makeDateS(y int, m time.Month, d int) int64 {
	if y == 1900 && int(m) <= 2 {
		d--
	}
	date := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	return date.Unix()
}

func isLeapYear(year int) bool {
	if year == year/400*400 {
		return true
	}
	if year == year/100*100 {
		return false
	}
	return year == year/4*4
}

func daysBetween(startDate, endDate int64) float64 {
	return float64(int(0.5 + float64((endDate-startDate)/86400)))
}

func feb29Between(date1, date2 time.Time) bool {
	date1S := date1.Unix()
	date2S := date2.Unix()
	year1 := date1.Year()
	mar1year1 := makeDateS(year1, time.March, 1)
	if isLeapYear(year1) && date1S < mar1year1 && date2S >= mar1year1 {
		return true
	}
	var year2 = date2.Year()
	var mar1year2 = makeDateS(year2, time.March, 1)
	return (isLeapYear(year2) && date2S >= mar1year2 && date1S < mar1year2)
}

func getDiff(from, to time.Time, basis int) float64 {
	if from.After(to) {
		from, to = to, from
	}
	diff := 0

	yFrom, mFromM, dFromOrig := from.Date()
	yTo, mToM, dToOrig := to.Date()

	mFrom, mTo := int(mFromM), int(mToM)
	dFrom, dTo := getDayOnBasis(yFrom, mFrom, dFromOrig, basis), getDayOnBasis(yTo, mTo, dToOrig, basis)

	if !basis30(basis) {
		return daysFromDate(yTo, mTo, dTo) - daysFromDate(yFrom, mFrom, dFrom)
	}

	if basis == 0 {
		if (mFrom == 2 || dFrom < 30) && dToOrig == 31 {
			dTo = 31
		} else if mTo == 2 && dTo == getDaysInMonth(yTo, mTo) {
			dTo = getDaysInMonth(yTo, 2)
		}
	} else {
		if mFrom == 2 && dFrom == 30 {
			dFrom = getDaysInMonth(yFrom, 2)
		}
		if mTo == 2 && dTo == 30 {
			dTo = getDaysInMonth(yTo, 2)
		}
	}

	if yFrom < yTo || (yFrom == yTo && mFrom < mTo) {
		diff = 30 - dFrom + 1
		dFromOrig = 1
		dFrom = 1
		fromNew := time.Date(yFrom, time.Month(mFrom), dFromOrig, 0, 0, 0, 0, time.UTC).AddDate(0, 1, 0)
		if fromNew.Year() < yTo {
			diff += getDaysInMonthRange(fromNew.Year(), int(fromNew.Month()), 12, basis)
			fromNew = fromNew.AddDate(0, 13-int(fromNew.Month()), 0)
			diff += getDaysInYearRange(fromNew.Year(), yTo-1, basis)
		}
		diff += getDaysInMonthRange(yTo, int(fromNew.Month()), mTo-1, basis)
		fromNew = fromNew.AddDate(0, mTo-int(fromNew.Month()), 0)
		mFrom = fromNew.Day()
	}
	diff += dTo - dFrom
	if diff > 0 {
		return float64(diff)
	} else {
		return 0
	}
}

func getDayOnBasis(year, month, dayOrig, basis int) int {
	if !basis30(basis) {
		return dayOrig
	}
	day := dayOrig
	dim := getDaysInMonth(year, month)
	if day > 30 || dayOrig >= dim || day >= dim {
		day = 30
	}
	return day
}

func getDaysInMonthRange(y, from, to, basis int) int {
	if from > to {
		return 0
	}
	if basis30(basis) {
		return (to - from + 1) * 30
	}
	days := 0
	for m := from; m <= to; m++ {
		days += getDaysInMonth(y, m)
	}
	return days
}

func getDaysInYearRange(from, to, basis int) int {
	if from > to {
		return 0
	}
	if basis30(basis) {
		return (to - from + 1) * 360
	}
	days := 0
	for y := from; y <= to; y++ {
		dy := 365
		if isLeapYear(y) {
			dy = 366
		}
		days += dy
	}
	return days
}

func basis30(basis int) bool {
	return basis == 0 || basis == 4
}

func parseDate(arg Result, dateName, funcName string) (float64, Result) {
	var date float64
	switch arg.Type {
	case ResultTypeNumber:
		date = float64(int(arg.ValueNumber))
	case ResultTypeString:
		dateFromString := DateValue([]Result{arg})
		if dateFromString.Type == ResultTypeError {
			return 0, MakeErrorResult("Incorrect " + dateName + " date for " + funcName)
		}
		date = dateFromString.ValueNumber
	default:
		return 0, MakeErrorResult("Incorrect argument for " + funcName)
	}
	if date < 0 {
		return 0, MakeErrorResultType(ErrorTypeNum, dateName+" should be non negative")
	}
	return date, empty
}
