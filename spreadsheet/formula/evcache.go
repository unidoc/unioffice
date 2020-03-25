// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

import "sync"

// evCache is a struct with collection of caching methods intended for add cache support to evaluators.
type evCache struct {
	cache map[string]Result
	lock  *sync.Mutex
}

func newEvCache() evCache {
	ev := evCache{}
	ev.cache = make(map[string]Result)
	ev.lock = &sync.Mutex{}
	return ev
}

func (ec *evCache) SetCache(key string, value Result) {
	ec.lock.Lock()
	ec.cache[key] = value
	ec.lock.Unlock()
}

func (ec *evCache) GetFromCache(key string) (Result, bool) {
	ec.lock.Lock()
	result, found := ec.cache[key]
	ec.lock.Unlock()
	return result, found
}
