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

func IsValidEmail(email string) bool {
	var count = len(email)
	if count == 0 {
		return false
	}
	var last rune
	var at, dot int
	for i, c := range email {
		if c == ',' ||
			c == '/' ||
			c == '`' ||
			(c >= ':' && c <= '?') ||
			(c >= '[' && c <= '^') ||
			(c <= '*' || c >= '{') {
			break // Invalid characters
		}
		if c == '.' {
			if count == 1 || // '.' is last
				last == 0 || // '.' is first
				last == '.' || // '.' after '.'
				last == '@' || // '.' is after '@'
				(at > 0 && last == '-') { // '.' is after '-' in domain
				break
			}
			dot = i
		} else if c == '-' {
			if count == 1 || // '-' is last
				last == '@' || // '-' is after '@'
				(at > 0 && last == '.') { // '-' is after '.' in domain
				break
			}
		} else if c == '@' {
			if count == 1 || // '@' is last
				last == 0 || // '@' is first
				at > 0 { // '@' appears twice
				break
			}
			at = i
		} else if at > 0 && (c == '_' || c == '+') {
			break // Invalid characters in domain
		}
		last = c
		count--
	}
	return count == 0 && // Reached End
		at > 0 && // We have a '@'
		dot > at && // We have a '.' after '@'
		IsValidTLD(email[dot+1:])
}
