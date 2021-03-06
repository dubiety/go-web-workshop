// Copyright 2017 Google Inc. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to writing, software distributed
// under the License is distributed on a "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func doPut() {
	body := strings.NewReader("Hi there!! I'm D.")
	req, err := http.NewRequest("PUT", "https://http-methods.appspot.com/MyNameIs/Message", body)
	if err != nil {
		log.Fatalf("could not create request: %v", err)
	}

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("http request failed: %v", err)
	}

	fmt.Println(res.Status)
}

func main() {
	doPut()
}
