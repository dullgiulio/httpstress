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
	up := []string{"https://google.com", "http://google.com"} // These URLs should pass.
	down := []string{"http://test.invalid"}                   // This should fail.
	invalid := []string{"ftp://localhost"}                    // Error. Non HTTP/HTTPS URL.

	if _, err := Test(1, 1, invalid); err == nil {
		t.Errorf("%s is ok (should be an error)", invalid)
	}

	if err, _ := Test(1, 1, up); len(err) > 0 {
		t.Errorf("%s is down (should be up)", up)
	}

	if err, _ := Test(1, 1, down); len(err) == 0 {
		t.Errorf("%s is up (should be down)", down)
	}
}
