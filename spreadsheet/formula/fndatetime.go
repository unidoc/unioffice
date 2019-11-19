// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting support@unidoc.io.

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
	RegisterFunction("NOW", Now)
	RegisterFunction("TIME", Time)
	RegisterFunction("TIMEVALUE", TimeValue)
	RegisterFunction("TODAY", Today)
	RegisterFunctionComplex("YEAR", Year)
	RegisterFunctionComplex("YEARFRAC", YearFrac)
}

var date1900 int64 = makeDateS(1900, time.January, 0)
var daysTo1970 float64 = 25569.0

var timeFormats map[string]*regexp.Regexp = map[string]*regexp.Regexp{}
const datePrefix = `^(([0-9])+/([0-9])+/([0-9])+ )?`

func initRegexpTime() {
	timeFormats["hh"] = regexp.MustCompile(datePrefix + `(([0-9])+) (am|pm)$`)
	timeFormats["hh:mm"] = regexp.MustCompile(datePrefix + `(([0-9])+):(([0-9])+)( (am|pm))?$`)
	timeFormats["mm:ss"] = regexp.MustCompile(datePrefix + `(([0-9])+):(([0-9])+\.([0-9])+)( (am|pm))?$`)
	timeFormats["hh:mm:ss"] = regexp.MustCompile(datePrefix + `(([0-9])+):(([0-9])+):(([0-9])+(\.([0-9])+)?)( (am|pm))?$`)
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

func daysFromDate(y,m,d int) float64 {
	return float64(makeDateS(y, time.Month(m), d) / 86400) + daysTo1970
}

// DateDif is an implementation of the Excel DATEDIF() function.
func DateDif(args []Result) Result {
	if len(args) != 3 || args[0].Type != ResultTypeNumber || args[1].Type != ResultTypeNumber || args[2].Type != ResultTypeString {
		return MakeErrorResultType(ErrorTypeValue, "DATEDIF requires two number and one string argument")
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

// Now is an implementation of the Excel NOW() function.
func Now(args []Result) Result {
	if len(args) > 0 {
		return MakeErrorResult("NOW doesn't require arguments")
	}
	now := time.Now()
	_, offset := now.Zone()
	nowS := daysTo1970 + float64(now.Unix() + int64(offset))/86400
	return MakeNumberResult(nowS)
}

// Today is an implementation of the Excel TODAY() function.
func Today(args []Result) Result {
	if len(args) > 0 {
		return MakeErrorResult("TODAY doesn't require arguments")
	}
	now := time.Now()
	_, offset := now.Zone()
	nowS := daysBetween(date1900, now.Unix() + int64(offset)) + 1
	return MakeNumberResult(nowS)
}

func daysFromTime(hours, minutes, seconds float64) float64 {
	return (hours * 3600 + minutes * 60 + seconds) / 86400
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
		return MakeErrorResult(tvErrMsg)
	}
	submatch = submatch[5:] // cut off date

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
			return MakeErrorResult(tvErrMsg)
		}
		minutes = 0
		seconds = 0
	case "hh:mm":
		hours, err = strconv.Atoi(submatch[0])
		if err != nil {
			return MakeErrorResult(tvErrMsg)
		}
		minutes, err = strconv.Atoi(submatch[2])
		if err != nil {
			return MakeErrorResult(tvErrMsg)
		}
		seconds = 0
	case "mm:ss":
		hours = 0
		minutes, err = strconv.Atoi(submatch[0])
		if err != nil {
			return MakeErrorResult(tvErrMsg)
		}
		seconds, err = strconv.ParseFloat(submatch[2], 64)
		if err != nil {
			return MakeErrorResult(tvErrMsg)
		}
	case "hh:mm:ss":
		hours, err = strconv.Atoi(submatch[0])
		if err != nil {
			return MakeErrorResult(tvErrMsg)
		}
		minutes, err = strconv.Atoi(submatch[2])
		if err != nil {
			return MakeErrorResult(tvErrMsg)
		}
		seconds, err = strconv.ParseFloat(submatch[4], 64)
		if err != nil {
			return MakeErrorResult(tvErrMsg)
		}
	}
	if minutes >= 60 {
		return MakeErrorResult(tvErrMsg)
	}
	if am || pm {
		if hours > 12 || seconds >= 60 {
			return MakeErrorResult(tvErrMsg)
		} else if hours == 12 {
			hours = 0
		}
	} else if hours >= 24 || seconds >= 10000 {
		return MakeErrorResult(tvErrMsg)
	}
	resultValue := daysFromTime(float64(hours), float64(minutes), seconds)
	if pm {
		resultValue += 0.5
	} else if resultValue >= 1 {
		resultValue -= float64(int(resultValue))
	}
	return MakeNumberResult(resultValue)
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
func YearFrac(ctx Context, ev Evaluator, args []Result) Result {
	argsNum := len(args)
	if (argsNum != 2 && argsNum != 3) || args[0].Type != ResultTypeNumber || args[1].Type != ResultTypeNumber {
		return MakeErrorResult("YEARFRAC requires two or three number arguments")
	}

	basis := 0
	if argsNum == 3 {
		if args[2].Type != ResultTypeNumber {
			return MakeErrorResult("YEARFRAC requires two or three number arguments")
		}
		basis = int(args[2].ValueNumber)
	}

	epoch := ctx.GetEpoch()
	startDate, err := getValueAsTime(args[0].Value(), epoch)
	if err != nil {
		return MakeErrorResult("incorrect start date")
	}
	startDateS := startDate.Unix()
	endDate, err := getValueAsTime(args[1].Value(), epoch)
	if err != nil {
		return MakeErrorResult("incorrect end date")
	}
	endDateS := endDate.Unix()
	sy, sm, sd := startDate.Date()
	ey, em, ed := endDate.Date()

	switch basis {
	case 0:
		if sd == 31 && ed == 31 {
			sd = 30
			ed = 30
		} else if sd == 31 {
			sd = 30
		} else if sd == 30 && ed == 31 {
			ed = 30
		}
		return MakeNumberResult(float64(((ed + int(em) * 30 + ey * 360) - (sd + int(sm) * 30 + sy * 360))) / 360)
	case 1:
		var ylength = 365.0
		if (sy == ey || ((sy + 1) == ey) && ((sm > em) || ((sm == em) && (sd >= ed)))) {
			if ((sy == ey && isLeapYear(sy)) || feb29Between(startDate, endDate) || (em == 1 && ed == 29)) {
				ylength = 366.0
			}
			return MakeNumberResult(daysBetween(startDateS, endDateS) / ylength)
		}
		var years = float64((ey - sy) + 1)
		var days = float64((makeDateS(ey + 1, time.January, 1) - makeDateS(sy, time.January, 1)) / 86400)
		var average = days / years
		return MakeNumberResult(daysBetween(startDateS, endDateS) / average)
	case 2:
		return MakeNumberResult(daysBetween(startDateS, endDateS) / 360.0)
	case 3:
		return MakeNumberResult(daysBetween(startDateS, endDateS) / 365.0)
	case 4:
		return MakeNumberResult(float64(((ed + int(em) * 30 + ey * 360) - (sd + int(sm) * 30 + sy * 360))) / 360.0)
	}
	return MakeErrorResultType(ErrorTypeValue, "")
}

func makeDateS(y int, m time.Month, d int) int64 {
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
	return float64(int(0.5 + float64((endDate - startDate) / 86400)))
}

func feb29Between(date1, date2 time.Time) bool {
	date1S := date1.Unix()
	date2S := date2.Unix()
	year1 := date1.Year()
	mar1year1 := makeDateS(year1, time.March, 1)
	if (isLeapYear(year1) && date1S < mar1year1 && date2S >= mar1year1) {
		return true
	}
	var year2 = date2.Year()
	var mar1year2 = makeDateS(year2, time.March, 1)
	return (isLeapYear(year2) && date2S >= mar1year2 && date1S < mar1year2)
}
