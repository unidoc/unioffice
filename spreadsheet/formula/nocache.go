// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package formula

// noCache is a struct with collection of caching methods stubs intended for evaluators without cache.
type noCache struct{}

func (nc *noCache) SetCache(key string, value Result) {}

func (nc *noCache) GetFromCache(key string) (Result, bool) {
	return empty, false
}
