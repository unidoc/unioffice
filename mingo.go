// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package unioffice

// MinGoVersion is used to cause a compile time error if unioffice is compiled with
// an older version of go.  Specifically it requires a feature in go 1.8
// regarding collecting all attributes from arbitrary xml used in decode
// unioffice.XSDAny.
const MinGoVersion = requires_go_18
