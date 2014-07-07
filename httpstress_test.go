/* Test the httpstress.Test() function. */
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

import "testing"

func TestTest(t *testing.T) {
	up := make([]string, 2)
	up[0] = "https://google.com"    // These URLs should pass.
	up[1] = "http://google.com"
	down := make([]string, 1)
	down[0] = "http://test.invalid" // This should fail.
	error := make([]string, 1)
	error[0] = "ftp://localhost"    // Error. Non HTTP/HTTPS URL.

	if _, err := Test(1, 1, error); err == nil {
		t.Errorf("%s is ok (should be an error)", error)
	}

	if err, _ := Test(1, 1, up); len(err) > 0 {
		t.Errorf("%s is down (should be up)", up)
	}

	if err, _ := Test(1, 1, down); len(err) == 0 {
		t.Errorf("%s is up (should be down)", down)
	}
}
