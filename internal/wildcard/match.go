/*
 * MinIO Cloud Storage, (C) 2015, 2016 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package wildcard

// MatchSimple - finds whether the text matches/satisfies the pattern string.
// supports only '*' wildcard in the pattern.
// considers a file system path as a flat name space.
func MatchSimple(pattern, name string) bool {
	if pattern == "" {
		return name == pattern
	}
	if pattern == "*" {
		return true
	}
	rname := make([]rune, 0, len(name))
	rpattern := make([]rune, 0, len(pattern))
	for _, r := range name {
		rname = append(rname, r)
	}
	for _, r := range pattern {
		rpattern = append(rpattern, r)
	}
	simple := true // Does only wildcard '*' match.
	return deepMatchRune(rname, rpattern, simple)
}

// Match -  finds whether the text matches/satisfies the pattern string.
// supports  '*' and '?' wildcards in the pattern string.
// unlike path.Match(), considers a path as a flat name space while matching the pattern.
// The difference is illustrated in the example here https://play.golang.org/p/Ega9qgD4Qz .
func Match(pattern, name string) (matched bool) {
	if pattern == "" {
		return name == pattern
	}
	if pattern == "*" {
		return true
	}
	rname := make([]rune, 0, len(name))
	rpattern := make([]rune, 0, len(pattern))
	for _, r := range name {
		rname = append(rname, r)
	}
	for _, r := range pattern {
		rpattern = append(rpattern, r)
	}
	simple := false // Does extended wildcard '*' and '?' match.
	return deepMatchRune(rname, rpattern, simple)
}

func deepMatchRune(str, pattern []rune, simple bool) bool {
	for len(pattern) > 0 {
		switch pattern[0] {
		default:
			if len(str) == 0 || str[0] != pattern[0] {
				return false
			}
		case '?':
			if len(str) == 0 && !simple {
				return false
			}
		case '*':
			return deepMatchRune(str, pattern[1:], simple) ||
				(len(str) > 0 && deepMatchRune(str[1:], pattern, simple))
		}
		str = str[1:]
		pattern = pattern[1:]
	}
	return len(str) == 0 && len(pattern) == 0
}

// Index - finds a position of substring which matches/satisfies the pattern string.
// Supports  '*' and '?' wildcards in the pattern string.
// If nothing is found, it returns -1.
func Index(pattern, name string) (index int) {
	if pattern == "" || pattern == "*" {
		return 0
	}
	rname := make([]rune, 0, len(name))
	rpattern := make([]rune, 0, len(pattern))
	for _, r := range name {
		rname = append(rname, r)
	}
	for _, r := range pattern {
		rpattern = append(rpattern, r)
	}
	return deepIndexRune(rname, rpattern, 0)
}

// deepIndexRune is a recursive function which searches for a position of substring which matches given wildcard pattern.
func deepIndexRune(str, pattern []rune, position int) int {
	for len(pattern) > 0 {
		switch pattern[0] {
		default:
			if len(str) == 0 {
				return -1
			} // if pattern is not empty and string is empty, then nothing is found
			if str[0] != pattern[0] {
				return deepIndexRune(str[1:], pattern, position+1)
			} // if first characters don't match, try search from the next position
		case '?':
			if len(str) == 0 {
				return -1
			} // if string is empty, nothing is found. Otherwise as '?' means any character just throw both first string and first pattern symbols and continue search
		case '*':
			if len(str) == 0 {
				return -1
			} // if str is empty, nothing is found
			subIndex := deepIndexRune(str, pattern[1:], position) // throw the '*' and search with the rest of the pattern
			if subIndex != -1 {
				return position
			} else { // if nothing is found, continue from the next string position
				subIndex = deepIndexRune(str[1:], pattern, position)
				if subIndex != -1 {
					return position
				} else {
					return -1
				}
			}
		}
		str = str[1:]
		pattern = pattern[1:]
	}
	return position // if pattern becomes empty and -1 never returned, that means that the position of the substring is found
}
