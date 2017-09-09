// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package measurement

// Distance represents a distance and is automatically converted
// to the units needed internally in the various ECMA 376 formats.
type Distance float64

// Constants for various distance units
const (
	Zero           Distance = 0
	Point                   = 1
	Pixel72                 = 1.0 / 72.0 * Inch
	Pixel96                 = 1.0 / 96.0 * Inch
	HalfPoint               = 1.0 / 2.0 * Point
	Character               = 7 * Point
	Millimeter              = 2.83465 * Point
	Centimeter              = 10 * Millimeter
	Inch                    = 72 * Point
	Foot                    = 12 * Inch
	Twips                   = 1.0 / 20.0 * Point
	EMU                     = 1.0 / 914400.0 * Inch
	HundredthPoint          = 1 / 100.0
	Dxa                     = Twips
)
