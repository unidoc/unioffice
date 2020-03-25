// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import (
	"math"
	"strconv"
	"strings"
	"time"
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
	RegisterFunction("IRR", Irr)
	RegisterFunction("ISPMT", Ispmt)
	RegisterFunction("MDURATION", Mduration)
	RegisterFunction("MIRR", Mirr)
	RegisterFunction("NOMINAL", Nominal)
	RegisterFunction("NPER", Nper)
	RegisterFunction("NPV", Npv)
	RegisterFunction("ODDLPRICE", Oddlprice)
	RegisterFunction("ODDLYIELD", Oddlyield)
	RegisterFunction("PDURATION", Pduration)
	RegisterFunction("_xlfn.PDURATION", Pduration)
	RegisterFunction("PMT", Pmt)
	RegisterFunction("PPMT", Ppmt)
	RegisterFunction("PRICE", Price)
	RegisterFunction("PRICEDISC", Pricedisc)
	RegisterFunction("PRICEMAT", Pricemat)
	RegisterFunction("PV", Pv)
	RegisterFunction("RATE", Rate)
	RegisterFunction("RECEIVED", Received)
	RegisterFunction("RRI", Rri)
	RegisterFunction("_xlfn.RRI", Rri)
	RegisterFunction("SLN", Sln)
	RegisterFunction("SYD", Syd)
	RegisterFunction("TBILLEQ", Tbilleq)
	RegisterFunction("TBILLPRICE", Tbillprice)
	RegisterFunction("TBILLYIELD", Tbillyield)
	RegisterFunction("VDB", Vdb)
	RegisterFunction("XIRR", Xirr)
	RegisterFunction("XNPV", Xnpv)
	RegisterFunction("YIELD", Yield)
	RegisterFunction("YIELDDISC", Yielddisc)
	RegisterFunction("YIELDMAT", Yieldmat)
}

func getSettlementMaturity(settlementResult, maturityResult Result, funcName string) (float64, float64, Result) {
	settlementDate, errResult := parseDate(settlementResult, "settlement date", funcName)
	if errResult.Type == ResultTypeError {
		return 0, 0, errResult
	}
	maturityDate, errResult := parseDate(maturityResult, "maturity date", funcName)
	if errResult.Type == ResultTypeError {
		return 0, 0, errResult
	}
	if settlementDate >= maturityDate {
		return 0, 0, MakeErrorResultType(ErrorTypeNum, funcName+" requires maturity date to be later than settlement date")
	}
	return settlementDate, maturityDate, empty
}

type couponArgs struct {
	settlementDate float64
	maturityDate   float64
	freq           int
	basis          int
}

// Coupdaybs implements the Excel COUPDAYBS function.
func Coupdaybs(args []Result) Result {
	parsedArgs, err := parseCouponArgs(args, "COUPDAYBS")
	if err.Type == ResultTypeError {
		return err
	}
	return MakeNumberResult(coupdaybs(parsedArgs.settlementDate, parsedArgs.maturityDate, parsedArgs.freq, parsedArgs.basis))
}

// coupdaybs returns the number of days from the beginning of the coupon period to the settlement date.
func coupdaybs(settlementDateF, maturityDateF float64, freq, basis int) float64 {
	settlementDate := dateFromDays(settlementDateF)
	maturityDate := dateFromDays(maturityDateF)
	pcd := couppcd(settlementDate, maturityDate, freq, basis)
	return getDiff(pcd, settlementDate, basis)
}

// Coupdays implements the Excel COUPDAYS function.
func Coupdays(args []Result) Result {
	parsedArgs, err := parseCouponArgs(args, "COUPDAYS")
	if err.Type == ResultTypeError {
		return err
	}
	return MakeNumberResult(coupdays(parsedArgs.settlementDate, parsedArgs.maturityDate, parsedArgs.freq, parsedArgs.basis))
}

// coupdays returns the number of days in the coupon period that contains the settlement date.
func coupdays(settlementDateF, maturityDateF float64, freq, basis int) float64 {
	settlementDate := dateFromDays(settlementDateF)
	maturityDate := dateFromDays(maturityDateF)
	if basis == 1 {
		pcd := couppcd(settlementDate, maturityDate, freq, 1)
		next := pcd.AddDate(0, 12/freq, 0)
		return getDiff(pcd, next, basis)
	}
	return float64(getDaysInYear(0, basis)) / float64(freq)
}

// Coupdaysnc implements the Excel COUPDAYSNC function.
func Coupdaysnc(args []Result) Result {
	parsedArgs, err := parseCouponArgs(args, "COUPDAYSNC")
	if err.Type == ResultTypeError {
		return err
	}
	return MakeNumberResult(coupdaysnc(parsedArgs.settlementDate, parsedArgs.maturityDate, parsedArgs.freq, parsedArgs.basis))
}

// coupdaysnc returns the number of days from the settlement date to the next coupon date.
func coupdaysnc(settlementDateF, maturityDateF float64, freq, basis int) float64 {
	settlementDate := dateFromDays(settlementDateF)
	maturityDate := dateFromDays(maturityDateF)
	ncd := coupncd(settlementDate, maturityDate, freq)
	return getDiff(settlementDate, ncd, basis)
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
	freq := parsedArgs.freq
	basis := parsedArgs.basis
	cn, err := coupnum(parsedArgs.settlementDate, parsedArgs.maturityDate, freq, basis)
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

// coupncd finds next coupon date after settlement.
func coupncd(settlementDate, maturityDate time.Time, freq int) time.Time {
	ncd := time.Date(settlementDate.Year(), maturityDate.Month(), maturityDate.Day(), 0, 0, 0, 0, time.UTC)
	if ncd.After(settlementDate) {
		ncd = ncd.AddDate(-1, 0, 0)
	}
	for !ncd.After(settlementDate) {
		ncd = ncd.AddDate(0, 12/freq, 0)
	}
	return ncd
}

func parseCouponArgs(args []Result, funcName string) (*couponArgs, Result) {
	argsNum := len(args)
	if argsNum != 3 && argsNum != 4 {
		return nil, MakeErrorResult(funcName + " requires three or four arguments")
	}
	settlementDate, maturityDate, errResult := getSettlementMaturity(args[0], args[1], funcName)
	if errResult.Type == ResultTypeError {
		return nil, errResult
	}
	if args[2].Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires frequency to be number argument")
	}
	freq := args[2].ValueNumber
	if !checkFreq(freq) {
		return nil, MakeErrorResult("Incorrect frequency for " + funcName)
	}
	basis := 0
	if argsNum == 4 && args[3].Type != ResultTypeEmpty {
		if args[3].Type != ResultTypeNumber {
			return nil, MakeErrorResult(funcName + " requires basis to be number argument")
		}
		basis = int(args[3].ValueNumber)
		if !checkBasis(basis) {
			return nil, MakeErrorResultType(ErrorTypeNum, "Incorrect basis argument for "+funcName)
		}
	}
	return &couponArgs{
		settlementDate,
		maturityDate,
		int(freq),
		basis,
	}, empty
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
func coupnum(settlementDateF, maturityDateF float64, freq, basis int) (float64, Result) {
	settlementDate, maturityDate := dateFromDays(settlementDateF), dateFromDays(maturityDateF)
	if maturityDate.After(settlementDate) {
		aDate := couppcd(settlementDate, maturityDate, freq, basis)
		months := (maturityDate.Year()-aDate.Year())*12 + int(maturityDate.Month()) - int(aDate.Month())
		return float64(months*freq) / 12.0, empty
	}
	return 0, MakeErrorResultType(ErrorTypeNum, "Settlement date should be before maturity date")
}

// getDuration returns the Macauley duration for an assumed par value of $100. It is defined as the weighted average of the present value of cash flows, and is used as a measure of a bond price's response to changes in yield.
func getDuration(settlementDate, maturityDate, coup, yield, freq float64, basis int) Result {
	frac, errResult := yearFrac(settlementDate, maturityDate, basis)
	if errResult.Type == ResultTypeError {
		return errResult
	}
	coups, err := coupnum(settlementDate, maturityDate, int(freq), basis)
	if err.Type == ResultTypeError {
		return err
	}
	duration := 0.0
	p := 0.0
	coup *= 100 / freq
	yield /= freq
	yield++
	diff := frac*freq - coups
	for t := 1.0; t < coups; t++ {
		tDiff := t + diff
		add := coup / math.Pow(yield, tDiff)
		p += add
		duration += tDiff * add
	}

	add := (coup + 100) / math.Pow(yield, coups+diff)

	p += add
	duration += (coups + diff) * add

	duration /= p
	duration /= freq

	return MakeNumberResult(duration)
}

type durationArgs struct {
	settlementDate float64
	maturityDate   float64
	coupon         float64
	yield          float64
	freq           float64
	basis          int
}

// validateDurationData returns settlement date, maturity date, coupon rate, yield rate, frequency of payments, day count basis and error result by parsing incoming arguments
func parseDurationData(args []Result, funcName string) (*durationArgs, Result) {
	argsNum := len(args)
	if argsNum != 5 && argsNum != 6 {
		return nil, MakeErrorResult(funcName + " requires five or six arguments")
	}
	settlementDate, maturityDate, errResult := getSettlementMaturity(args[0], args[1], funcName)
	if errResult.Type == ResultTypeError {
		return nil, errResult
	}
	couponResult := args[2]
	if couponResult.Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires coupon rate of type number")
	}
	coupon := couponResult.ValueNumber
	if coupon < 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, "Coupon rate should not be negative")
	}
	yieldResult := args[3]
	if yieldResult.Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires yield rate of type number")
	}
	yield := yieldResult.ValueNumber
	if yield < 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, "Yield rate should not be negative")
	}
	freqResult := args[4]
	if freqResult.Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires frequency of type number")
	}
	freq := float64(int(freqResult.ValueNumber))
	if !checkFreq(freq) {
		return nil, MakeErrorResultType(ErrorTypeNum, "Incorrect frequence value")
	}
	basis := 0
	if argsNum == 6 && args[5].Type != ResultTypeEmpty {
		basisResult := args[5]
		if basisResult.Type != ResultTypeNumber {
			return nil, MakeErrorResult(funcName + " requires basis of type number")
		}
		basis = int(basisResult.ValueNumber)
		if !checkBasis(basis) {
			return nil, MakeErrorResultType(ErrorTypeNum, "Incorrect basis value for "+funcName)
		}
	}
	return &durationArgs{
		settlementDate,
		maturityDate,
		coupon,
		yield,
		freq,
		basis,
	}, empty
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
	issueDate, errResult := parseDate(args[0], "issue date", "ACCRINTM")
	if errResult.Type == ResultTypeError {
		return errResult
	}
	settlementDate, errResult := parseDate(args[1], "settlement date", "ACCRINTM")
	if errResult.Type == ResultTypeError {
		return errResult
	}
	if issueDate >= settlementDate {
		return MakeErrorResultType(ErrorTypeNum, "Issue date should be earlier than settlement date")
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
	if argsNum == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("ACCRINTM requires basis to be number argument")
		}
		basis = int(args[4].ValueNumber)
		if !checkBasis(basis) {
			return MakeErrorResultType(ErrorTypeNum, "Incorrect basis argument for ACCRINTM")
		}
	}
	frac, errResult := yearFrac(issueDate, settlementDate, basis)
	if errResult.Type == ResultTypeError {
		return errResult
	}
	return MakeNumberResult(par * rate * frac)
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
	yf, errResult := yearFrac(datePurchased, firstPeriod, basis)
	if errResult.Type == ResultTypeError {
		return MakeErrorResult("incorrect dates for AMORDEGRC")
	}
	nRate := mathRound(yf * rate * cost)
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

	yf, errResult := yearFrac(datePurchased, firstPeriod, basis)
	if errResult.Type == ResultTypeError {
		return MakeErrorResult("incorrect dates for AMORLINC")
	}
	r0 := yf * rate * cost
	if period == 0 {
		return MakeNumberResult(r0)
	}

	oneRate := cost * rate
	costDelta := cost - salvage
	numOfFullPeriods := int((costDelta - r0) / oneRate)

	if period <= numOfFullPeriods {
		return MakeNumberResult(oneRate)
	} else if period == numOfFullPeriods+1 {
		return MakeNumberResult(costDelta - oneRate*float64(numOfFullPeriods) - r0)
	} else {
		return MakeNumberResult(0)
	}
}

type amorArgs struct {
	cost          float64
	datePurchased float64
	firstPeriod   float64
	salvage       float64
	period        int
	rate          float64
	basis         int
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
		return nil, MakeErrorResultType(ErrorTypeNum, funcName+" requires cost to be non negative")
	}
	datePurchased, errResult := parseDate(args[1], "date purchased", funcName)
	if errResult.Type == ResultTypeError {
		return nil, errResult
	}
	firstPeriod, errResult := parseDate(args[2], "first period", funcName)
	if errResult.Type == ResultTypeError {
		return nil, errResult
	}
	if firstPeriod < datePurchased {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName+" requires first period to be later than date purchased")
	}
	if args[3].Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires salvage to be number argument")
	}
	salvage := args[3].ValueNumber
	if salvage < 0 || salvage > cost {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName+" requires salvage to be between 0 and the initial cost")
	}
	if args[4].Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires period to be number argument")
	}
	period := int(args[4].ValueNumber)
	if period < 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName+" requires period to be non-negative")
	}
	if args[5].Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires depreciation rate to be number argument")
	}
	rate := args[5].ValueNumber
	if rate < 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName+" requires depreciation rate to be more than 0 and less than 0.5")
	}
	basis := 0
	if argsNum == 7 && args[6].Type != ResultTypeEmpty {
		if args[6].Type != ResultTypeNumber {
			return nil, MakeErrorResult(funcName + " requires basis to be number argument")
		}
		basis = int(args[6].ValueNumber)
		if !checkBasis(basis) || basis == 2 {
			return nil, MakeErrorResultType(ErrorTypeNum, "Incorrect basis argument for "+funcName)
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
	}, empty
}

func mathRound(x float64) float64 {
	return float64(int(x + 0.5))
}

type cumulArgs struct {
	rate         float64
	nPer         float64
	presentValue float64
	startPeriod  float64
	endPeriod    float64
	t            int
}

// Cumipmt implements the Excel CUMIPMT function.
func Cumipmt(args []Result) Result {
	parsedArgs, err := parseCumulArgs(args, "CUMIPMT")
	if err.Type == ResultTypeError {
		return err
	}
	rate := parsedArgs.rate
	nPer := parsedArgs.nPer
	presentValue := parsedArgs.presentValue
	startPeriod := parsedArgs.startPeriod
	endPeriod := parsedArgs.endPeriod
	t := parsedArgs.t

	payment := pmt(rate, nPer, presentValue, 0, t)
	interest := 0.0
	if startPeriod == 1 {
		if t == 0 {
			interest = -presentValue
			startPeriod++
		}
	}
	for i := startPeriod; i <= endPeriod; i++ {
		if t == 1 {
			interest += fv(rate, i-2, payment, presentValue, 1) - payment
		} else {
			interest += fv(rate, i-1, payment, presentValue, 0)
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
	presentValue := parsedArgs.presentValue
	startPeriod := parsedArgs.startPeriod
	endPeriod := parsedArgs.endPeriod
	t := parsedArgs.t

	payment := pmt(rate, nPer, presentValue, 0, t)
	principal := 0.0
	if startPeriod == 1 {
		if t == 0 {
			principal = payment + presentValue*rate
		} else {
			principal = payment
		}
		startPeriod++
	}
	for i := startPeriod; i <= endPeriod; i++ {
		if t == 1 {
			principal += payment - (fv(rate, i-2, payment, presentValue, 1)-payment)*rate
		} else {
			principal += payment - fv(rate, i-1, payment, presentValue, 0)*rate
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
		return nil, MakeErrorResultType(ErrorTypeNum, funcName+" requires rate to be positive number argument")
	}
	if args[1].Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires number of periods to be number argument")
	}
	nPer := args[1].ValueNumber
	if nPer <= 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName+" requires number of periods to be positive number argument")
	}
	if args[2].Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires present value to be number argument")
	}
	presentValue := args[2].ValueNumber
	if presentValue <= 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName+" requires present value to be positive number argument")
	}
	if args[3].Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires start period to be number argument")
	}
	startPeriod := args[3].ValueNumber
	if startPeriod <= 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName+" requires start period to be positive number argument")
	}
	if args[4].Type != ResultTypeNumber {
		return nil, MakeErrorResult(funcName + " requires end period to be number argument")
	}
	endPeriod := args[4].ValueNumber
	if endPeriod <= 0 {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName+" requires end period to be positive number argument")
	}
	if endPeriod < startPeriod {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName+" requires end period to be later or equal to start period")
	}
	if endPeriod > nPer {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName+" requires periods to be in periods range")
	}
	t := int(args[5].ValueNumber)
	if t != 0 && t != 1 {
		return nil, MakeErrorResultType(ErrorTypeNum, funcName+" requires type to be 0 or 1")
	}
	return &cumulArgs{
		rate,
		nPer,
		presentValue,
		startPeriod,
		endPeriod,
		t,
	}, empty
}

func pmt(rate, periods, present, future float64, t int) float64 {
	var result float64
	if rate == 0 {
		result = (present + future) / periods
	} else {
		term := math.Pow(1+rate, periods)
		if t == 1 {
			result = (future*rate/(term-1) + present*rate/(1-1/term)) / (1 + rate)
		} else {
			result = future*rate/(term-1) + present*rate/(1-1/term)
		}
	}
	return -result
}

func fv(rate, periods, payment, value float64, t int) float64 {
	var result float64
	if rate == 0 {
		result = value + payment*periods
	} else {
		term := math.Pow(1+rate, periods)
		if t == 1 {
			result = value*term + payment*(1+rate)*(term-1)/rate
		} else {
			result = value*term + payment*(term-1)/rate
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
	if period-life > 1 {
		return MakeErrorResultType(ErrorTypeNum, "Incorrect period for DB")
	}
	month := 12.0
	if argsNum == 5 && args[4].Type != ResultTypeEmpty {
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
	rate := 1 - math.Pow(salvage/cost, 1/life)
	rate = float64(int(rate*1000+0.5)) / 1000 // round to 3 decimal places
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
		return MakeErrorResultType(ErrorTypeNum, "DDB requires period to be not less than one")
	}
	if period > life {
		return MakeErrorResultType(ErrorTypeNum, "Incorrect period for DDB")
	}
	factor := 2.0
	if argsNum == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("DDB requires factor to be number argument")
		}
		factor = args[4].ValueNumber
		if factor < 0 {
			return MakeErrorResultType(ErrorTypeNum, "DDB requires factor to be non negative")
		}
	}

	return MakeNumberResult(getDDB(cost, salvage, life, period, factor))
}

// Disc implements the Excel DISC function.
func Disc(args []Result) Result {
	argsNum := len(args)
	if argsNum != 4 && argsNum != 5 {
		return MakeErrorResult("DISC requires four or five arguments")
	}
	settlementDate, maturityDate, errResult := getSettlementMaturity(args[0], args[1], "DISC")
	if errResult.Type == ResultTypeError {
		return errResult
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
	if argsNum == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("DISC requires basis to be number argument")
		}
		basis = int(args[4].ValueNumber)
		if !checkBasis(basis) {
			return MakeErrorResultType(ErrorTypeNum, "Incorrect basis argument for DISC")
		}
	}
	frac, errResult := yearFrac(settlementDate, maturityDate, basis)
	if errResult.Type == ResultTypeError {
		return errResult
	}
	return MakeNumberResult((redemption - pr) / redemption / frac)
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
	dollarde := dollarInt + dollarFrac/fraction
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

	dollarfr := dollarFrac*fraction/math.Pow(10, float64(int(math.Log10(fraction)))+1) + dollarInt
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
		return 0, 0, MakeErrorResultType(ErrorTypeNum, funcName+" requires fraction to be non negative number")
	}
	return dollar, fraction, empty
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
		return MakeErrorResultType(ErrorTypeNum, "EFFECT requires number of compounding periods to be 1 or more")
	}
	return MakeNumberResult(math.Pow((1+nominal/npery), npery) - 1)
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
	presentValue := 0.0
	if argsNum >= 4 && args[3].Type != ResultTypeEmpty {
		if args[3].Type != ResultTypeNumber {
			return MakeErrorResult("FV requires present value to be number argument")
		}
		presentValue = args[3].ValueNumber
	}
	t := 0
	if argsNum == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("FV requires type to be number argument")
		}
		t = int(args[4].ValueNumber)
		if t != 0 {
			t = 1
		}
	}
	return MakeNumberResult(fv(rate, nPer, pmt, presentValue, t))
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
		return MakeNumberResult(principal * (args[1].ValueNumber + 1))
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
		return MakeErrorResult("FVSCHEDULE requires schedule to be of number or array type")
	}
}

// Intrate implements the Excel INTRATE function.
func Intrate(args []Result) Result {
	argsNum := len(args)
	if argsNum != 4 && argsNum != 5 {
		return MakeErrorResult("INTRATE requires four or five arguments")
	}
	settlementDate, maturityDate, errResult := getSettlementMaturity(args[0], args[1], "INTRATE")
	if errResult.Type == ResultTypeError {
		return errResult
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
	if argsNum == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("INTRATE requires basis to be number argument")
		}
		basis = int(args[4].ValueNumber)
		if !checkBasis(basis) {
			return MakeErrorResultType(ErrorTypeNum, "Incorrect basis argument for INTRATE")
		}
	}
	frac, errResult := yearFrac(settlementDate, maturityDate, basis)
	if errResult.Type == ResultTypeError {
		return errResult
	}
	return MakeNumberResult((redemption - investment) / investment / frac)
}

// Ipmt implements the Excel IPMT function.
func Ipmt(args []Result) Result {
	argsNum := len(args)
	if argsNum < 4 || argsNum > 6 {
		return MakeErrorResult("IPMT requires number of arguments in range between four and six")
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
	if argsNum > 4 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("IPMT requires future value to be number argument")
		}
		futureValue = args[4].ValueNumber
	}
	t := 0
	if argsNum == 6 && args[5].Type != ResultTypeEmpty {
		if args[5].Type != ResultTypeNumber {
			return MakeErrorResult("IPMT requires type to be number argument")
		}
		t = int(args[5].ValueNumber)
		if t != 0 {
			t = 1
		}
	}

	return MakeNumberResult(ipmt(rate, period, nPer, presentValue, futureValue, t))
}

func ipmt(rate, period, nPer, presentValue, futureValue float64, t int) float64 {
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
			interest = fv(rate, period-2, payment, presentValue, 1) - payment
		} else {
			interest = fv(rate, period-1, payment, presentValue, 0)
		}
	}
	return interest * rate
}

// Irr implements the Excel IRR function.
func Irr(args []Result) Result {
	argsNum := len(args)
	if argsNum > 2 {
		return MakeErrorResult("IRR requires one or two arguments")
	}
	if args[0].Type != ResultTypeList && args[0].Type != ResultTypeArray {
		return MakeErrorResult("IRR requires values to be of array type")
	}
	valuesR := arrayFromRange(args[0])
	values := []float64{}
	for _, row := range valuesR {
		for _, vR := range row {
			if vR.Type == ResultTypeNumber && !vR.IsBoolean {
				values = append(values, vR.ValueNumber)
			}
		}
	}
	vlen := len(values)
	if len(values) < 2 {
		return MakeErrorResultType(ErrorTypeNum, "")
	}
	guess := 0.1
	if argsNum == 2 && args[1].Type != ResultTypeEmpty {
		if args[1].Type != ResultTypeNumber {
			return MakeErrorResult("IRR requires guess to be number argument")
		}
		guess = args[1].ValueNumber
		if guess <= -1 {
			return MakeErrorResult("IRR requires guess to be more than -1")
		}
	}

	dates := []float64{}

	for i := 0; i < vlen; i++ {
		if i == 0 {
			dates = append(dates, 0)
		} else {
			dates = append(dates, dates[i-1]+365)
		}
	}
	return irr(values, dates, guess)
}

// irr is used to calculate results for Irr and Xirr as method is the same
func irr(values, dates []float64, guess float64) Result {

	positive := false
	negative := false

	for i := 0; i < len(values); i++ {
		if values[i] > 0 {
			positive = true
		}
		if values[i] < 0 {
			negative = true
		}
	}

	if !positive || !negative {
		return MakeErrorResultType(ErrorTypeNum, "")
	}

	resultRate := guess
	epsMax := 1e-10
	iter := 0
	maxIter := 50
	isErr := false

	for {
		resultValue := irrResult(values, dates, resultRate)
		newRate := resultRate - resultValue/irrResultDeriv(values, dates, resultRate)
		epsRate := math.Abs(newRate - resultRate)
		resultRate = newRate
		iter++
		if epsRate <= epsMax || math.Abs(resultValue) <= epsMax {
			break
		}
		if iter > maxIter {
			isErr = true
			break
		}
	}
	if isErr || math.IsNaN(resultRate) || math.IsInf(resultRate, 0) {
		return MakeErrorResultType(ErrorTypeNum, "")
	}
	return MakeNumberResult(resultRate)
}

func irrResult(values, dates []float64, rate float64) float64 {
	r := rate + 1
	result := values[0]
	vlen := len(values)
	firstDate := dates[0]
	for i := 1; i < vlen; i++ {
		result += values[i] / math.Pow(r, (dates[i]-firstDate)/365)
	}
	return result
}

func irrResultDeriv(values, dates []float64, rate float64) float64 {
	r := rate + 1
	result := 0.0
	vlen := len(values)
	firstDate := dates[0]
	for i := 1; i < vlen; i++ {
		frac := (dates[i] - firstDate) / 365
		result -= frac * values[i] / math.Pow(r, frac+1)
	}
	return result
}

// Ispmt implements the Excel ISPMT function.
func Ispmt(args []Result) Result {
	if len(args) != 4 {
		return MakeErrorResult("ISPMT requires four arguments")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("ISPMT requires rate to be number argument")
	}
	rate := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("ISPMT requires period to be number argument")
	}
	period := args[1].ValueNumber
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("ISPMT requires number of periods to be number argument")
	}
	nPer := args[2].ValueNumber
	if nPer <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "ISPMT requires number of periods to be positive number argument")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("ISPMT requires present value to be number argument")
	}
	presentValue := args[3].ValueNumber

	return MakeNumberResult(presentValue * rate * (period/nPer - 1))
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
	mDuration := duration.ValueNumber / (1.0 + yield/freq)
	return MakeNumberResult(mDuration)
}

// Mirr implements the Excel MIRR function.
func Mirr(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("MIRR requires three arguments")
	}
	if args[0].Type != ResultTypeList && args[0].Type != ResultTypeArray {
		return MakeErrorResult("MIRR requires values to be of array type")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("MIRR requires finance rate to be number argument")
	}
	finRate := args[1].ValueNumber + 1
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("MIRR requires reinvest rate to be number argument")
	}
	reinvRate := args[2].ValueNumber + 1
	if reinvRate == 0 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "")
	}

	valuesR := arrayFromRange(args[0])
	n := float64(len(valuesR))
	npvInvest, npvReinvest := 0.0, 0.0
	powInvest, powReinvest := 1.0, 1.0
	hasPositive, hasNegative := false, false
	for _, row := range valuesR {
		for _, vR := range row {
			if vR.Type == ResultTypeNumber && !vR.IsBoolean {
				v := vR.ValueNumber
				if v == 0 {
					continue
				} else {
					if v > 0 {
						hasPositive = true
						npvReinvest += vR.ValueNumber * powReinvest
					} else {
						hasNegative = true
						npvInvest += vR.ValueNumber * powInvest
					}
					powInvest /= finRate
					powReinvest /= reinvRate
				}
			}
		}
	}

	if !hasPositive || !hasNegative {
		return MakeErrorResultType(ErrorTypeDivideByZero, "")
	}

	result := -npvReinvest / npvInvest
	result *= math.Pow(reinvRate, n-1)
	result = math.Pow(result, 1/(n-1))
	return MakeNumberResult(result - 1)
}

// Nominal implements the Excel NOMINAL function.
func Nominal(args []Result) Result {
	if len(args) != 2 {
		return MakeErrorResult("NOMINAL requires two arguments")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("NOMINAL requires nominal interest rate to be number argument")
	}
	effect := args[0].ValueNumber
	if effect <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "NOMINAL requires effect interest rate to be positive")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("NOMINAL requires number of compounding periods to be number argument")
	}
	npery := float64(int(args[1].ValueNumber))
	if npery < 1 {
		return MakeErrorResultType(ErrorTypeNum, "NOMINAL requires number of compounding periods to be 1 or more")
	}
	return MakeNumberResult((math.Pow(effect+1, 1/npery) - 1) * npery)
}

// Nper implements the Excel NPER function.
func Nper(args []Result) Result {
	argsNum := len(args)
	if argsNum < 3 || argsNum > 5 {
		return MakeErrorResult("NPER requires number of arguments in range of 3 and 5")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("NPER requires rate to be number argument")
	}
	rate := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("NPER requires payment to be number argument")
	}
	pmt := args[1].ValueNumber
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("NPER requires present value to be number argument")
	}
	presentValue := args[2].ValueNumber
	futureValue := 0.0
	if argsNum >= 4 && args[3].Type != ResultTypeEmpty {
		if args[3].Type != ResultTypeNumber {
			return MakeErrorResult("NPER requires future value to be number argument")
		}
		futureValue = args[3].ValueNumber
	}
	t := 0.0
	if argsNum == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("NPER requires type to be number argument")
		}
		t = args[4].ValueNumber
		if t != 0 {
			t = 1
		}
	}
	num := pmt*(1+rate*t) - futureValue*rate
	den := (presentValue*rate + pmt*(1+rate*t))
	return MakeNumberResult(math.Log(num/den) / math.Log(1+rate))
}

// Npv implements the Excel NPV function.
func Npv(args []Result) Result {
	argsNum := len(args)
	if argsNum < 2 {
		return MakeErrorResult("NPV requires two or more arguments")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("NPV requires rate to be number argument")
	}
	rate := args[0].ValueNumber
	if rate == -1 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "")
	}
	values := []float64{}
	for _, arg := range args[1:] {
		switch arg.Type {
		case ResultTypeNumber:
			values = append(values, arg.ValueNumber)
		case ResultTypeArray, ResultTypeList:
			rangeR := arrayFromRange(arg)
			for _, r := range rangeR {
				for _, vR := range r {
					if vR.Type == ResultTypeNumber && !vR.IsBoolean {
						values = append(values, vR.ValueNumber)
					}
				}
			}
		}
	}
	npv := 0.0
	for i, value := range values {
		npv += value / math.Pow(1+rate, float64(i)+1)
	}
	return MakeNumberResult(npv)
}

// Oddlprice implements the Excel ODDLPRICE function.
func Oddlprice(args []Result) Result {
	if len(args) != 8 && len(args) != 9 {
		return MakeErrorResult("ODDLPRICE requires eight or nine arguments")
	}
	settlementDate, maturityDate, errResult := getSettlementMaturity(args[0], args[1], "ODDLPRICE")
	if errResult.Type == ResultTypeError {
		return errResult
	}
	lastInterestDate, errResult := parseDate(args[2], "issue date", "ODDLPRICE")
	if errResult.Type == ResultTypeError {
		return errResult
	}
	if lastInterestDate >= settlementDate {
		return MakeErrorResultType(ErrorTypeNum, "Last interest date should be before settlement date")
	}
	rateResult := args[3]
	if rateResult.Type != ResultTypeNumber {
		return MakeErrorResult("ODDLPRICE requires rate of type number")
	}
	rate := rateResult.ValueNumber
	if rate < 0 {
		return MakeErrorResultType(ErrorTypeNum, "Rate should be non negative")
	}
	yieldResult := args[4]
	if yieldResult.Type != ResultTypeNumber {
		return MakeErrorResult("ODDLPRICE requires yield of type number")
	}
	yield := yieldResult.ValueNumber
	if yield < 0 {
		return MakeErrorResultType(ErrorTypeNum, "Yield should be non negative")
	}
	redemptionResult := args[5]
	if redemptionResult.Type != ResultTypeNumber {
		return MakeErrorResult("ODDLPRICE requires redemption of type number")
	}
	redemption := redemptionResult.ValueNumber
	if redemption < 0 {
		return MakeErrorResultType(ErrorTypeNum, "Yield should be non negative")
	}
	freqResult := args[6]
	if freqResult.Type != ResultTypeNumber {
		return MakeErrorResult("ODDLPRICE requires frequency of type number")
	}
	freq := float64(int(freqResult.ValueNumber))
	if !checkFreq(freq) {
		return MakeErrorResultType(ErrorTypeNum, "Incorrect frequence value")
	}
	basis := 0
	if len(args) == 8 && args[7].Type != ResultTypeEmpty {
		basisResult := args[7]
		if basisResult.Type != ResultTypeNumber {
			return MakeErrorResult("ODDLPRICE requires basis of type number")
		}
		basis = int(basisResult.ValueNumber)
		if !checkBasis(basis) {
			return MakeErrorResultType(ErrorTypeNum, "Incorrect basis value for ODDLPRICE")
		}
	}

	dc, errResult := yearFrac(lastInterestDate, maturityDate, basis)
	if errResult.Type == ResultTypeError {
		return errResult
	}
	dc *= freq
	dsc, errResult := yearFrac(settlementDate, maturityDate, basis)
	if errResult.Type == ResultTypeError {
		return errResult
	}
	dsc *= freq
	a, errResult := yearFrac(lastInterestDate, settlementDate, basis)
	if errResult.Type == ResultTypeError {
		return errResult
	}
	a *= freq

	p := redemption + dc*100*rate/freq
	p /= dsc*yield/freq + 1
	p -= a * 100 * rate / freq

	return MakeNumberResult(p)
}

// Oddlyield implements the Excel ODDLYIELD function.
func Oddlyield(args []Result) Result {
	if len(args) != 7 && len(args) != 8 {
		return MakeErrorResult("ODDLYIELD requires seven or eight arguments")
	}
	settlementDate, maturityDate, errResult := getSettlementMaturity(args[0], args[1], "ODDLYIELD")
	if errResult.Type == ResultTypeError {
		return errResult
	}
	lastInterestDate, errResult := parseDate(args[2], "issue date", "ODDLPRICE")
	if errResult.Type == ResultTypeError {
		return errResult
	}
	if lastInterestDate >= settlementDate {
		return MakeErrorResultType(ErrorTypeNum, "Last interest date should be before settlement date")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("ODDLYIELD requires rate of type number")
	}
	rate := args[3].ValueNumber
	if rate < 0 {
		return MakeErrorResultType(ErrorTypeNum, "Rate should be non negative")
	}
	if args[4].Type != ResultTypeNumber {
		return MakeErrorResult("ODDLYIELD requires present value of type number")
	}
	pr := args[4].ValueNumber
	if pr <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "Present value should be positive")
	}
	if args[5].Type != ResultTypeNumber {
		return MakeErrorResult("ODDLYIELD requires redemption of type number")
	}
	redemption := args[5].ValueNumber
	if redemption < 0 {
		return MakeErrorResultType(ErrorTypeNum, "Yield should be non negative")
	}
	if args[6].Type != ResultTypeNumber {
		return MakeErrorResult("ODDLYIELD requires frequency of type number")
	}
	freq := float64(int(args[6].ValueNumber))
	if !checkFreq(freq) {
		return MakeErrorResultType(ErrorTypeNum, "Incorrect frequence value")
	}
	basis := 0
	if len(args) == 8 && args[7].Type != ResultTypeEmpty {
		if args[7].Type != ResultTypeNumber {
			return MakeErrorResult("ODDLYIELD requires basis of type number")
		}
		basis = int(args[7].ValueNumber)
		if !checkBasis(basis) {
			return MakeErrorResultType(ErrorTypeNum, "Incorrect basis value for ODDLYIELD")
		}
	}

	dc, errResult := yearFrac(lastInterestDate, maturityDate, basis)
	if errResult.Type == ResultTypeError {
		return errResult
	}
	dc *= freq
	dsc, errResult := yearFrac(settlementDate, maturityDate, basis)
	if errResult.Type == ResultTypeError {
		return errResult
	}
	dsc *= freq
	a, errResult := yearFrac(lastInterestDate, settlementDate, basis)
	if errResult.Type == ResultTypeError {
		return errResult
	}
	a *= freq

	yield := redemption + dc*100*rate/freq
	yield /= pr + a*100*rate/freq
	yield--
	yield *= freq / dsc

	return MakeNumberResult(yield)
}

// Pduration implements the Excel PDURATION function.
func Pduration(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("PDURATION requires three arguments")
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
	return MakeNumberResult((math.Log10(specifiedValue) - math.Log10(currentValue)) / math.Log10(1+rate))
}

// Pmt implements the Excel PMT function.
func Pmt(args []Result) Result {
	argsNum := len(args)
	if argsNum < 3 || argsNum > 5 {
		return MakeErrorResult("PMT requires number of arguments in range of 3 and 5")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("PMT requires rate to be number argument")
	}
	rate := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("PMT requires number of periods to be number argument")
	}
	nPer := args[1].ValueNumber
	if nPer == 0 {
		return MakeErrorResultType(ErrorTypeNum, "PMT requires number of periods to be not equal to 0")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("PMT requires present value to be number argument")
	}
	presentValue := args[2].ValueNumber
	futureValue := 0.0
	if argsNum >= 4 && args[3].Type != ResultTypeEmpty {
		if args[3].Type != ResultTypeNumber {
			return MakeErrorResult("PMT requires future value to be number argument")
		}
		futureValue = args[3].ValueNumber
	}
	t := 0.0
	if argsNum == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("PMT requires type to be number argument")
		}
		t = args[4].ValueNumber
		if t != 0 {
			t = 1
		}
	}
	var result float64
	if rate == 0 {
		result = (presentValue + futureValue) / nPer
	} else {
		term := math.Pow(1+rate, nPer)
		if t == 1 {
			result = (futureValue*rate/(term-1) + presentValue*rate/(1-1/term)) / (1 + rate)
		} else {
			result = futureValue*rate/(term-1) + presentValue*rate/(1-1/term)
		}
	}
	return MakeNumberResult(-result)
}

// Ppmt implements the Excel PPPMT function.
func Ppmt(args []Result) Result {
	argsNum := len(args)
	if argsNum < 4 || argsNum > 6 {
		return MakeErrorResult("PPMT requires number of arguments in range of four and six")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("PPMT requires rate to be number argument")
	}
	rate := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("PPMT requires period to be number argument")
	}
	period := args[1].ValueNumber
	if period <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "PPMT requires period to be positive")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("PPMT requires number of periods to be number argument")
	}
	nPer := args[2].ValueNumber
	if nPer < period {
		return MakeErrorResultType(ErrorTypeNum, "PPMT requires number of periods to be not less than period")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("PPMT requires present value to be number argument")
	}
	presentValue := args[3].ValueNumber
	futureValue := 0.0
	if argsNum >= 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("PPMT requires future value to be number argument")
		}
		futureValue = args[4].ValueNumber
	}
	t := 0
	if argsNum == 6 && args[5].Type != ResultTypeEmpty {
		if args[5].Type != ResultTypeNumber {
			return MakeErrorResult("PPMT requires type to be number argument")
		}
		t = int(args[5].ValueNumber)
		if t != 0 {
			t = 1
		}
	}
	return MakeNumberResult(pmt(rate, nPer, presentValue, futureValue, t) - ipmt(rate, period, nPer, presentValue, futureValue, t))
}

// Price implements the Excel PRICE function.
func Price(args []Result) Result {
	argsNum := len(args)
	if argsNum != 6 && argsNum != 7 {
		return MakeErrorResult("PRICE requires six or seven arguments")
	}
	settlementDate, maturityDate, errResult := getSettlementMaturity(args[0], args[1], "PRICE")
	if errResult.Type == ResultTypeError {
		return errResult
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("PRICE requires rate of type number")
	}
	rate := args[2].ValueNumber
	if rate < 0 {
		return MakeErrorResultType(ErrorTypeNum, "PRICE requires rate to not be negative")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("PRICE requires yield of type number")
	}
	yield := args[3].ValueNumber
	if yield < 0 {
		return MakeErrorResultType(ErrorTypeNum, "PRICE requires yield to not be negative")
	}
	if args[4].Type != ResultTypeNumber {
		return MakeErrorResult("PRICE requires redemption to be number argument")
	}
	redemption := args[4].ValueNumber
	if redemption <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "PRICE requires redemption to be positive number argument")
	}
	freqResult := args[5]
	if freqResult.Type != ResultTypeNumber {
		return MakeErrorResult("PRICE requires frequency of type number")
	}
	freq := freqResult.ValueNumber
	if !checkFreq(freq) {
		return MakeErrorResultType(ErrorTypeNum, "Incorrect frequence value")
	}
	basis := 0
	if argsNum == 7 && args[6].Type != ResultTypeEmpty {
		if args[6].Type != ResultTypeNumber {
			return MakeErrorResult("PRICE requires basis to be number argument")
		}
		basis = int(args[6].ValueNumber)
		if !checkBasis(basis) {
			return MakeErrorResultType(ErrorTypeNum, "Incorrect basis argument for PRICE")
		}
	}
	price, errResult := getPrice(settlementDate, maturityDate, rate, yield, redemption, freq, basis)
	if errResult.Type == ResultTypeError {
		return errResult
	}
	return MakeNumberResult(price)
}

func getPrice(settlementDate, maturityDate, rate, yield, redemption, freqF float64, basis int) (float64, Result) {
	freq := int(freqF)
	e := coupdays(settlementDate, maturityDate, freq, basis)
	dsc := coupdaysnc(settlementDate, maturityDate, freq, basis) / e
	n, errResult := coupnum(settlementDate, maturityDate, freq, basis)
	if errResult.Type == ResultTypeError {
		return 0, errResult
	}
	a := coupdaybs(settlementDate, maturityDate, freq, basis)
	ret := redemption / math.Pow(1+yield/freqF, n-1+dsc)
	ret -= 100 * rate / freqF * a / e
	t1 := 100 * rate / freqF
	t2 := 1 + yield/freqF
	for k := 0.0; k < n; k++ {
		ret += t1 / math.Pow(t2, k+dsc)
	}
	return ret, MakeEmptyResult()
}

// Pricedisc implements the Excel PRICEDISC function.
func Pricedisc(args []Result) Result {
	argsNum := len(args)
	if argsNum != 4 && argsNum != 5 {
		return MakeErrorResult("PRICEDISC requires four or five arguments")
	}
	settlementDate, maturityDate, errResult := getSettlementMaturity(args[0], args[1], "PRICEDISC")
	if errResult.Type == ResultTypeError {
		return errResult
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("PRICEDISC requires discount of type number")
	}
	discount := args[2].ValueNumber
	if discount <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "PRICEDISC requires discount to be positive")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("PRICEDISC requires redemption of type number")
	}
	redemption := args[3].ValueNumber
	if redemption <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "PRICEDISC requires redemption to be positive")
	}
	basis := 0
	if argsNum == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("PRICEDISC requires basis to be number argument")
		}
		basis = int(args[4].ValueNumber)
		if !checkBasis(basis) {
			return MakeErrorResultType(ErrorTypeNum, "Incorrect basis argument for PRICEDISC")
		}
	}
	yf, errResult := yearFrac(settlementDate, maturityDate, basis)
	if errResult.Type == ResultTypeError {
		return errResult
	}
	return MakeNumberResult(redemption * (1 - discount*yf))
}

// Pricemat implements the Excel PRICEMAT function.
func Pricemat(args []Result) Result {
	argsNum := len(args)
	if argsNum != 5 && argsNum != 6 {
		return MakeErrorResult("PRICEMAT requires five or six arguments")
	}
	settlementDate, maturityDate, errResult := getSettlementMaturity(args[0], args[1], "PRICEMAT")
	if errResult.Type == ResultTypeError {
		return errResult
	}
	issueDate, errResult := parseDate(args[2], "issue date", "PRICEMAT")
	if errResult.Type == ResultTypeError {
		return errResult
	}
	if issueDate >= settlementDate {
		return MakeErrorResult("PRICEMAT requires issue date to be before settlement date")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("PRICEMAT requires rate of type number")
	}
	rate := args[3].ValueNumber
	if rate < 0 {
		return MakeErrorResultType(ErrorTypeNum, "PRICEMAT requires rate to be non negative")
	}
	if args[4].Type != ResultTypeNumber {
		return MakeErrorResult("PRICEMAT requires yield of type number")
	}
	yield := args[4].ValueNumber
	if yield < 0 {
		return MakeErrorResultType(ErrorTypeNum, "PRICEMAT requires yield to be non negative")
	}
	basis := 0
	if argsNum == 6 && args[5].Type != ResultTypeEmpty {
		if args[5].Type != ResultTypeNumber {
			return MakeErrorResult("PRICEMAT requires basis to be number argument")
		}
		basis = int(args[5].ValueNumber)
		if !checkBasis(basis) {
			return MakeErrorResultType(ErrorTypeNum, "Incorrect basis argument for PRICEMAT")
		}
	}
	dsm, errResult := yearFrac(settlementDate, maturityDate, basis)
	if errResult.Type == ResultTypeError {
		return errResult
	}
	dim, errResult := yearFrac(issueDate, maturityDate, basis)
	if errResult.Type == ResultTypeError {
		return errResult
	}
	dis, errResult := yearFrac(issueDate, settlementDate, basis)
	if errResult.Type == ResultTypeError {
		return errResult
	}

	num := 1 + dim*rate
	den := 1 + dsm*yield

	return MakeNumberResult((num/den - dis*rate) * 100)
}

// Pv implements the Excel PV function.
func Pv(args []Result) Result {
	argsNum := len(args)
	if argsNum < 3 || argsNum > 5 {
		return MakeErrorResult("PV requires number of arguments in range of 3 and 5")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("PV requires rate to be number argument")
	}
	rate := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("PV requires number of periods to be number argument")
	}
	nPer := args[1].ValueNumber
	if nPer != float64(int(nPer)) {
		return MakeErrorResultType(ErrorTypeNum, "PV requires number of periods to be integer number argument")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("PV requires payment to be number argument")
	}
	pmt := args[2].ValueNumber
	futureValue := 0.0
	if argsNum >= 4 && args[3].Type != ResultTypeEmpty {
		if args[3].Type != ResultTypeNumber {
			return MakeErrorResult("PV requires future value to be number argument")
		}
		futureValue = args[3].ValueNumber
	}
	t := 0.0
	if argsNum == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("PV requires type to be number argument")
		}
		t = args[4].ValueNumber
		if t != 0 {
			t = 1
		}
	}
	if rate == 0 {
		return MakeNumberResult(-pmt*nPer - futureValue)
	} else {
		return MakeNumberResult((((1-math.Pow(1+rate, nPer))/rate)*pmt*(1+rate*t) - futureValue) / math.Pow(1+rate, nPer))
	}
}

// Rate implements the Excel RATE function.
func Rate(args []Result) Result {
	argsNum := len(args)
	if argsNum < 3 || argsNum > 6 {
		return MakeErrorResult("RATE requires number of arguments in range of three and six")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("RATE requires number of periods to be number argument")
	}
	nPer := args[0].ValueNumber
	if nPer != float64(int(nPer)) {
		return MakeErrorResultType(ErrorTypeNum, "RATE requires number of periods to be integer number argument")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("RATE requires payment to be number argument")
	}
	pmt := args[1].ValueNumber
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("RATE requires present value to be number argument")
	}
	presentValue := args[2].ValueNumber
	futureValue := 0.0
	if argsNum >= 4 && args[3].Type != ResultTypeEmpty {
		if args[3].Type != ResultTypeNumber {
			return MakeErrorResult("RATE requires future value to be number argument")
		}
		futureValue = args[3].ValueNumber
	}
	t := 0.0
	if argsNum >= 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("RATE requires type to be number argument")
		}
		t = args[4].ValueNumber
		if t != 0 {
			t = 1
		}
	}
	guess := 0.1
	if argsNum >= 6 && args[5].Type != ResultTypeEmpty {
		if args[5].Type != ResultTypeNumber {
			return MakeErrorResult("RATE requires guess to be number argument")
		}
		guess = args[5].ValueNumber
	}

	maxIter := 100
	iter := 0
	close := false
	epsMax := 1e-6

	rate := guess
	for iter < maxIter && !close {
		t1 := math.Pow(rate+1, nPer)
		t2 := math.Pow(rate+1, nPer-1)
		rt := rate*t + 1
		p0 := pmt * (t1 - 1)
		f1 := futureValue + t1*presentValue + p0*rt/rate
		f2 := nPer*t2*presentValue - p0*rt/math.Pow(rate, 2)
		f3 := (nPer*pmt*t2*rt + p0*t) / rate

		delta := f1 / (f2 + f3)
		if math.Abs(delta) < epsMax {
			close = true
		}
		iter++
		rate -= delta
	}

	return MakeNumberResult(rate)
}

// Received implements the Excel RECEIVED function.
func Received(args []Result) Result {
	argsNum := len(args)
	if argsNum != 4 && argsNum != 5 {
		return MakeErrorResult("RECEIVED requires four or five arguments")
	}
	settlementDate, maturityDate, errResult := getSettlementMaturity(args[0], args[1], "RECEIVED")
	if errResult.Type == ResultTypeError {
		return errResult
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("RECEIVED requires investment to be number argument")
	}
	investment := args[2].ValueNumber
	if investment <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "RECEIVED requires investment to be positive number argument")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("RECEIVED requires discount to be number argument")
	}
	discount := args[3].ValueNumber
	if discount <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "RECEIVED requires discount to be positive number argument")
	}
	basis := 0
	if argsNum == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("RECEIVED requires basis to be number argument")
		}
		basis = int(args[4].ValueNumber)
		if !checkBasis(basis) {
			return MakeErrorResultType(ErrorTypeNum, "Incorrect basis argument for RECEIVED")
		}
	}
	frac, errResult := yearFrac(settlementDate, maturityDate, basis)
	if errResult.Type == ResultTypeError {
		return errResult
	}
	return MakeNumberResult(investment / (1 - discount*frac))
}

// Rri implements the Excel RRI function.
func Rri(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("RRI requires three arguments")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("RRI requires number of periods to be number argument")
	}
	nPer := args[0].ValueNumber
	if nPer <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "RRI requires number of periods to be positive")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("RRI requires present value to be number argument")
	}
	presentValue := args[1].ValueNumber
	if presentValue <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "RRI requires present value to be positive")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("RRI requires future value to be number argument")
	}
	futureValue := args[2].ValueNumber
	if futureValue < 0 {
		return MakeErrorResultType(ErrorTypeNum, "RRI requires future value to be non negative")
	}

	return MakeNumberResult(math.Pow(futureValue/presentValue, 1/nPer) - 1)
}

// Sln implements the Excel SLN function.
func Sln(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("SLN requires three arguments")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("SLN requires cost to be number argument")
	}
	cost := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("SLN requires salvage to be number argument")
	}
	salvage := args[1].ValueNumber
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("SLN requires life to be number argument")
	}
	life := args[2].ValueNumber
	if life == 0 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "SLN requires life to be non zero")
	}

	return MakeNumberResult((cost - salvage) / life)
}

// Syd implements the Excel SYD function.
func Syd(args []Result) Result {
	if len(args) != 4 {
		return MakeErrorResult("SYD requires four arguments")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("SYD requires cost to be number argument")
	}
	cost := args[0].ValueNumber
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("SYD requires salvage to be number argument")
	}
	salvage := args[1].ValueNumber
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("SYD requires life to be number argument")
	}
	life := args[2].ValueNumber
	if life <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "SYD requires life to be positive")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("SYD requires period to be number argument")
	}
	per := args[3].ValueNumber
	if per <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "SYD requires period to be positive")
	}
	if per > life {
		return MakeErrorResultType(ErrorTypeNum, "SYD requires period to be equal or less than life")
	}

	num := (cost - salvage) * (life - per + 1) * 2
	den := life * (life + 1)

	return MakeNumberResult(num / den)
}

// Tbilleq implements the Excel TBILLEQ function.
func Tbilleq(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("TBILLEQ requires three arguments")
	}
	settlementDate, maturityDate, errResult := getSettlementMaturity(args[0], args[1], "TBILLEQ")
	if errResult.Type == ResultTypeError {
		return errResult
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("TBILLEQ requires discount to be number argument")
	}
	dsm := maturityDate - settlementDate
	if dsm > 365 {
		return MakeErrorResultType(ErrorTypeNum, "TBILLEQ requires maturity to be not more than one year after settlement")
	}
	discount := args[2].ValueNumber
	if discount <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "TBILLEQ requires discount to be positive number argument")
	}
	return MakeNumberResult((365 * discount) / (360 - discount*dsm))
}

// Tbillprice implements the Excel TBILLPRICE function.
func Tbillprice(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("TBILLPRICE requires three arguments")
	}
	settlementDate, maturityDate, errResult := getSettlementMaturity(args[0], args[1], "TBILLPRICE")
	if errResult.Type == ResultTypeError {
		return errResult
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("TBILLPRICE requires discount to be number argument")
	}
	dsm := maturityDate - settlementDate
	if dsm > 365 {
		return MakeErrorResultType(ErrorTypeNum, "TBILLPRICE requires maturity to be not more than one year after settlement")
	}
	discount := args[2].ValueNumber
	if discount <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "TBILLPRICE requires discount to be positive number argument")
	}
	return MakeNumberResult(100 * (1 - discount*dsm/360))
}

// Tbillyield implements the Excel TBILLYIELD function.
func Tbillyield(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("TBILLYIELD requires three arguments")
	}
	settlementDate, maturityDate, errResult := getSettlementMaturity(args[0], args[1], "TBILLYIELD")
	if errResult.Type == ResultTypeError {
		return errResult
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("TBILLYIELD requires discount to be number argument")
	}
	dsm := maturityDate - settlementDate
	if dsm > 365 {
		return MakeErrorResultType(ErrorTypeNum, "TBILLYIELD requires maturity to be not more than one year after settlement")
	}
	pr := args[2].ValueNumber
	if pr <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "TBILLYIELD requires pr to be positive number argument")
	}
	m1 := (100 - pr) / pr
	m2 := 360 / dsm
	return MakeNumberResult(m1 * m2)
}

// Vdb implements the Excel VDB function.
func Vdb(args []Result) Result {
	argsNum := len(args)
	if argsNum < 5 || argsNum > 7 {
		return MakeErrorResult("VDB requires number of arguments to be in range between five and seven")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("VDB requires cost to be number argument")
	}
	cost := args[0].ValueNumber
	if cost < 0 {
		return MakeErrorResultType(ErrorTypeNum, "VDB requires cost to be non negative")
	}
	if args[1].Type != ResultTypeNumber {
		return MakeErrorResult("VDB requires salvage to be number argument")
	}
	salvage := args[1].ValueNumber
	if salvage < 0 {
		return MakeErrorResultType(ErrorTypeNum, "VDB requires salvage to be non negative")
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("VDB requires life to be number argument")
	}
	life := args[2].ValueNumber
	if life == 0 {
		return MakeErrorResultType(ErrorTypeDivideByZero, "VDB requires life to be positive")
	}
	if life < 0 {
		return MakeErrorResultType(ErrorTypeNum, "VDB requires life to be positive")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("VDB requires start period to be number argument")
	}
	startPeriod := args[3].ValueNumber
	if startPeriod < 0 {
		return MakeErrorResultType(ErrorTypeNum, "VDB requires start period to be not less than one")
	}
	if args[4].Type != ResultTypeNumber {
		return MakeErrorResult("VDB requires end period to be number argument")
	}
	endPeriod := args[4].ValueNumber
	if startPeriod > endPeriod {
		return MakeErrorResultType(ErrorTypeNum, "Incorrect start period for VDB")
	}
	if endPeriod > life {
		return MakeErrorResultType(ErrorTypeNum, "Incorrect end period for VDB")
	}
	factor := 2.0
	if argsNum > 5 {
		if args[5].Type == ResultTypeEmpty {
			factor = 0.0
		} else {
			if args[5].Type != ResultTypeNumber {
				return MakeErrorResult("VDB requires factor to be number argument")
			}
			factor = args[5].ValueNumber
			if factor < 0 {
				return MakeErrorResultType(ErrorTypeNum, "VDB requires factor to be non negative")
			}
		}
	}
	noSwitch := false
	if argsNum > 6 && args[6].Type != ResultTypeEmpty {
		if args[6].Type != ResultTypeNumber {
			return MakeErrorResult("VDB requires no_switch to be number argument")
		}
		noSwitch = args[6].ValueNumber != 0
	}

	vdb := 0.0
	startInt := math.Floor(startPeriod)
	endInt := math.Ceil(endPeriod)

	if noSwitch {
		for i := startInt + 1; i <= endInt; i++ {
			term := getDDB(cost, salvage, life, i, factor)
			if i == startInt+1 {
				term *= math.Min(endPeriod, startInt+1) - startPeriod
			} else if i == endInt {
				term *= endPeriod + 1 - endInt
			}
			vdb += term
		}
	} else {
		life1 := life
		var part float64
		if !approxEqual(startPeriod, math.Floor(startPeriod)) {
			if factor == 1 {
				l2 := life / 2
				if startPeriod > l2 || approxEqual(startPeriod, l2) {
					part = startPeriod - l2
					startPeriod = l2
					endPeriod -= part
					life1++
				}
			}
		}
		if factor != 0 {
			cost -= interVDB(cost, salvage, life, life1, startPeriod, factor)
		}
		vdb = interVDB(cost, salvage, life, life-startPeriod, endPeriod-startPeriod, factor)
	}
	return MakeNumberResult(vdb)
}

func interVDB(cost, salvage, life, life1, period, factor float64) float64 {
	var ddb, term float64
	vdb := 0.0
	endInt := math.Ceil(period)
	cs := cost - salvage
	nowSln := false
	sln := 0.0

	for i := 1.0; i <= endInt; i++ {
		if !nowSln {
			ddb = getDDB(cost, salvage, life, i, factor)
			sln = cs / (life - i + 1)
			if sln > ddb {
				term = sln
				nowSln = true
			} else {
				term = ddb
				cs -= ddb
			}
		} else {
			term = sln
		}
		if i == endInt {
			term *= period + 1 - endInt
		}
		vdb += term
	}
	return vdb
}

func getDDB(cost, salvage, life, period, factor float64) float64 {
	var oldValue float64
	rate := factor / life
	if rate >= 1 {
		rate = 1
		if period == 1 {
			oldValue = cost
		} else {
			oldValue = 0
		}
	} else {
		oldValue = cost * math.Pow(1-rate, period-1)
	}
	newValue := cost * math.Pow(1-rate, period)

	var ddb float64

	if newValue < salvage {
		ddb = oldValue - salvage
	} else {
		ddb = oldValue - newValue
	}
	if ddb < 0 {
		ddb = 0
	}
	return ddb
}

// Yielddisc implements the Excel YIELDDISC function.
func Yielddisc(args []Result) Result {
	argsNum := len(args)
	if argsNum != 4 && argsNum != 5 {
		return MakeErrorResult("YIELDDISC requires four or five arguments")
	}
	settlementDate, maturityDate, errResult := getSettlementMaturity(args[0], args[1], "YIELDDISC")
	if errResult.Type == ResultTypeError {
		return errResult
	}
	if args[2].Type != ResultTypeNumber {
		return MakeErrorResult("YIELDDISC requires pr to be number argument")
	}
	pr := args[2].ValueNumber
	if pr <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "YIELDDISC requires pr to be positive number argument")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("YIELDDISC requires redemption to be number argument")
	}
	redemption := args[3].ValueNumber
	if redemption <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "YIELDDISC requires redemption to be positive number argument")
	}
	basis := 0
	if argsNum == 5 && args[4].Type != ResultTypeEmpty {
		if args[4].Type != ResultTypeNumber {
			return MakeErrorResult("YIELDDISC requires basis to be number argument")
		}
		basis = int(args[4].ValueNumber)
		if !checkBasis(basis) {
			return MakeErrorResultType(ErrorTypeNum, "Incorrect basis argument for YIELDDISC")
		}
	}
	frac, errResult := yearFrac(settlementDate, maturityDate, basis)
	if errResult.Type == ResultTypeError {
		return errResult
	}

	return MakeNumberResult((redemption/pr - 1) / frac)
}

func approxEqual(a, b float64) bool {
	return math.Abs(a-b) < 1.0e-6
}

// Xirr implements the Excel XIRR function.
func Xirr(args []Result) Result {
	argsNum := len(args)
	if argsNum != 2 && argsNum != 3 {
		return MakeErrorResult("XIRR requires two or three arguments")
	}
	xStruct, errResult := getXargs(args[0], args[1], "XIRR")
	if errResult.Type == ResultTypeError {
		return errResult
	}
	values := xStruct.values
	dates := xStruct.dates
	guess := 0.1
	if argsNum == 3 && args[2].Type != ResultTypeEmpty {
		if args[2].Type != ResultTypeNumber {
			return MakeErrorResult("XIRR requires guess to be number argument")
		}
		guess = args[2].ValueNumber
		if guess <= -1 {
			return MakeErrorResult("XIRR requires guess to be more than -1")
		}
	}
	return irr(values, dates, guess)
}

// Xnpv implements the Excel XNPV function.
func Xnpv(args []Result) Result {
	if len(args) != 3 {
		return MakeErrorResult("XNPV requires three arguments")
	}
	if args[0].Type != ResultTypeNumber {
		return MakeErrorResult("XNPV requires rate to be number argument")
	}
	rate := args[0].ValueNumber
	if rate <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "XNPV requires rate to be positive")
	}
	xStruct, errResult := getXargs(args[1], args[2], "XNPV")
	if errResult.Type == ResultTypeError {
		return errResult
	}
	values := xStruct.values
	dates := xStruct.dates
	xnpv := 0.0
	firstDate := dates[0]
	for i, value := range values {
		xnpv += value / math.Pow(1+rate, (dates[i]-firstDate)/365)
	}
	return MakeNumberResult(xnpv)
}

type xargs struct {
	values []float64
	dates  []float64
}

func getXargs(valuesR, datesR Result, funcName string) (*xargs, Result) {
	if valuesR.Type != ResultTypeList && valuesR.Type != ResultTypeArray {
		return nil, MakeErrorResult(funcName + " requires values to be of array type")
	}
	valuesArr := arrayFromRange(valuesR)
	values := []float64{}
	for _, row := range valuesArr {
		for _, vR := range row {
			if vR.Type == ResultTypeNumber && !vR.IsBoolean {
				values = append(values, vR.ValueNumber)
			} else {
				return nil, MakeErrorResult(funcName + "requires values to be numbers")
			}
		}
	}
	vlen := len(values)
	if len(values) < 2 {
		return nil, MakeErrorResultType(ErrorTypeNum, "")
	}
	if datesR.Type != ResultTypeList && datesR.Type != ResultTypeArray {
		return nil, MakeErrorResult(funcName + " requires dates to be of array type")
	}
	datesArr := arrayFromRange(datesR)
	dates := []float64{}
	lastDate := 0.0
	for _, row := range datesArr {
		for _, vR := range row {
			if vR.Type == ResultTypeNumber && !vR.IsBoolean {
				newDate := float64(int(vR.ValueNumber))
				if newDate < lastDate {
					return nil, MakeErrorResultType(ErrorTypeNum, funcName+" requires dates to be in ascending order")
				}
				dates = append(dates, newDate)
				lastDate = newDate
			} else {
				return nil, MakeErrorResult(funcName + "requires dates to be numbers")
			}
		}
	}
	if len(dates) != vlen {
		return nil, MakeErrorResultType(ErrorTypeNum, "")
	}
	return &xargs{values, dates}, MakeEmptyResult()
}

// Yield implements the Excel YIELD function.
func Yield(args []Result) Result {
	argsNum := len(args)
	if argsNum != 6 && argsNum != 7 {
		return MakeErrorResult("YIELD requires six or seven arguments")
	}
	settlementDate, maturityDate, errResult := getSettlementMaturity(args[0], args[1], "YIELD")
	if errResult.Type == ResultTypeError {
		return errResult
	}
	rateResult := args[2]
	if rateResult.Type != ResultTypeNumber {
		return MakeErrorResult("YIELD requires rate of type number")
	}
	rate := rateResult.ValueNumber
	if rate < 0 {
		return MakeErrorResultType(ErrorTypeNum, "Rate should be non negative")
	}
	prResult := args[3]
	if prResult.Type != ResultTypeNumber {
		return MakeErrorResult("YIELD requires pr of type number")
	}
	pr := prResult.ValueNumber
	if pr <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "pr should be positive")
	}
	redemptionResult := args[4]
	if redemptionResult.Type != ResultTypeNumber {
		return MakeErrorResult("YIELD requires redemption of type number")
	}
	redemption := redemptionResult.ValueNumber
	if redemption < 0 {
		return MakeErrorResultType(ErrorTypeNum, "Yield should be non negative")
	}
	freqResult := args[5]
	if freqResult.Type != ResultTypeNumber {
		return MakeErrorResult("YIELD requires frequency of type number")
	}
	freq := float64(int(freqResult.ValueNumber))
	if !checkFreq(freq) {
		return MakeErrorResultType(ErrorTypeNum, "Incorrect frequence value")
	}
	basis := 0
	if argsNum == 7 && args[6].Type != ResultTypeEmpty {
		basisResult := args[6]
		if basisResult.Type != ResultTypeNumber {
			return MakeErrorResult("YIELD requires basis of type number")
		}
		basis = int(basisResult.ValueNumber)
		if !checkBasis(basis) {
			return MakeErrorResultType(ErrorTypeNum, "Incorrect basis value for YIELD")
		}
	}

	priceN := 0.0
	yield1 := 0.0
	yield2 := 1.0
	price1, errResult := getPrice(settlementDate, maturityDate, rate, yield1, redemption, freq, basis)
	if errResult.Type == ResultTypeError {
		return errResult
	}
	price2, errResult := getPrice(settlementDate, maturityDate, rate, yield2, redemption, freq, basis)
	if errResult.Type == ResultTypeError {
		return errResult
	}

	yieldN := (yield2 - yield1) * 0.5

	for iter := 0; iter < 100 && priceN != pr; iter++ {
		priceN, errResult = getPrice(settlementDate, maturityDate, rate, yieldN, redemption, freq, basis)
		if errResult.Type == ResultTypeError {
			return errResult
		}
		if pr == price1 {
			return MakeNumberResult(yield1)
		} else if pr == price2 {
			return MakeNumberResult(yield2)
		} else if pr == priceN {
			return MakeNumberResult(yieldN)
		} else if pr < price2 {
			yield2 *= 2.0
			price2, errResult = getPrice(settlementDate, maturityDate, rate, yield2, redemption, freq, basis)
			if errResult.Type == ResultTypeError {
				return errResult
			}
			yieldN = (yield2 - yield1) * 0.5
		} else {
			if pr < priceN {
				yield1 = yieldN
				price1 = priceN
			} else {
				yield2 = yieldN
				price2 = priceN
			}
			yieldN = yield2 - (yield2-yield1)*((pr-price2)/(price1-price2))
		}
	}

	return MakeNumberResult(yieldN)
}

// Yieldmat implements the Excel YIELDMAT function.
func Yieldmat(args []Result) Result {
	argsNum := len(args)
	if argsNum != 5 && argsNum != 6 {
		return MakeErrorResult("YIELDMAT requires five or six arguments")
	}
	settlementDate, maturityDate, errResult := getSettlementMaturity(args[0], args[1], "YIELDMAT")
	if errResult.Type == ResultTypeError {
		return errResult
	}
	issueDate, errResult := parseDate(args[2], "issue date", "YIELDMAT")
	if errResult.Type == ResultTypeError {
		return errResult
	}
	if issueDate >= settlementDate {
		return MakeErrorResult("YIELDMAT requires issue date to be before settlement date")
	}
	if args[3].Type != ResultTypeNumber {
		return MakeErrorResult("YIELDMAT requires rate of type number")
	}
	rate := args[3].ValueNumber
	if rate < 0 {
		return MakeErrorResultType(ErrorTypeNum, "YIELDMAT requires rate to be non negative")
	}
	if args[4].Type != ResultTypeNumber {
		return MakeErrorResult("YIELDMAT requires yield of type number")
	}
	pr := args[4].ValueNumber
	if pr <= 0 {
		return MakeErrorResultType(ErrorTypeNum, "YIELDMAT requires pr to be positive")
	}
	basis := 0
	if argsNum == 6 && args[5].Type != ResultTypeEmpty {
		if args[5].Type != ResultTypeNumber {
			return MakeErrorResult("YIELDMAT requires basis to be number argument")
		}
		basis = int(args[5].ValueNumber)
		if !checkBasis(basis) {
			return MakeErrorResultType(ErrorTypeNum, "Incorrect basis argument for YIELDMAT")
		}
	}
	dim, errResult := yearFrac(issueDate, maturityDate, basis)
	if errResult.Type == ResultTypeError {
		return errResult
	}
	dis, errResult := yearFrac(issueDate, settlementDate, basis)
	if errResult.Type == ResultTypeError {
		return errResult
	}
	dsm, errResult := yearFrac(settlementDate, maturityDate, basis)
	if errResult.Type == ResultTypeError {
		return errResult
	}

	y := 1 + dim*rate
	y /= pr/100 + dis*rate
	y--
	y /= dsm

	return MakeNumberResult(y)
}
