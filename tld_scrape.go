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

// +build ignore

package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"runtime"
	"strings"
)

func main() {
	// Get path to this file
	_, self, _, ok := runtime.Caller(0)
	if !ok {
		panic("Failed to get path")
	}

	// Scrape TLD list
	r, err := http.Get("http://data.iana.org/TLD/tlds-alpha-by-domain.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(r.Body)
	common := []string{"COM", "NET", "ORG"}
	list := []string{}
	list = append(list, common...)
Next:
	for scanner.Scan() {
		tld := string(scanner.Text())
		if tld[0] == '#' {
			continue
		}
		if tld[:2] == "XN" {
			continue
		}
		for _, ctld := range common {
			if tld == ctld {
				continue Next
			}
		}
		list = append(list, tld)
	}

	// Write tld_list.go
	buf := new(bytes.Buffer)
	buf.WriteString("package validate\n\n")
	buf.WriteString("var tldList = []string")
	buf.WriteByte('{')
	buf.WriteByte('"')
	buf.WriteString(strings.Join(list, `","`))
	buf.WriteByte('"')
	buf.WriteByte('}')
	ioutil.WriteFile(path.Dir(self)+"/tld_list.go", buf.Bytes(), 0644)
}
