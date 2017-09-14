// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package formula

import (
	"log"
	"sync"
)

type Function func(args []Result) Result

var regLock sync.Mutex
var registered = map[string]Function{}

func RegisterFunction(name string, fn Function) {
	regLock.Lock()
	defer regLock.Unlock()
	if _, ok := registered[name]; ok {
		log.Printf("duplicate registration of function %s", name)
	}
	registered[name] = fn

}

func LookupFunction(name string) Function {
	regLock.Lock()
	defer regLock.Unlock()
	if fn, ok := registered[name]; ok {
		return fn
	}
	log.Printf("unknown function %s", name)
	return nil
}
