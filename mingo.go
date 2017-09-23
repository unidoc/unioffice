// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package gooxml

// MinGoVersion is used to cause a compile time error if gooxml is compiled with
// an older version of go.  Specifically it requires a feature in go 1.8
// regarding collecting all attributes from arbitrary xml used in decode
// gooxml.XSDAny.
const MinGoVersion = requires_go_18
