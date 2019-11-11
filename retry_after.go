// Copyright 2019 Miles Barr <milesbarr2@gmail.com>
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

package httpext

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// ParseRetryAfter parses a Retry-After header.
func ParseRetryAfter(s string) (time.Duration, error) {
	if s != "" {
		// Try parsing as an integer.
		i, err := strconv.ParseInt(s, 10, 64)
		if err == nil && i >= 0 {
			return time.Duration(i) * time.Second, nil
		}

		// Try parsing as a time.
		t, err := http.ParseTime(s)
		now := time.Now()
		if err == nil {
			if t.Before(now) {
				return 0, fmt.Errorf(
					"httpext.ParseRetryAfter: parsed time \"%v\" is in the past", t)
			}
			return t.Sub(now), nil
		}
	}
	return 0, fmt.Errorf("httpext.ParseRetryAfter: cannot parse \"%s\"", s)
}
