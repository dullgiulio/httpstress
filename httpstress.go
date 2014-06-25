/* httpstress is a library for HTTP stress testing.
It launches one goroutine per concurrent connection and does not count successful attempts.

A CLI utility is avaliable at github.com/chillum/httpstress-go */
package httpstress

/* Copyright 2014 Chai Chillum

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License. */

import (
	"errors"
	"net/http"
	"regexp"
)

// Library version
const Version = "1.0.3"

/* httpstress.Test launches {conn} goroutines to fetch HTTP locations in {urls} list.
If {max} is more than {conn}, more goroutines will spawn as other are finished,
resulting in {max} queries (but no more than {conn} in every moment).
Returns map: {url}/{failed} or error (failed URL message). */
func Test(conn int, max int, urls []string) (results map[string]int, err error) {
	for _, i := range urls {
		if m, _ := regexp.MatchString("^http://", i); !m {
			err = errors.New("Not a HTTP URL: " + i)
			return
		}
	}

	results = make(map[string]int)
	failures := make(chan string)
	finished := make(chan bool)
	total := len(urls) - 1
	trans := &http.Transport{MaxIdleConnsPerHost: conn} // Use persistent connections.
	client := &http.Client{Transport: trans}
	n := 0
	i := 0

	go logger(failures, results)
	for ; i < conn; i++ { // Launch initial workers.
		go worker(&urls[n], failures, finished, client)

		if n < total {
			n++
		} else {
			n = 0
		}
	}
	for ; i < max; i++ { // Launch more workers as initial finish.
		if <-finished {
			go worker(&urls[n], failures, finished, client)

			if n < total {
				n++
			} else {
				n = 0
			}
		}
	}
	for i := 0; i < conn; i++ { // Wait for active workers.
		<-finished
	}
	return
}

func logger(failures <-chan string, results map[string]int) {
	for {
		select {
		case url := <-failures:
			results[url]++
		}
	}
}

func worker(url *string, failures chan<- string, finished chan<- bool, client *http.Client) {
	resp, err := client.Get(*url)
	if err != nil {
		failures <- *url
	}
	if resp != nil {
		resp.Body.Close()
	}
	finished <- true
}
