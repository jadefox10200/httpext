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
	"net/http"
	"time"

	"github.com/tradyfinance/backoff"
)

// A LazyRateLimiter implements rate limiting for an http.RoundTripper object.
// Rate limiting start when the server responds with 429 Too Many Requests.
type LazyRateLimiter struct {
	transport     http.RoundTripper
	backoffPolicy backoff.Policy
}

// NewLazyRateLimiter returns a new LazyRateLimiter given a transport and
// backoff policy. The transport will use http.DefaultTransport when nil. The
// backoff policy will use backoff.Default when nil.
func NewLazyRateLimiter(transport http.RoundTripper, backoffPolicy backoff.Policy) LazyRateLimiter {
	// Use http.DefaultTransport when nil.
	if transport == nil {
		transport = http.DefaultTransport
	}

	// Use backoff.Default when nil.
	if backoffPolicy == nil {
		backoffPolicy = backoff.Default()
	}

	return LazyRateLimiter{transport, backoffPolicy}
}

// RoundTrip implemented the http.RoundTripper interface.
func (lrl LazyRateLimiter) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := lrl.transport.RoundTrip(req)

	// Decrease the backoff policy delay if the server does not respond with 429
	// Too Many Requests.
	if resp.StatusCode != http.StatusTooManyRequests {
		lrl.backoffPolicy.Decrease()
	} else {
		// Repeat until the server does not respond with 429 Too Many Requests.
		for resp.StatusCode == http.StatusTooManyRequests {
			if err := resp.Body.Close(); err != nil {
				return nil, err
			}

			// Rate limit based on the Retry-After header or otherwise use the
			// backoff policy.
			d, err := ParseRetryAfter(resp.Header.Get("Retry-After"))
			if err != nil {
				d = lrl.backoffPolicy.Increase()
			}
			time.Sleep(d)

			resp, err = lrl.transport.RoundTrip(req)
			if err != nil {
				return nil, err
			}
		}
	}

	return resp, err
}
