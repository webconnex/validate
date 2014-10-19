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
	"sort"
	"strings"
)

func main() {
	// Get path to this file
	_, self, _, ok := runtime.Caller(0)
	if !ok {
		panic("Failed to get path")
	}

	// Scrape TLD list
	url := "http://data.iana.org/TLD/tlds-alpha-by-domain.txt"
	r, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(r.Body)
	comments := []string{}
	tldList := []string{}
	for scanner.Scan() {
		line := string(scanner.Text())
		if line[0] == '#' {
			comments = append(comments, strings.TrimLeft(line, "# "))
			continue
		}
		if len(line) > 4 && line[:4] == "XN--" {
			continue
		}
		tldList = append(tldList, line)
	}
	sort.Sort(sort.StringSlice(tldList))

	// Write tld_list.go
	buf := new(bytes.Buffer)
	buf.WriteString("package validate\n\n")
	buf.WriteString("// ")
	buf.WriteString(url)
	buf.WriteByte('\n')
	if len(comments) > 0 {
		buf.WriteString("// ")
		buf.WriteString(comments[0])
		buf.WriteByte('\n')
	}
	buf.WriteByte('\n')
	buf.WriteString("var tldList = []string{\n\t\"")
	buf.WriteString(strings.Join(tldList, "\",\n\t\""))
	buf.WriteString("\",\n}")
	ioutil.WriteFile(path.Dir(self)+"/tld_list.go", buf.Bytes(), 0644)
}
