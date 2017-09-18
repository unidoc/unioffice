// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package format

import "time"

func asLocal(d time.Time) time.Time {
	d = d.UTC()
	return time.Date(d.Year(), d.Month(), d.Day(), d.Hour(),
		d.Minute(), d.Second(), d.Nanosecond(), time.Local)
}
func asUTC(d time.Time) time.Time {
	// Excel appears to interpret and serial dates in the local timezone, so
	// first ensure the time is converted internally.
	d = d.Local()

	// Then to avoid any daylight savings differences showing up between our
	// epoch and the current time, we 'cast' the time to UTC and later subtract
	// from the epoch in UTC.
	return time.Date(d.Year(), d.Month(), d.Day(), d.Hour(),
		d.Minute(), d.Second(), d.Nanosecond(), time.UTC)
}
