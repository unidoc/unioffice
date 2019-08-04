// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package format

import (
	"bytes"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/unidoc/unioffice"
)

// constants used when formatting generic values to determine when to start
// rounding
const maxGeneric = 1e11
const minGeneric = 1e-10

// Format is a parsed number format.
type Format struct {
	Whole         []Token
	Fractional    []Token
	Exponent      []Token
	IsExponential bool

	isFraction   bool
	isPercent    bool
	isGeneral    bool
	hasThousands bool
	skipNext     bool
	seenDecimal  bool

	denom       int64
	denomDigits int
}

// FmtType is the type of a format token.
//go:generate stringer -type=FmtType
type FmtType byte

// Format type constants.
const (
	FmtTypeLiteral FmtType = iota
	FmtTypeDigit
	FmtTypeDigitOpt
	FmtTypeComma
	FmtTypeDecimal
	FmtTypePercent
	FmtTypeDollar
	FmtTypeDigitOptThousands
	FmtTypeUnderscore
	FmtTypeDate
	FmtTypeTime
	FmtTypeFraction
	FmtTypeText
)

// Token is a format token in the Excel format string.
type Token struct {
	Type     FmtType
	Literal  byte
	DateTime string
}

// AddToken adds a format token to the format.
func (f *Format) AddToken(t FmtType, l []byte) {
	if f.skipNext {
		f.skipNext = false
		return
	}
	switch t {
	case FmtTypeDecimal:
		f.seenDecimal = true
	case FmtTypeUnderscore:
		f.skipNext = true
	case FmtTypeText:
		f.Whole = append(f.Whole, Token{Type: t})
	case FmtTypeDate, FmtTypeTime:
		f.Whole = append(f.Whole, Token{Type: t, DateTime: string(l)})
	case FmtTypePercent:
		f.isPercent = true
		t = FmtTypeLiteral
		l = []byte{'%'}
		fallthrough
	case FmtTypeDigitOpt:
		fallthrough
	case FmtTypeLiteral, FmtTypeDigit, FmtTypeDollar, FmtTypeComma:
		if l == nil {
			l = []byte{0}
		}
		for _, c := range l {
			if f.IsExponential {
				f.Exponent = append(f.Exponent, Token{Type: t, Literal: c})
			} else if !f.seenDecimal {
				f.Whole = append(f.Whole, Token{Type: t, Literal: c})
			} else {
				f.Fractional = append(f.Fractional, Token{Type: t, Literal: c})
			}
		}
	case FmtTypeDigitOptThousands:
		f.hasThousands = true
	case FmtTypeFraction:
		sp := strings.Split(string(l), "/")
		if len(sp) == 2 {
			f.isFraction = true
			f.denom, _ = strconv.ParseInt(sp[1], 10, 64)

			// also count the placeholder digits in the denominator for things
			// like ?/? and ?/??
			for _, c := range sp[1] {
				if c == '?' || c == '0' {
					f.denomDigits++
				}
			}
			// TODO: if anyone cares, parse and use the numerator format.
		}
	default:
		unioffice.Log("unsupported ph type in parse %v", t)
	}
}

// Number is used to format a number with a format string.  If the format
// string is empty, then General number formatting is used which attempts to mimic
// Excel's general formatting.
func Number(v float64, f string) string {
	if f == "" || f == "General" {
		return NumberGeneric(v)
	}
	fmts := Parse(f)
	if len(fmts) == 1 {
		return number(v, fmts[0], false)
	} else if len(fmts) > 1 && v < 0 {
		return number(v, fmts[1], true)
	} else if len(fmts) > 2 && v == 0 {
		return number(v, fmts[2], false)
	}
	return number(v, fmts[0], false)
}

// Value formats a value as a number or string depending on  if it appears to be
// a number or string.
func Value(v string, f string) string {
	if IsNumber(v) {
		v, _ := strconv.ParseFloat(v, 64)
		return Number(v, f)
	}
	return String(v, f)
}

// String returns the string formatted according to the type.  In format strings
// this is the fourth item, where '@' is used as a placeholder for text.
func String(v string, f string) string {
	fmts := Parse(f)
	var fm Format
	if len(fmts) == 1 {
		fm = fmts[0]
	} else if len(fmts) == 4 {
		fm = fmts[3]
	}
	fmtHasText := false
	for _, w := range fm.Whole {
		if w.Type == FmtTypeText {
			fmtHasText = true
		}
	}
	// no text placeholder in the format (@), so just return the original string
	if !fmtHasText {
		return v
	}
	b := bytes.Buffer{}
	for _, w := range fm.Whole {
		// these appear to be the only formats that matter for string
		switch w.Type {
		case FmtTypeLiteral:
			b.WriteByte(w.Literal)
		case FmtTypeText:
			b.WriteString(v)
		}
	}
	return b.String()
}

func reverse(b []byte) []byte {
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - 1 - i
		b[i], b[j] = b[j], b[i]
	}
	return b
}

func number(vOrig float64, f Format, isNeg bool) string {
	if f.isGeneral {
		return NumberGeneric(vOrig)
	}
	buf := make([]byte, 0, 20)
	wasNeg := math.Signbit(vOrig)
	v := math.Abs(vOrig)

	fractNum := int64(0)
	exp := int64(0)
	if f.IsExponential {
		for v >= 10 {
			exp++
			v /= 10
		}
		for v < 1 {
			exp--
			v *= 10
		}
	} else if f.isPercent {
		// percent symbol implies multiplying the value by 100
		v *= 100
	} else if f.isFraction {
		// if a denominator is not specified, we calculate it based on the number
		// of digits we are allowed
		if f.denom == 0 {
			maxDenom := math.Pow(10, float64(f.denomDigits))
			bestDenom, bestDenomError := 1.0, 1.0
			_ = bestDenom
			for i := 1.0; i < maxDenom; i++ {
				_, rem := math.Modf(v * float64(i))
				if rem < bestDenomError {
					bestDenomError = rem
					bestDenom = i
					if rem == 0 {
						break
					}
				}
			}
			f.denom = int64(bestDenom)
		}

		fractNum = int64(v*float64(f.denom) + 0.5)
		if len(f.Whole) > 0 && fractNum > f.denom {
			fractNum = int64(v*float64(f.denom)) % f.denom
			// subtract the portion we're displaying as a fraction
			v -= float64(fractNum) / float64(f.denom)
		} else {
			v -= float64(fractNum) / float64(f.denom)
			// This handles cases of '# ?/?' where we don't have a whole number
			// portion. The logic would print a '0' for the whole number portion
			// below, so we just clear f.Whole here to prevent anything from
			// being printed.
			if math.Abs(v) < 1 {
				isOnlyOpt := true
				for _, v := range f.Whole {
					// don't care about '#'
					if v.Type == FmtTypeDigitOpt {
						continue
					}
					// or spaces
					if v.Type == FmtTypeLiteral && v.Literal == ' ' {
						continue
					}
					isOnlyOpt = false
				}
				if isOnlyOpt {
					f.Whole = nil
				}
			}
		}
	}

	// round up now as this avoids rounding up on just the decimal portion which
	// is complicated due to a possible carry over into the whole number portion
	rndUp := 1
	for _, ph := range f.Fractional {
		if ph.Type == FmtTypeDigit || ph.Type == FmtTypeDigitOpt {
			rndUp++
		}
	}
	v += 5 * math.Pow10(-rndUp)

	// split into whole and decimal portions
	pre, post := math.Modf(v)
	buf = append(buf, formatWholeNumber(pre, vOrig, f)...)
	buf = append(buf, formatFractional(post, vOrig, f)...)
	buf = append(buf, formatExponential(exp, f)...)

	// fractions are special, the whole number portion is handled above (if
	// len(f.whole) > 0).  This is for the fractional portion, or in the case if
	// no whole portion, the numerator will be greater than the denominator.
	if f.isFraction {
		buf = strconv.AppendInt(buf, fractNum, 10)
		buf = append(buf, '/')
		buf = strconv.AppendInt(buf, f.denom, 10)
	}
	// if the number was negative, but this isn't a 'negative' format, then
	// we need to prepend a negative sign
	if !isNeg && wasNeg {
		return "-" + string(buf)
	}
	return string(buf)
}

func formatWholeNumber(pre, vOrig float64, f Format) []byte {
	if len(f.Whole) == 0 {
		return nil
	}
	epoch := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)
	t := epoch.Add(time.Duration(vOrig * float64(24*time.Hour)))
	t = asLocal(t)

	raw := strconv.AppendFloat(nil, pre, 'f', -1, 64)
	op := make([]byte, 0, len(raw))
	consumed := 0
	lastIdx := 1
lfor:
	for i := len(f.Whole) - 1; i >= 0; i-- {
		bidx := len(raw) - 1 - consumed
		ph := f.Whole[i]
		switch ph.Type {
		// '0' consumes a digit or prints a '0' if there is no digit
		case FmtTypeDigit:
			if bidx >= 0 {
				op = append(op, raw[bidx])
				consumed++
				lastIdx = i
			} else {
				op = append(op, '0')
			}
		// '#' consumes a digit or prints nothing
		case FmtTypeDigitOpt:
			if bidx >= 0 {
				op = append(op, raw[bidx])
				consumed++
				lastIdx = i
			} else {
				// we don't skip everything, just #/,/?. This is used so
				// that formats like (#,###) with '1' turn into '(1)' and
				// not '1)'
				for j := i; j >= 0; j-- {
					c := f.Whole[j]
					if c.Type == FmtTypeLiteral {
						op = append(op, c.Literal)
					}
				}
				break lfor
			}
		case FmtTypeDollar:
			for i := consumed; i < len(raw); i++ {
				op = append(op, raw[len(raw)-1-i])
				consumed++
			}
			op = append(op, '$')
		case FmtTypeComma:
			if !f.hasThousands {
				op = append(op, ',')
			}
		case FmtTypeLiteral:
			op = append(op, ph.Literal)
		case FmtTypeDate:
			op = append(op, reverse(dDate(t, ph.DateTime))...)
		case FmtTypeTime:
			op = append(op, reverse(dTime(t, vOrig, ph.DateTime))...)
		default:
			unioffice.Log("unsupported type in whole %v", ph)
		}
	}

	buf := reverse(op)
	// didn't consume all of the number characters, so insert the rest where
	// we were last inserting
	if consumed < len(raw) && (consumed != 0 || f.seenDecimal) {
		rem := len(raw) - consumed
		o := make([]byte, len(buf)+rem)
		copy(o, buf[0:lastIdx])
		copy(o[lastIdx:], raw[0:])
		copy(o[lastIdx+rem:], buf[lastIdx:])
		buf = o
	}
	if f.hasThousands {
		b := bytes.Buffer{}
		nonTerm := 0
		for i := len(buf) - 1; i >= 0; i-- {
			if !(buf[i] >= '0' && buf[i] <= '9') {
				nonTerm++
			} else {
				break
			}
		}
		for i := 0; i < len(buf); i++ {
			idx := (len(buf) - i - nonTerm)
			if idx%3 == 0 && idx != 0 && i != 0 {
				b.WriteByte(',')
			}
			b.WriteByte(buf[i])

		}
		buf = b.Bytes()
	}
	return buf
}

func formatFractional(post, vOrig float64, f Format) []byte {
	if len(f.Fractional) == 0 {
		return nil
	}

	raw := strconv.AppendFloat(nil, post, 'f', -1, 64)
	if len(raw) > 2 {
		raw = raw[2:] // skip the decimal portion (ie. '0.')
	} else {
		raw = nil
	}

	op := make([]byte, 0, len(raw))
	op = append(op, '.')
	consumed := 0
lforPost:
	for i := 0; i < len(f.Fractional); i++ {
		bidx := i
		ph := f.Fractional[i]
		switch ph.Type {
		// '0' consumes a digit or prints a '0' if there is no digit
		case FmtTypeDigit:
			if bidx < len(raw) {
				op = append(op, raw[bidx])
				consumed++
			} else {
				op = append(op, '0')
			}
		// '#' consumes a digit or prints nothing
		case FmtTypeDigitOpt:
			if bidx >= 0 {
				op = append(op, raw[bidx])
				consumed++
			} else {
				break lforPost
			}
		case FmtTypeLiteral:
			op = append(op, ph.Literal)
		default:
			unioffice.Log("unsupported type in fractional %v", ph)
		}
	}
	// remaining digits are truncated
	return op
}

func absi64(i int64) int64 {
	if i < 0 {
		return -i
	}
	return i
}

func formatExponential(exp int64, f Format) []byte {
	if !f.IsExponential || len(f.Exponent) == 0 {
		return nil
	}

	raw := strconv.AppendInt(nil, absi64(exp), 10)

	op := make([]byte, 0, len(raw)+2)
	op = append(op, 'E')
	if exp >= 0 {
		op = append(op, '+')
	} else {
		op = append(op, '-')
		exp *= -1
	}
	consumed := 0
lexfor:
	for i := len(f.Exponent) - 1; i >= 0; i-- {
		bidx := len(raw) - 1 - consumed
		ph := f.Exponent[i]
		switch ph.Type {
		// '0' consumes a digit or prints a '0' if there is no digit
		case FmtTypeDigit:
			if bidx >= 0 {
				op = append(op, raw[bidx])
				consumed++
			} else {
				op = append(op, '0')
			}
		// '#' consumes a digit or prints nothing
		case FmtTypeDigitOpt:
			if bidx >= 0 {
				op = append(op, raw[bidx])
				consumed++
			} else {
				for j := i; j >= 0; j-- {
					c := f.Exponent[j]
					if c.Type == FmtTypeLiteral {
						op = append(op, c.Literal)
					}
				}
				break lexfor
			}
		case FmtTypeLiteral:
			op = append(op, ph.Literal)
		default:
			unioffice.Log("unsupported type in exp %v", ph)
		}
	}
	// remaining non-consumed digits in the exponent
	if consumed < len(raw) {
		op = append(op, raw[len(raw)-consumed-1:consumed-1]...)
	}

	reverse(op[2:])
	return op
}

// NumberGeneric formats the number with the generic format which attemps to
// mimic Excel's general formatting.
func NumberGeneric(v float64) string {
	if math.Abs(v) >= maxGeneric || math.Abs(v) <= minGeneric && v != 0 {
		return formatExpNumberGeneric(v)
	}

	b := make([]byte, 0, 15)
	b = strconv.AppendFloat(b, v, 'f', -1, 64)
	if len(b) > 11 {
		ntlDigit := b[11] - '0'
		// if after truncating, do we need to round?
		if ntlDigit >= 5 && ntlDigit <= 9 {
			b[10]++
			b = b[0:11]
			b = performCarries(b)
		}
		b = b[0:11]
	} else if len(b) == 11 {
		// Excel rounds up a series of consecutive 9's if the total length is 11 digits
		if b[len(b)-1] == '9' {
			b[len(b)-1]++
			b = performCarries(b)
		}
	}

	b = trimTrailingZeros(b)
	return string(b)
}

func formatExpNumberGeneric(v float64) string {
	s1 := strconv.FormatFloat(v, 'E', -1, 64)
	s2 := strconv.FormatFloat(v, 'E', 5, 64)
	if len(s1) < len(s2) {
		return strconv.FormatFloat(v, 'E', 2, 64)
	}
	return s2
}

// trim any trailing zeros for numbers like '1.23000000' => '1.23' while not
// triming zeros on things like '10000'
func trimTrailingZeros(b []byte) []byte {
	end := len(b)
	sawDecimal := false
	sawDigit := false
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] == '0' && !sawDigit && !sawDecimal {
			end = i
		} else if b[i] == '.' {
			sawDecimal = true
		} else {
			sawDigit = true
		}
	}
	if sawDecimal && sawDigit {
		if b[end-1] == '.' {
			end--
		}
		return b[0:end]
	}
	return b
}

func performCarries(b []byte) []byte {
	// perform any required carries caused by rounding
	for i := len(b) - 1; i > 0; i-- {
		if b[i] == '9'+1 {
			b[i] = '0'
			if b[i-1] == '.' {
				i--
			}
			b[i-1]++
		}
	}
	// perform carry on initial digit by prepending a 1
	if b[0] == '9'+1 {
		b[0] = '0'
		copy(b[1:], b[0:])
		b[0] = '1'
	}
	return b
}

// dDate formats a time with an Excel format date string.
func dDate(t time.Time, f string) []byte {
	ret := []byte{}
	beg := 0
	for i := 0; i < len(f); i++ {
		var s string
		// split on '/'
		if f[i] == '/' {
			s = string(f[beg:i])
			beg = i + 1
		} else if i == len(f)-1 {
			s = string(f[beg : i+1])
		} else {
			continue
		}
		// Mon Jan 2 15:04:05 MST 2006
		switch s {
		case "yy":
			ret = t.AppendFormat(ret, "06")
		case "yyyy":
			ret = t.AppendFormat(ret, "2006")
		case "m":
			ret = t.AppendFormat(ret, "1")
		case "mm":
			ret = t.AppendFormat(ret, "01")
		case "mmm":
			ret = t.AppendFormat(ret, "Jan")
		case "mmmm":
			ret = t.AppendFormat(ret, "January")
		case "mmmmm":
			switch t.Month() {
			case time.January, time.July, time.June:
				ret = append(ret, 'J')
			case time.February:
				ret = append(ret, 'M')
			case time.March, time.May:
				ret = append(ret, 'M')
			case time.April, time.August:
				ret = append(ret, 'A')
			case time.September:
				ret = append(ret, 'S')
			case time.October:
				ret = append(ret, 'O')
			case time.November:
				ret = append(ret, 'N')
			case time.December:
				ret = append(ret, 'D')
			}
		case "d":
			ret = t.AppendFormat(ret, "2")
		case "dd":
			ret = t.AppendFormat(ret, "02")
		case "ddd":
			ret = t.AppendFormat(ret, "Mon")
		case "dddd":
			ret = t.AppendFormat(ret, "Monday")
		default:
			unioffice.Log("unsupported date format %s", s)
		}
		if f[i] == '/' {
			ret = append(ret, '/')
		}
	}
	return ret
}

// dTime formats a time with an Excel format time string.
func dTime(t time.Time, v float64, f string) []byte {
	ret := []byte{}
	beg := 0
	for i := 0; i < len(f); i++ {
		var s string
		// split on ':'
		if f[i] == ':' {
			s = string(f[beg:i])
			beg = i + 1
		} else if i == len(f)-1 {
			s = string(f[beg : i+1])
		} else {
			continue
		}

		// Mon Jan 2 15:04:05 MST 2006
		switch s {
		case "d":
			ret = t.AppendFormat(ret, "2")
		case "h":
			ret = t.AppendFormat(ret, "3")
		case "hh":
			ret = t.AppendFormat(ret, "15")
		case "m":
			ret = t.AppendFormat(ret, "4")
		case "mm":
			ret = t.AppendFormat(ret, "04")
		case "s":
			ret = t.Round(time.Second).AppendFormat(ret, "5")
		case "s.0":
			ret = t.Round(time.Second/10).AppendFormat(ret, "5.0")
		case "s.00":
			ret = t.Round(time.Second/100).AppendFormat(ret, "5.00")
		case "s.000":
			ret = t.Round(time.Second/1000).AppendFormat(ret, "5.000")
		case "ss":
			ret = t.Round(time.Second).AppendFormat(ret, "05")
		case "ss.0":
			ret = t.Round(time.Second/10).AppendFormat(ret, "05.0")
		case "ss.00":
			ret = t.Round(time.Second/100).AppendFormat(ret, "05.00")
		case "ss.000":
			ret = t.Round(time.Second/1000).AppendFormat(ret, "05.000")
		case "AM/PM":
			ret = t.AppendFormat(ret, "PM")
		case "[h]":
			ret = strconv.AppendInt(ret, int64(v*24), 10)
		case "[m]":
			ret = strconv.AppendInt(ret, int64(v*24*60), 10)
		case "[s]":
			ret = strconv.AppendInt(ret, int64(v*24*60*60), 10)
		case "":
		default:
			unioffice.Log("unsupported time format %s", s)
		}
		if f[i] == ':' {
			ret = append(ret, ':')
		}
	}
	return ret
}
