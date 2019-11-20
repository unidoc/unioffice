// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package unioffice

// MinGoVersion is used to cause a compile time error if unioffice is compiled with
// an older version of go.  Specifically it requires a feature in go 1.8
// regarding collecting all attributes from arbitrary xml used in decode
// unioffice.XSDAny.
const MinGoVersion = requires_go_18
