// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased on https://unidoc.io.

package formula

// noCache is a struct with collection of caching methods stubs intended for evaluators without cache.
type noCache struct{}

func (nc *noCache) SetCache(key string, value Result) {}

func (nc *noCache) GetFromCache(key string) (Result, bool) {
	return empty, false
}
