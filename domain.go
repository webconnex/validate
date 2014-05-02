// Copyright 2013 Webconnex, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validate

func readDomain(domain string) (int, int, bool) {
	count := len(domain)
	if count == 0 {
		return 0, 0, false
	}
	var last rune
	var fdot, dot int
	for i, c := range domain {
		if c == '/' ||
			(c >= ':' && c <= '@') ||
			(c >= '[' && c <= '`') ||
			(c <= ',' || c >= '{') {
			break // Invalid characters
		}
		if c == '.' {
			if count == 1 || // '.' is last
				last == 0 || // '.' is first
				last == '.' || // '.' after '.'
				last == '-' { // '.' is after '-'
				break
			}
			if fdot == 0 {
				fdot = i
			}
			dot = i
		} else if c == '-' {
			if count == 1 || // '-' is last
				last == '.' { // '-' is after '.'
				break
			}
		}
		last = c
		count--
	}
	return fdot, dot, count == 0
}

func IsValidDomain(domain string, reverse bool) bool {
	fdot, dot, valid := readDomain(domain)
	if valid {
		if dot == 0 { // We don't have a '.'
			valid = false
		} else if reverse {
			valid = IsValidTLD(domain[:fdot])
		} else {
			valid = IsValidTLD(domain[dot+1:])
		}
	}
	return valid
}

func IsValidSubdomain(subdomain string, allowSep bool) bool {
	_, dot, valid := readDomain(subdomain)
	if valid && dot > 0 != allowSep {
		valid = false
	}
	return valid
}
