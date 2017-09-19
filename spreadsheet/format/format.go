// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package format

import (
	"bytes"
	"log"
	"math"
	"strconv"
	"time"
)

const maxGeneric = 1e11
const minGeneric = 1e-10

type Format struct {
	Whole         []PlaceHolder
	Fractional    []PlaceHolder
	Exponent      []PlaceHolder
	IsExponential bool

	isPercent    bool
	isGeneral    bool
	hasThousands bool
	skipNext     bool
	seenDecimal  bool
}

//go:generate stringer -type=FmtType
type FmtType byte

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
)

type PlaceHolder struct {
	Type     FmtType
	Literal  byte
	DateTime string
}

func (f *Format) AddPlaceholder(t FmtType, l []byte) {
	if f.skipNext {
		f.skipNext = false
		return
	}
	switch t {
	case FmtTypeDecimal:
		f.seenDecimal = true
	case FmtTypeUnderscore:
		f.skipNext = true
	case FmtTypeDate, FmtTypeTime:
		f.Whole = append(f.Whole, PlaceHolder{Type: t, DateTime: string(l)})
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
				f.Exponent = append(f.Exponent, PlaceHolder{Type: t, Literal: c})
			} else if !f.seenDecimal {
				f.Whole = append(f.Whole, PlaceHolder{Type: t, Literal: c})
			} else {
				f.Fractional = append(f.Fractional, PlaceHolder{Type: t, Literal: c})
			}
		}
	case FmtTypeDigitOptThousands:
		f.hasThousands = true
	default:
		log.Printf("unsupported ph type in parse %s", t)
	}
}

// Number is used to format a number with a format string.  If the format
// string is empty, then General number formatting is used which attempts to mimic
// Excel's general formatting.
func Number(v float64, f string) string {
	if f == "" {
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

func String(v string, f string) string {
	return v
}

func reverse(b []byte) []byte {
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - 1 - i
		b[i], b[j] = b[j], b[i]
	}
	return b
}

func number(vOrig float64, f Format, isNeg bool) string {
	epoch := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)
	t := epoch.Add(time.Duration(vOrig * float64(24*time.Hour)))
	t = asLocal(t)
	if f.isGeneral {
		return NumberGeneric(vOrig)
	}
	buf := make([]byte, 0, 20)
	wasNeg := math.Signbit(vOrig)
	v := math.Abs(vOrig)

	// percent symbol implies multiplying the value by 100
	if f.isPercent {
		v *= 100
	}
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

	if len(f.Whole) > 0 {
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
				log.Fatalf("unsupported type in whole %v", ph)
			}
		}
		buf = append(buf, reverse(op)...)

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
	}

	if len(f.Fractional) != 0 {
		buf = append(buf, '.')
		raw := strconv.AppendFloat(nil, post, 'f', -1, 64)
		if len(raw) > 2 {
			raw = raw[2:] // skip the decimal portion (ie. '0.')
		} else {
			raw = nil
		}
		op := make([]byte, 0, len(raw))
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
				log.Fatalf("unsupported type in fractional %v", ph)
			}
		}
		// remaining digits are truncated
		buf = append(buf, op...)
	}

	if f.IsExponential {
		if len(f.Exponent) > 0 {
			buf = append(buf, 'E')
			if exp >= 0 {
				buf = append(buf, '+')
			} else {
				buf = append(buf, '-')
				exp *= -1
			}
			raw := strconv.AppendInt(nil, exp, 10)
			op := make([]byte, 0, len(raw))
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
					log.Printf("unsupported type in exp %v", ph)
				}
			}
			// remaining non-consumed digits in the exponent
			if consumed < len(raw) {
				op = append(op, raw[len(raw)-consumed-1:consumed-1]...)
			}
			buf = append(buf, reverse(op)...)

		}
	}
	// if the number was negative, but this isn't a 'negative' format, then
	// we need to prepend a negative sign
	if !isNeg && wasNeg {
		return "-" + string(buf)
	}
	return string(buf)
}

// NumberGeneric formats the number with the generic format which attemps to
// mimic Excel's general formatting.
func NumberGeneric(v float64) string {
	if math.Abs(v) >= maxGeneric || math.Abs(v) <= minGeneric {
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
			log.Fatalf("unsupported date format %s", s)
		}
		if f[i] == '/' {
			ret = append(ret, '/')
		}
	}
	return ret
}

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
			log.Fatalf("unsupported time format %s", s)
		}
		if f[i] == ':' {
			ret = append(ret, ':')
		}
	}
	return ret
}
