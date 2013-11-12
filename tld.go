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

func IsValidTLD(str string) bool {
	length := len(str)
	match := 0
	for _, tld := range tldList {
		if len(tld) != length {
			continue
		}
		match = 0
		for i := 0; i < length; i++ {
			if tld[i] == str[i] || tld[i]+32 == str[i] {
				match++
			}
		}
		if match == length {
			break
		}
	}
	return match == length
}
