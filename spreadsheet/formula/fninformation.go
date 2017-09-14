// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

func init() {
	RegisterFunction("NA", NA)
}

// NA is an implementation of the Excel NA() function that just returns the #N/A! error.
func NA(args []Result) Result {
	if len(args) != 0 {
		MakeErrorResult("NA() accepts no arguments")
	}
	return MakeErrorResultType(ErrorTypeNA, "")
}
