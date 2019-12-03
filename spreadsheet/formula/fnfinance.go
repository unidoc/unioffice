// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import (
	//"fmt"
	"time"
	"math"
)

func init() {
	RegisterFunction("DURATION", Duration)
	RegisterFunction("MDURATION", MDuration)
}

// getCouppcd finds last coupon date before settlement (can be equal to settlement).
func getCouppcd(settlementDate, maturityDate time.Time, freq int) time.Time {
	rDate := maturityDate
	diffYears := settlementDate.Year() - maturityDate.Year()
	rDate = rDate.AddDate(diffYears, 0, 0)
	if settlementDate.After(rDate) {
		rDate = rDate.AddDate(1, 0, 0)
	}
	monthsToAdd := -12 / freq // FREQ
	for rDate.After(settlementDate) {
		rDate = rDate.AddDate(0, monthsToAdd, 0)
	}
	return rDate
}

// getCoupnum gets count of coupon dates.
func getCoupnum(settlementDate, maturityDate time.Time, freq, basis int) float64 {
	if maturityDate.After(settlementDate) {
		aDate := getCouppcd(settlementDate, maturityDate, freq)
		months := (maturityDate.Year() - aDate.Year()) * 12 + int(maturityDate.Month()) - int(aDate.Month())
		return float64(months * freq) / 12.0 // FREQ
	}
	return 0.0 // replace for error
}

// getDuration returns the Macauley duration for an assumed par value of $100. It is defined as the weighted average of the present value of cash flows, and is used as a measure of a bond price's response to changes in yield.
func getDuration(settlementDate, maturityDate time.Time, coup, yield, freq float64, basis int) Result {
	fracResult := yearFrac(settlementDate, maturityDate, basis)
	if fracResult.Type == ResultTypeError {
		return fracResult
	}
	frac := fracResult.ValueNumber
	nCoups := getCoupnum(settlementDate, maturityDate, int(freq), basis)
	duration := 0.0
	p := 0.0
	coup *= 100 / float64(freq)
	yield /= float64(freq) // FREQ
	yield++
	diff := frac * float64(freq) - nCoups // FREQ
	for t := 1.0; t < nCoups; t++ {
		tDiff := t + diff
		add := coup / math.Pow(yield, tDiff)
		p += add
		duration += tDiff * add
	}

	add := (coup + 100) / math.Pow(yield, nCoups + diff)

	p += add
	duration += (nCoups + diff) * add

	duration /= p
	duration /= float64(freq) // FREQ

	return MakeNumberResult(duration)
}

// Duration implements the Excel DURATION function.
func Duration(args []Result) Result {
	settlementDate, maturityDate, coupon, yield, freq, basis, err := parseDurationData(args, "DURATION")
	if err.Type == ResultTypeError {
		return err
	}
	return getDuration(dateFromDays(settlementDate), dateFromDays(maturityDate), coupon, yield, freq, basis)
}

// MDuration implements the Excel MDURATION function.
func MDuration(args []Result) Result {
	settlementDate, maturityDate, coupon, yield, freq, basis, err := parseDurationData(args, "MDURATION")
	if err.Type == ResultTypeError {
		return err
	}
	duration := getDuration(dateFromDays(settlementDate), dateFromDays(maturityDate), coupon, yield, freq, basis)
	if duration.Type == ResultTypeError {
		return duration
	}
	mDuration := duration.ValueNumber / (1.0 + yield / float64(freq))
	return MakeNumberResult(mDuration)
}

// validateDurationData returns settlement date, maturity date, coupon rate, yield rate, frequency of payments, day count basis and error result by parsing incoming arguments
func parseDurationData(args []Result, funcName string) (float64, float64, float64, float64, float64, int, Result) {
	if len(args) != 5 && len(args) != 6 {
		return 0, 0, 0, 0, 0, 0, MakeErrorResult(funcName + " requires five or six arguments")
	}
	var settlementDate, maturityDate float64
	settlementResult := args[0]
	switch settlementResult.Type {
	case ResultTypeNumber:
		settlementDate = float64(int(settlementResult.ValueNumber))
	case ResultTypeString:
		settlementFromString := DateValue([]Result{settlementResult})
		if settlementFromString.Type == ResultTypeError {
			return 0, 0, 0, 0, 0, 0, MakeErrorResult("Incorrect settltment date for " + funcName)
		}
		settlementDate = settlementFromString.ValueNumber
	default:
		return 0, 0, 0, 0, 0, 0, MakeErrorResult("Incorrect argument for " + funcName)
	}
	maturityResult := args[1]
	switch maturityResult.Type {
	case ResultTypeNumber:
		maturityDate = float64(int(maturityResult.ValueNumber))
	case ResultTypeString:
		maturityFromString := DateValue([]Result{maturityResult})
		if maturityFromString.Type == ResultTypeError {
			return 0, 0, 0, 0, 0, 0, MakeErrorResult("Incorrect settltment date for " + funcName)
		}
		maturityDate = maturityFromString.ValueNumber
	default:
		return 0, 0, 0, 0, 0, 0, MakeErrorResult("Incorrect argument for " + funcName)
	}
	if settlementDate >= maturityDate {
		return 0, 0, 0, 0, 0, 0, MakeErrorResultType(ErrorTypeNum, "Settlement date should be before maturity date")
	}
	couponResult := args[2]
	if couponResult.Type != ResultTypeNumber {
		return 0, 0, 0, 0, 0, 0, MakeErrorResult(funcName + " requires third argument of type number")
	}
	coupon := couponResult.ValueNumber
	if coupon < 0 {
		return 0, 0, 0, 0, 0, 0, MakeErrorResultType(ErrorTypeNum, "Coupon rate should not be negative")
	}
	yieldResult := args[3]
	if yieldResult.Type != ResultTypeNumber {
		return 0, 0, 0, 0, 0, 0, MakeErrorResult(funcName + " requires fourth argument of type number")
	}
	yield := yieldResult.ValueNumber
	if yield < 0 {
		return 0, 0, 0, 0, 0, 0, MakeErrorResultType(ErrorTypeNum, "Yield rate should not be negative")
	}
	freqResult := args[4]
	if freqResult.Type != ResultTypeNumber {
		return 0, 0, 0, 0, 0, 0, MakeErrorResult(funcName + " requires fifth argument of type number")
	}
	freq := float64(int(freqResult.ValueNumber))
	if freq != 1 && freq != 2 && freq != 4 {
		return 0, 0, 0, 0, 0, 0, MakeErrorResultType(ErrorTypeNum, "Incorrect frequence value")
	}
	basis := 0
	if len(args) == 6 {
		basisResult := args[5]
		if basisResult.Type != ResultTypeNumber {
			return 0, 0, 0, 0, 0, 0, MakeErrorResult(funcName + " requires sixth argument of type number")
		}
		basis = int(basisResult.ValueNumber)
		if basis < 0 || basis > 4 {
			return 0, 0, 0, 0, 0, 0, MakeErrorResultType(ErrorTypeNum, "Incorrect basis value")
		}
	}
	return settlementDate, maturityDate, coupon, yield, freq, basis, MakeEmptyResult()
}
