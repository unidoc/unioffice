// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import (
	"time"
	"math"
	"strconv"
	"strings"
)

func init() {
	RegisterFunction("ACCRINTM", Accrintm)
	RegisterFunction("AMORDEGRC", Amordegrc)
	RegisterFunction("AMORLINC", Amorlinc)
	RegisterFunction("COUPDAYBS", Coupdaybs)
	RegisterFunction("COUPDAYS", Coupdays)
	RegisterFunction("COUPDAYSNC", Coupdaysnc)
	RegisterFunction("COUPNUM", Coupnum)
	RegisterFunction("COUPNCD", Coupncd)
	RegisterFunction("COUPPCD", Couppcd)
	RegisterFunction("CUMIPMT", Cumipmt)
	RegisterFunction("CUMPRINC", Cumprinc)
	RegisterFunction("DB", Db)
	RegisterFunction("DDB", Ddb)
	RegisterFunction("DISC", Disc)
	RegisterFunction("DOLLARDE", Dollarde)
	RegisterFunction("DOLLARFR", Dollarfr)
	RegisterFunction("DURATION", Duration)
	RegisterFunction("EFFECT", Effect)
	RegisterFunction("FV", Fv)
	RegisterFunction("FVSCHEDULE", Fvschedule)
	RegisterFunction("INTRATE", Intrate)
	RegisterFunction("IPMT", Ipmt)
	RegisterFunction("MDURATION", Mduration)
	RegisterFunction("PDURATION", Pduration)
	RegisterFunction("_xlfn.PDURATION", Pduration)
}

// Duration implements the Excel DURATION function.
func Duration(args []Result) Result {
	parsedArgs, err := parseDurationData(args, "DURATION")
	if err.Type == ResultTypeError {
		return err
	}
	settlementDate := parsedArgs.settlementDate
	maturityDate := parsedArgs.maturityDate
	coupon := parsedArgs.coupon
	yield := parsedArgs.yield
	freq := parsedArgs.freq
	basis := parsedArgs.basis

	return getDuration(settlementDate, maturityDate, coupon, yield, freq, basis)
}

// Mduration implements the Excel MDURATION function.
func Mduration(args []Result) Result {
	parsedArgs, err := parseDurationData(args, "MDURATION")
	if err.Type == ResultTypeError {
		return err
	}
	settlementDate := parsedArgs.settlementDate
	maturityDate := parsedArgs.maturityDate
	coupon := parsedArgs.coupon
	yield := parsedArgs.yield
	freq := parsedArgs.freq
	basis := parsedArgs.basis

	duration := getDuration(settlementDate, maturityDate, coupon, yield, freq, basis)
	if duration.Type == ResultTypeError {
		return duration
	}
	mDuration := duration.ValueNumber / (1.0 + yield / freq)
	return MakeNumberResult(mDuration)
}

// Pduration implements the Excel PDURATION function.
func Pduration(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("PDURATION requires three number arguments")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("PDURATION requires rate to be number argument")
	}
	rate := args[0].ValueNumber
	if rate <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "PDURATION requires rate to be positive")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("PDURATION requires current value to be number argument")
	}
	currentValue := args[1].ValueNumber
	if currentValue <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "PDURATION requires current value to be positive")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("PDURATION requires specified value to be number argument")
	}
	specifiedValue := args[2].ValueNumber
	if specifiedValue <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "PDURATION requires specified value to be positive")
	}
	return MakeNumberResult((math.Log10(specifiedValue) - math.Log10(currentValue)) / math.Log10(1 + rate))
}

type couponArgs struct {
	settlementDate float64
	maturityDate float64
	freq int
	basis int
}

// Coupdaybs implements the Excel COUPDAYBS function.
func Coupdaybs(args []Result) Result {
	parsedArgs, err := parseCouponArgs(args, "COUPDAYBS")
	if err.Type == ResultTypeError {
		return err
	}
	settlementDate := dateFromDays(parsedArgs.settlementDate)
	maturityDate := dateFromDays(parsedArgs.maturityDate)
	freq := parsedArgs.freq
	basis := parsedArgs.basis
	pcd := couppcd(settlementDate, maturityDate, freq, basis)
	return MakeNumberResult(getDiff(pcd, settlementDate, basis))
}

// Coupdays implements the Excel COUPDAYS function.
func Coupdays(args []Result) Result {
	parsedArgs, err := parseCouponArgs(args, "COUPDAYS")
	if err.Type == ResultTypeError {
		return err
	}
	settlementDate := dateFromDays(parsedArgs.settlementDate)
	maturityDate := dateFromDays(parsedArgs.maturityDate)
	freq := parsedArgs.freq
	basis := parsedArgs.basis
	if basis == 1 {
		pcd := couppcd(settlementDate, maturityDate, freq, 1)
		next := pcd.AddDate(0, 12 / freq, 0)
		return MakeNumberResult(getDiff(pcd, next, basis))
	}
	return MakeNumberResult(float64(getDaysInYear(0, basis)) / float64(freq))
}

// Coupdaysnc implements the Excel COUPDAYSNC function.
func Coupdaysnc(args []Result) Result {
	parsedArgs, err := parseCouponArgs(args, "COUPDAYSNC")
	if err.Type == ResultTypeError {
		return err
	}
	settlementDate := dateFromDays(parsedArgs.settlementDate)
	maturityDate := dateFromDays(parsedArgs.maturityDate)
	freq := parsedArgs.freq
	basis := parsedArgs.basis
	ncd := coupncd(settlementDate, maturityDate, freq)
	return MakeNumberResult(getDiff(settlementDate, ncd, basis))
}

// Couppcd implements the Excel COUPPCD function.
func Couppcd(args []Result) Result {
	parsedArgs, err := parseCouponArgs(args, "COUPPCD")
	if err.Type == ResultTypeError {
		return err
	}
	settlementDate := dateFromDays(parsedArgs.settlementDate)
	maturityDate := dateFromDays(parsedArgs.maturityDate)
	freq := parsedArgs.freq
	basis := parsedArgs.basis
	pcd := couppcd(settlementDate, maturityDate, freq, basis)
	y, m, d := pcd.Date()
	return MakeNumberResult(daysFromDate(y, int(m), d))
}

// Coupnum implements the Excel COUPNUM function.
func Coupnum(args []Result) Result {
	parsedArgs, err := parseCouponArgs(args, "COUPNUM")
	if err.Type == ResultTypeError {
		return err
	}
	settlementDate := dateFromDays(parsedArgs.settlementDate)
	maturityDate := dateFromDays(parsedArgs.maturityDate)
	freq := parsedArgs.freq
	basis := parsedArgs.basis
	cn, err := coupnum(settlementDate, maturityDate, freq, basis)
	if err.Type == ResultTypeError {
		return err
	}
	return MakeNumberResult(cn)
}

// Coupncd implements the Excel COUPNCD function.
func Coupncd(args []Result) Result {
	parsedArgs, err := parseCouponArgs(args, "COUPNCD")
	if err.Type == ResultTypeError {
		return err
	}
	settlementDate := dateFromDays(parsedArgs.settlementDate)
	maturityDate := dateFromDays(parsedArgs.maturityDate)
	freq := parsedArgs.freq
	ncd := coupncd(settlementDate, maturityDate, freq)
	y, m, d := ncd.Date()
	return MakeNumberResult(daysFromDate(y, int(m), d))
}

func coupncd(settlementDate, maturityDate time.Time, freq int) time.Time {
	ncd := time.Date(settlementDate.Year(), maturityDate.Month(), maturityDate.Day(), 0, 0, 0, 0, time.UTC)
	if ncd.After(settlementDate) {
		ncd = ncd.AddDate(-1, 0, 0)
	}
	for !ncd.After(settlementDate) {
		ncd = ncd.AddDate(0, 12 / freq, 0)
	}
	return ncd
}

func parseCouponArgs(args []Result, funcName string) (*couponArgs, Result) {
	argsNum := len(args)
	if argsNum != 3 && argsNum != 4 {
		return nil, MakeErrorResult(funcName + " requires four arguments")
	}
	if args[0].Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires settlement date to be number argument")
	}
	settlementDate := args[0].ValueNumber
	if settlementDate < 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName + " requires settlement date to be non negative")
	}
	if args[1].Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires maturity date to be number argument")
	}
	maturityDate := args[1].ValueNumber
	if maturityDate <= settlementDate {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName + " requires maturity date to be later than settlement date")
	}
	if args[2].Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires frequency to be number argument")
	}
	freq := args[2].ValueNumber
	if !checkFreq(freq) {
		return nil, MakeErrorResult("Incorrect frequency for " + funcName)
	}
	basis := 0
	if argsNum == 4 {
		if args[3].Type != ResultTypeNumber {
			return nil, MakeErrorResult(funcName + " requires basis to be number argument")
		}
		basis = int(args[3].ValueNumber)
		if !checkBasis(basis) {
			return nil, MakeErrorResultType(ErrorTypeNum, "Incorrect basis argument for " + funcName)
		}
	}
	return &couponArgs{
		settlementDate,
		maturityDate,
		int(freq),
		basis,
	}, MakeEmptyResult()
}

// couppcd finds last coupon date before settlement (can be equal to settlement).
func couppcd(settlementDate, maturityDate time.Time, freq, basis int) time.Time {
	rDate := maturityDate
	diffYears := settlementDate.Year() - maturityDate.Year()
	rDate = rDate.AddDate(diffYears, 0, 0)
	if settlementDate.After(rDate) {
		rDate = rDate.AddDate(1, 0, 0)
	}
	monthsToAdd := -12 / freq
	for rDate.After(settlementDate) {
		rDate = rDate.AddDate(0, monthsToAdd, 0)
	}
	return rDate
}

// coupnum gets count of coupon dates.
func coupnum(settlementDate, maturityDate time.Time, freq, basis int) (float64, Result) {
	if maturityDate.After(settlementDate) {
		aDate := couppcd(settlementDate, maturityDate, freq, basis)
		months := (maturityDate.Year() - aDate.Year()) * 12 + int(maturityDate.Month()) - int(aDate.Month())
		return float64(months * freq) / 12.0, MakeEmptyResult()
	}
	return 0, MakeErrorResultType(ErrorTypeNum, "Settlement date should be before maturity date")
}

// getDuration returns the Macauley duration for an assumed par value of $100. It is defined as the weighted average of the present value of cash flows, and is used as a measure of a bond price's response to changes in yield.
func getDuration(settlementDate, maturityDate, coup, yield, freq float64, basis int) Result {
	fracResult := yearFrac(settlementDate, maturityDate, basis)
	if fracResult.Type == ResultTypeError {
		return fracResult
	}
	frac := fracResult.ValueNumber
	coups, err := coupnum(dateFromDays(settlementDate), dateFromDays(maturityDate), int(freq), basis)
	if err.Type == ResultTypeError {
		return err
	}
	duration := 0.0
	p := 0.0
	coup *= 100 / freq
	yield /= freq
	yield++
	diff := frac * freq - coups
	for t := 1.0; t < coups; t++ {
		tDiff := t + diff
		add := coup / math.Pow(yield, tDiff)
		p += add
		duration += tDiff * add
	}

	add := (coup + 100) / math.Pow(yield, coups + diff)

	p += add
	duration += (coups + diff) * add

	duration /= p
	duration /= freq

	return MakeNumberResult(duration)
}

type durationArgs struct {
	settlementDate float64
	maturityDate float64
	coupon float64
	yield float64
	freq float64
	basis int
}

// validateDurationData returns settlement date, maturity date, coupon rate, yield rate, frequency of payments, day count basis and error result by parsing incoming arguments
func parseDurationData(args []Result, funcName string) (*durationArgs, Result) {
	if len(args) != 5 && len(args) != 6 {
		return nil, MakeErrorResult(funcName + " requires five or six arguments")
	}
	var settlementDate, maturityDate float64
	settlementResult := args[0]
	switch settlementResult.Type {
	case ResultTypeNumber:
		settlementDate = float64(int(settlementResult.ValueNumber))
	case ResultTypeString:
		settlementFromString := DateValue([]Result{settlementResult})
		if settlementFromString.Type == ResultTypeError {
			return nil, MakeErrorResult("Incorrect settltment date for " + funcName)
		}
		settlementDate = settlementFromString.ValueNumber
	default:
		return nil, MakeErrorResult("Incorrect argument for " + funcName)
	}
	maturityResult := args[1]
	switch maturityResult.Type {
	case ResultTypeNumber:
		maturityDate = float64(int(maturityResult.ValueNumber))
	case ResultTypeString:
		maturityFromString := DateValue([]Result{maturityResult})
		if maturityFromString.Type == ResultTypeError {
			return nil, MakeErrorResult("Incorrect settltment date for " + funcName)
		}
		maturityDate = maturityFromString.ValueNumber
	default:
		return nil, MakeErrorResult("Incorrect argument for " + funcName)
	}
	if settlementDate >= maturityDate {
		return nil, MakeErrorResultType(ErrorTypeNum, "Settlement date should be before maturity date")
	}
	couponResult := args[2]
	if couponResult.Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires third argument of type number")
	}
	coupon := couponResult.ValueNumber
	if coupon < 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, "Coupon rate should not be negative")
	}
	yieldResult := args[3]
	if yieldResult.Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires fourth argument of type number")
	}
	yield := yieldResult.ValueNumber
	if yield < 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, "Yield rate should not be negative")
	}
	freqResult := args[4]
	if freqResult.Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires fifth argument of type number")
	}
	freq := float64(int(freqResult.ValueNumber))
	if !checkFreq(freq) {
		return nil, MakeErrorResultType(ErrorTypeNum, "Incorrect frequence value")
	}
	basis := 0
	if len(args) == 6 {
		basisResult := args[5]
		if basisResult.Type != ResultTypeNumber {
			return nil, MakeErrorResult(funcName + " requires sixth argument of type number")
		}
		basis = int(basisResult.ValueNumber)
		if !checkBasis(basis) {
			return nil, MakeErrorResultType(ErrorTypeNum, "Incorrect basis value")
		}
	}
	return &durationArgs{
		settlementDate,
		maturityDate,
		coupon,
		yield,
		freq,
		basis,
	}, MakeEmptyResult()
}

func checkFreq(freq float64) bool {
	return freq == 1 || freq == 2 || freq == 4
}

func checkBasis(basis int) bool {
	return basis >= 0 && basis <= 4
}

// Accrintm implements the Excel ACCRINTM function.
func Accrintm(args []Result) Result {
	argsNum := len(args)
	if argsNum != 4 && argsNum != 5 {
		return MakeErrorResult("ACCRINTM requires four or five arguments")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("ACCRINTM requires issue date to be number argument")
	}
	issue := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("ACCRINTM requires settlement date to be number argument")
	}
	settlement := args[1].ValueNumber
	if issue >= settlement {
		return MakeErrorResultType(ErrorTypeNum, "ACCRINTM requires settlement date to be later than issue date")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("ACCRINTM requires rate to be number argument")
	}
	rate := args[2].ValueNumber
	if rate <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "ACCRINTM requires rate to be positive number argument")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("ACCRINTM requires par to be number argument")
	}
	par := args[3].ValueNumber
	if par <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "ACCRINTM requires par to be positive number argument")
	}
	basis := 0
	if argsNum == 5 {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("ACCRINTM requires basis to be number argument")
		}
		basis = int(args[4].ValueNumber)
		if !checkBasis(basis) {
			return MakeErrorResultType(ErrorTypeNum, "Incorrect basis argument for ACCRINTM")
		}
	}
	fracResult := yearFrac(issue, settlement, basis)
	if fracResult.Type == ResultTypeError {
		return fracResult
	}
	return MakeNumberResult(par * rate * fracResult.ValueNumber)
}

// Amordegrc implements the Excel AMORDEGRC function.
func Amordegrc(args []Result) Result {
	parsedArgs, err := parseAmorArgs(args, "AMORDEGRC")
	if err.Type == ResultTypeError {
		return err
	}
	cost := parsedArgs.cost
	datePurchased := parsedArgs.datePurchased
	firstPeriod := parsedArgs.firstPeriod
	salvage := parsedArgs.salvage
	period := parsedArgs.period
	rate := parsedArgs.rate
	if rate >= 0.5 {
		return MakeErrorResultType(ErrorTypeNum, "AMORDEGRC requires rate to be less than 0.5")
	}
	basis := parsedArgs.basis

	lifeOfAssets := 1.0 / rate
	amorCoeff := 2.5
	if lifeOfAssets < 3 {
		amorCoeff = 1
	} else if lifeOfAssets < 5 {
		amorCoeff = 1.5
	} else if lifeOfAssets <= 6 {
		amorCoeff = 2
	}

	rate *= amorCoeff
	yfResult := yearFrac(datePurchased, firstPeriod, basis)
	if yfResult.Type == ResultTypeError {
		return MakeErrorResult("incorrect dates for AMORDEGRC")
	}
	nRate := mathRound(yfResult.ValueNumber * rate * cost)
	cost -= nRate
	rest := cost - salvage

	for n := 0; n < period; n++ {
		nRate = mathRound(rate * cost)
		rest -= nRate
		if rest < 0 {
			switch period - n {
			case 0:
			case 1:
				return MakeNumberResult(mathRound(cost * 0.5))
			default:
				return MakeNumberResult(0)
			}
		}
		cost -= nRate
	}

	return MakeNumberResult(nRate)
}

// Amorlinc implements the Excel AMORLINC function.
func Amorlinc(args []Result) Result {
	parsedArgs, err := parseAmorArgs(args, "AMORLINC")
	if err.Type == ResultTypeError {
		return err
	}
	cost := parsedArgs.cost
	datePurchased := parsedArgs.datePurchased
	firstPeriod := parsedArgs.firstPeriod
	salvage := parsedArgs.salvage
	period := parsedArgs.period
	rate := parsedArgs.rate
	basis := parsedArgs.basis

	yfResult := yearFrac(datePurchased, firstPeriod, basis)
	if yfResult.Type == ResultTypeError {
		return MakeErrorResult("incorrect dates for AMORLINC")
	}
	r0 := yfResult.ValueNumber * rate * cost
	if period == 0 {
		return MakeNumberResult(r0)
	}

	oneRate := cost * rate
	costDelta := cost - salvage
	numOfFullPeriods := int((costDelta - r0) / oneRate)

	if period <= numOfFullPeriods {
		return MakeNumberResult(oneRate)
	} else if period == numOfFullPeriods + 1 {
		return MakeNumberResult(costDelta - oneRate * float64(numOfFullPeriods) - r0)
	} else {
		return MakeNumberResult(0)
	}
}

type amorArgs struct {
	cost float64
	datePurchased float64
	firstPeriod float64
	salvage float64
	period int
	rate float64
	basis int
}

func parseAmorArgs(args []Result, funcName string) (*amorArgs, Result) {
	argsNum := len(args)
	if argsNum != 6 && argsNum != 7 {
		return nil, MakeErrorResult(funcName + " requires six or seven arguments")
	}
	if args[0].Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires cost to be number argument")
	}
	cost := args[0].ValueNumber
	if cost < 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName + " requires cost to be positive")
	}
	if args[1].Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires date purchased to be number argument")
	}
	datePurchased := args[1].ValueNumber
	if datePurchased < 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName + " requires date purchased to be positive")
	}
	if args[2].Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires first period to be number argument")
	}
	firstPeriod := args[2].ValueNumber
	if firstPeriod < datePurchased {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName + " requires first period to be later than date purchased")
	}
	if args[3].Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires salvage to be number argument")
	}
	salvage := args[3].ValueNumber
	if salvage < 0 || salvage > cost {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName + " requires salvage to be between 0 and the initial cost")
	}
	if args[4].Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires period to be number argument")
	}
	period := int(args[4].ValueNumber)
	if period < 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName + " requires period to be non-negative")
	}
	if args[5].Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires depreciation rate to be number argument")
	}
	rate := args[5].ValueNumber
	if rate < 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName + " requires rate to be more than 0 and less than 0.5")
	}
	basis := 0
	if argsNum == 7 {
		if args[6].Type != ResultTypeNumber {
			return nil, MakeErrorResult(funcName + " requires basis to be number argument")
		}
		basis = int(args[6].ValueNumber)
		if !checkBasis(basis) || basis == 2 {
			return nil, MakeErrorResultType(ErrorTypeNum, "Incorrect basis argument for " + funcName)
		}
	}
	return &amorArgs{
		cost,
		datePurchased,
		firstPeriod,
		salvage,
		period,
		rate,
		basis,
	}, MakeEmptyResult()
}

func mathRound(x float64) float64 {
	return float64(int(x + 0.5))
}

type cumulArgs struct {
	rate float64
	nPer float64
	pv float64
	startPeriod float64
	endPeriod float64
	t int
}

// Cumipmt implements the Excel CUMIPMT function.
func Cumipmt(args []Result) Result {
	parsedArgs, err := parseCumulArgs(args, "CUMIPMT")
	if err.Type == ResultTypeError {
		return err
	}
	rate := parsedArgs.rate
	nPer := parsedArgs.nPer
	pv := parsedArgs.pv
	startPeriod := parsedArgs.startPeriod
	endPeriod := parsedArgs.endPeriod
	t := parsedArgs.t

	payment := pmt(rate, nPer, pv, 0, t)
	interest := 0.0
	if startPeriod == 1 {
		if t == 0 {
			interest = -pv
			startPeriod++
		}
	}
	for i := startPeriod; i <= endPeriod; i++ {
		if t == 1 {
			interest += fv(rate, i - 2, payment, pv, 1) - payment
		} else {
			interest += fv(rate, i - 1, payment, pv, 0)
		}
	}
	interest *= rate
	return MakeNumberResult(interest)
}

// Cumprinc implements the Excel CUMPRINC function.
func Cumprinc(args []Result) Result {
	parsedArgs, err := parseCumulArgs(args, "CUMPRINC")
	if err.Type == ResultTypeError {
		return err
	}
	rate := parsedArgs.rate
	nPer := parsedArgs.nPer
	pv := parsedArgs.pv
	startPeriod := parsedArgs.startPeriod
	endPeriod := parsedArgs.endPeriod
	t := parsedArgs.t

	payment := pmt(rate, nPer, pv, 0, t)
	principal := 0.0
	if startPeriod == 1 {
		if t == 0 {
			principal = payment + pv * rate
		} else {
			principal = payment
		}
		startPeriod++
	}
	for i := startPeriod; i <= endPeriod; i++ {
		if t == 1 {
			principal += payment - (fv(rate, i - 2, payment, pv, 1) - payment) * rate
		} else {
			principal += payment - fv(rate, i - 1, payment, pv, 0) * rate
		}
	}
	return MakeNumberResult(principal)
}

func parseCumulArgs(args []Result, funcName string) (*cumulArgs, Result) {
	if len(args) != 6 {
		return nil, MakeErrorResult(funcName + " requires six arguments")
	}
	if args[0].Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires rate to be number argument")
	}
	rate := args[0].ValueNumber
	if rate <= 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName + " requires rate to be positive number argument")
	}
	if args[1].Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires number of periods to be number argument")
	}
	nPer := args[1].ValueNumber
	if nPer <= 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName + " requires number of periods to be positive number argument")
	}
	if args[2].Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires present value to be number argument")
	}
	pv := args[2].ValueNumber
	if pv <= 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName + " requires present value to be positive number argument")
	}
	if args[3].Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires start period to be number argument")
	}
	startPeriod := args[3].ValueNumber
	if startPeriod <= 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName + " requires start period to be positive number argument")
	}
	if args[4].Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires end period to be number argument")
	}
	endPeriod := args[4].ValueNumber
	if endPeriod <= 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName + " requires end period to be positive number argument")
	}
	if endPeriod < startPeriod {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName + " requires end period to be later or equal to start period")
	}
	if endPeriod > nPer {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName + " requires periods to be in number of periods range")
	}
	t := int(args[5].ValueNumber)
	if t != 0 && t != 1 {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName + " requires type to be 0 or 1")
	}
	return &cumulArgs{
		rate,
		nPer,
		pv,
		startPeriod,
		endPeriod,
		t,
	}, MakeEmptyResult()
}

func pmt(rate, periods, present, future float64, t int ) float64 {
	var result float64
	if rate == 0 {
		result = (present + future) / periods
	} else {
		term := math.Pow(1 + rate, periods)
		if t == 1 {
			result = (future * rate / (term - 1) + present * rate / (1 - 1 / term)) / (1 + rate)
		} else {
			result = future * rate / (term - 1) + present * rate / (1 - 1 / term)
		}
	}
	return -result
}

func fv(rate, periods, payment, value float64, t int) float64 {
	var result float64
	if rate == 0 {
		result = value + payment * periods
	} else {
		term := math.Pow(1 + rate, periods)
		if t == 1 {
			result = value * term + payment * (1 + rate) * (term - 1) / rate
		} else {
			result = value * term + payment * (term - 1) / rate
		}
	}
	return -result
}

// Db implements the Excel DB function.
func Db(args []Result) Result {
	argsNum := len(args)
	if argsNum != 4 && argsNum != 5 {
		return MakeErrorResult("DB requires four or five number arguments")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("DB requires cost to be number argument")
	}
	cost := args[0].ValueNumber
	if cost < 0 {
		return MakeErrorResultType(ErrorTypeNum, "DB requires cost to be non negative")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("DB requires salvage to be number argument")
	}
	salvage := args[1].ValueNumber
	if salvage < 0 {
		return MakeErrorResultType(ErrorTypeNum, "DB requires salvage to be non negative")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("DB requires life to be number argument")
	}
	life := args[2].ValueNumber
	if life <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "DB requires life to be positive")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("DB requires period to be number argument")
	}
	period := args[3].ValueNumber
	if period <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "DB requires period to be positive")
	}
	if period - life > 1 {
		return MakeErrorResultType(ErrorTypeNum, "Incorrect period for DB")
	}
	month := 12.0
	if argsNum == 5 {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("DB requires month to be number argument")
		}
		month = args[4].ValueNumber
		if month < 1 || month > 12 {
			return MakeErrorResultType(ErrorTypeNum, "DB requires month to be in range of 1 and 12")
		}
	}
	if month == 12 && period > life {
		return MakeErrorResultType(ErrorTypeNum, "Incorrect period for DB")
	}
	if salvage >= cost {
		return MakeNumberResult(0)
	}
	rate := 1 - math.Pow(salvage / cost, 1 / life)
	rate = float64(int(rate * 1000 + 0.5)) / 1000 // round to 3 decimal places
	initial := cost * rate * month / 12
	if period == 1 {
		return MakeNumberResult(initial)
	}
	total := initial
	current := 0.0
	ceiling := life
	if ceiling > period {
		ceiling = period
	}
	for i := 2.0; i <= ceiling; i++ {
		current = (cost - total) * rate
		total += current
	}
	if period > life {
		return MakeNumberResult((cost - total) * rate * (12 - month) / 12)
	}
	return MakeNumberResult(current)
}

// Ddb implements the Excel DDB function.
func Ddb(args []Result) Result {
	argsNum := len(args)
	if argsNum != 4 && argsNum != 5 {
		return MakeErrorResult("DDB requires four or five number arguments")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("DDB requires cost to be number argument")
	}
	cost := args[0].ValueNumber
	if cost < 0 {
		return MakeErrorResultType(ErrorTypeNum, "DDB requires cost to be non negative")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("DDB requires salvage to be number argument")
	}
	salvage := args[1].ValueNumber
	if salvage < 0 {
		return MakeErrorResultType(ErrorTypeNum, "DDB requires salvage to be non negative")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("DDB requires life to be number argument")
	}
	life := args[2].ValueNumber
	if life <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "DDB requires life to be positive")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("DDB requires period to be number argument")
	}
	period := args[3].ValueNumber
	if period < 1 {
		return MakeErrorResultType(ErrorTypeNum, "DDB requires period to be positive")
	}
	if period > life {
		return MakeErrorResultType(ErrorTypeNum, "Incorrect period for DDB")
	}
	factor := 2.0
	if argsNum == 5 {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("DDB requires factor to be number argument")
		}
		factor = args[4].ValueNumber
		if factor < 0 {
			return MakeErrorResultType(ErrorTypeNum, "DDB requires factor to be non negative")
		}
	}

	oldValue := 0.0
	rate := factor / life
	if rate >= 1 {
		rate = 1
		if period == 1 {
			oldValue = cost
		}
	} else {
		oldValue = cost * math.Pow(1 - rate, period - 1)
	}
	newValue := cost * math.Pow(1 - rate, period)

	var ddb float64

	if newValue < salvage {
		ddb = oldValue - salvage
	} else {
		ddb = oldValue - newValue
	}
	if ddb < 0 {
		ddb = 0
	}
	return MakeNumberResult(ddb)
}

// Disc implements the Excel DISC function.
func Disc(args []Result) Result {
	argsNum := len(args)
	if argsNum != 4 && argsNum != 5 {
		return MakeErrorResult("DISC requires four or five arguments")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("DISC requires settlement date to be number argument")
	}
	settlement := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("DISC requires maturity date to be number argument")
	}
	maturity := args[1].ValueNumber
	if settlement >= maturity {
		return MakeErrorResultType(ErrorTypeNum, "DISC requires maturity date to be later than settlement date")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("DISC requires pr to be number argument")
	}
	pr := args[2].ValueNumber
	if pr <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "DISC requires pr to be positive number argument")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("DISC requires redemption to be number argument")
	}
	redemption := args[3].ValueNumber
	if redemption <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "DISC requires redemption to be positive number argument")
	}
	basis := 0
	if argsNum == 5 {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("DISC requires basis to be number argument")
		}
		basis = int(args[4].ValueNumber)
		if !checkBasis(basis) {
			return MakeErrorResultType(ErrorTypeNum, "Incorrect basis argument for DISC")
		}
	}
	fracResult := yearFrac(settlement, maturity, basis)
	if fracResult.Type == ResultTypeError {
		return fracResult
	}
	return MakeNumberResult((redemption - pr) / redemption / fracResult.ValueNumber)
}

// Dollarde implements the Excel DOLLARDE function.
func Dollarde(args []Result) Result {
	dollar, fraction, resultErr := parseDollarArgs(args, "DOLLARDE")
	if resultErr.Type == ResultTypeError {
		return resultErr
	}
	if fraction < 1 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "DOLLARDE requires fraction to be equal or more than 1")
	}
	if dollar == 0 {
		return MakeNumberResult(0)
	}
	neg := dollar < 0
	if neg {
		dollar = -dollar
	}
	dollarStr := args[0].Value()
	split := strings.Split(dollarStr, ".")
	dollarInt := float64(int(dollar))
	dollarFracStr := split[1]
	dollarFracOrder := len(dollarFracStr)
	fractionOrder := int(math.Log10(fraction)) + 1
	power := float64(fractionOrder - dollarFracOrder)
	dollarFrac, err := strconv.ParseFloat(dollarFracStr, 64)
	if err != nil {
		return MakeErrorResult("Incorrect fraction argument for DOLLARDE")
	}
	dollarFrac *= math.Pow(10, power)
	dollarde := dollarInt + dollarFrac / fraction
	if neg {
		dollarde = -dollarde
	}
	return MakeNumberResult(dollarde)
}

// Dollarfr implements the Excel DOLLARFR function.
func Dollarfr(args []Result) Result {
	dollar, fraction, resultErr := parseDollarArgs(args, "DOLLARFR")
	if resultErr.Type == ResultTypeError {
		return resultErr
	}
	if dollar == 0 {
		return MakeNumberResult(0)
	}
	neg := dollar < 0
	if neg {
		dollar = -dollar
	}
	dollarInt := float64(int(dollar))
	dollarStr := args[0].Value()
	split := strings.Split(dollarStr, ".")
	dollarFracStr := split[1]
	dollarFrac, err := strconv.ParseFloat(dollarFracStr, 64)
	if err != nil {
		return MakeErrorResult("Incorrect fraction argument for DOLLARFR")
	}
	dollarFracOrder := float64(len(dollarFracStr))
	dollarFrac /= math.Pow(10, dollarFracOrder)

	dollarfr := dollarFrac * fraction / math.Pow(10, float64(int(math.Log10(fraction))) + 1) + dollarInt
	if neg {
		dollarfr = -dollarfr
	}
	return MakeNumberResult(dollarfr)
}

func parseDollarArgs(args []Result, funcName string) (float64, float64, Result) {
	if len(args) != 2 {
		return 0, 0, MakeErrorResult(funcName + " requires two arguments")
	}
	if args[0].Type != ResultTypeNumber {
		return 0, 0, MakeErrorResult(funcName + " requires fractional dollar to be number argument")
	}
	dollar := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return 0, 0, MakeErrorResult(funcName + " requires fraction to be number argument")
	}
	fraction := float64(int(args[1].ValueNumber))
	if fraction < 0 {
		return 0, 0, MakeErrorResultType(ErrorTypeNum, funcName + " requires fraction to be positive number")
	}
	return dollar, fraction, MakeEmptyResult()
}

// Effect implements the Excel EFFECT function.
func Effect(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("EFFECT requires two arguments")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("EFFECT requires nominal interest rate to be number argument")
	}
	nominal := args[0].ValueNumber
	if nominal <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "EFFECT requires nominal interest rate to be positive number argument")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("EFFECT requires number of compounding periods to be number argument")
	}
	npery := float64(int(args[1].ValueNumber))
	if npery < 1 {
		return MakeErrorResult("EFFECT requires number of compounding periods to be 1 or more")
	}
	return MakeNumberResult(math.Pow((1 + nominal / npery), npery) - 1)
}

// Fv implements the Excel FV function.
func Fv(args []Result) Result {
	argsNum := len(args)
	if argsNum < 3 || argsNum > 5 {
		return MakeErrorResult("FV requires number of arguments in range of 3 and 5")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("FV requires rate to be number argument")
	}
	rate := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("FV requires number of periods to be number argument")
	}
	nPer := args[1].ValueNumber
	if nPer != float64(int(nPer)) {
		return MakeErrorResultType(ErrorTypeNum, "FV requires number of periods to be integer number argument")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("FV requires payment to be number argument")
	}
	pmt := args[2].ValueNumber
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("FV requires payment to be number argument")
	}
	pv := 0.0
	if argsNum >= 4 {
		if args[3].Type != ResultTypeNumber {
			return MakeErrorResult("FV requires present value to be number argument")
		}
		pv = args[3].ValueNumber
	}
	t := 0
	if argsNum == 5 {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("FV requires type to be number argument")
		}
		t = int(args[4].ValueNumber)
		if t != 0 {
			t = 1
		}
	}
	return MakeNumberResult(fv(rate, nPer, pmt, pv, t))
}

// Fvschedule implements the Excel FVSCHEDULE function.
func Fvschedule(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("FVSCHEDULE requires two arguments")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("FVSCHEDULE requires principal to be number argument")
	}
	principal := args[0].ValueNumber
	switch args[1].Type {
	case ResultTypeNumber:
		return MakeNumberResult(principal * (args[1].ValueNumber+1))
	case ResultTypeList, ResultTypeArray:
		schedule := arrayFromRange(args[1])
		for _, row := range schedule {
			for _, rate := range row {
				if rate.Type != ResultTypeNumber || rate.IsBoolean {
					return MakeErrorResult("FVSCHEDULE requires rates to be numbers")
				}
				principal *= 1.0 + rate.ValueNumber
			}
		}
		return MakeNumberResult(principal)
	default:
		return MakeErrorResult("FVSCHEDULE requires schedule to be of array type")
	}
}

// Intrate implements the Excel INTRATE function.
func Intrate(args []Result) Result {
	argsNum := len(args)
	if argsNum != 4 && argsNum != 5 {
		return MakeErrorResult("INTRATE requires four or five arguments")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("INTRATE requires settlement date to be number argument")
	}
	settlement := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("INTRATE requires maturity date to be number argument")
	}
	maturity := args[1].ValueNumber
	if settlement >= maturity {
		return MakeErrorResultType(ErrorTypeNum, "INTRATE requires maturity date to be later than settlement date")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("INTRATE requires investment to be number argument")
	}
	investment := args[2].ValueNumber
	if investment <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "INTRATE requires investment to be positive number argument")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("INTRATE requires redemption to be number argument")
	}
	redemption := args[3].ValueNumber
	if redemption <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "INTRATE requires redemption to be positive number argument")
	}
	basis := 0
	if argsNum == 5 {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("INTRATE requires basis to be number argument")
		}
		basis = int(args[4].ValueNumber)
		if !checkBasis(basis) {
			return MakeErrorResultType(ErrorTypeNum, "Incorrect basis argument for INTRATE")
		}
	}
	fracResult := yearFrac(settlement, maturity, basis)
	if fracResult.Type == ResultTypeError {
		return fracResult
	}
	return MakeNumberResult((redemption - investment) / investment / fracResult.ValueNumber)
}

// Ipmt implements the Excel IPMT function.
func Ipmt(args []Result) Result {
	argsNum := len(args)
	if argsNum < 4 || argsNum > 6 {
		return MakeErrorResult("IPMT requires six arguments")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("IPMT requires rate to be number argument")
	}
	rate := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("IPMT requires period to be number argument")
	}
	period := args[1].ValueNumber
	if period <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "IPMT requires period to be positive number argument")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("IPMT requires number of periods to be number argument")
	}
	nPer := args[2].ValueNumber
	if nPer <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "IPMT requires number of periods to be positive number argument")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("IPMT requires present value to be number argument")
	}
	presentValue := args[3].ValueNumber
	futureValue := 0.0
	if argsNum > 4 {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("IPMT requires future value to be number argument")
		}
		futureValue = args[4].ValueNumber
	}
	t := 0
	if argsNum == 6 {
		if args[5].Type != ResultTypeNumber {
			return MakeErrorResult("IPMT requires start period to be number argument")
		}
		t = int(args[5].ValueNumber)
		if t != 0  {
			t = 1
		}
	}
	payment := pmt(rate, nPer, presentValue, futureValue, t)
	var interest float64
	if period == 1 {
		if t == 1 {
			interest = 0
		} else {
			interest = -presentValue
		}
	} else {
		if t == 1 {
			interest = fv(rate, period - 2, payment, presentValue, 1) - payment
		} else {
			interest = fv(rate, period - 1, payment, presentValue, 0)
		}
	}

	return MakeNumberResult(interest * rate)
}
