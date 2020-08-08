// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

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
