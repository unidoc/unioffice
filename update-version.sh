#!/bin/bash


cat << __EOF > version.go
// Copyright `date +%Y` Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package gooxml

import "time"

// Release is the last release version of the software.
var ReleaseVersion = "v0.5000"

// ReleaseDate is the release date of the source code for licensing purposes.
var ReleaseDate = time.Date(`date +%Y`,`date +%_m`,`date +%_d`,0,0,0,0,time.UTC)

__EOF
goimports -w version.go
