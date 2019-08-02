// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package unioffice

import (
	"log"
)

// Log is used to log content from within the library.  The intent is to use
// logging sparingly, preferring to return an error.  At the very least this
// allows redirecting logs to somewhere more appropriate than stdout.
var Log = log.Printf

// DisableLogging sets the Log function to a no-op so that any log messages are
// silently discarded.
func DisableLogging() {
	Log = func(string, ...interface{}) {}
}
