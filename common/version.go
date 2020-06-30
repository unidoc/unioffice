/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

// Package common contains common properties used by the subpackages.
package common

import (
	"time"
)

const releaseYear = 2020
const releaseMonth = 6
const releaseDay = 30
const releaseHour = 16
const releaseMin = 10

// Version holds version information, when bumping this make sure to bump the released at stamp also.
const Version = "1.4.0"

var ReleasedAt = time.Date(releaseYear, releaseMonth, releaseDay, releaseHour, releaseMin, 0, 0, time.UTC)
