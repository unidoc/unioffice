// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

import (
	"time"
)

func init() {
	RegisterFunctionComplex("YEAR", Year)
	RegisterFunctionComplex("YEARFRAC", YearFrac)
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
