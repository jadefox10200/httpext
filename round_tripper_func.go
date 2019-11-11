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

import "net/http"

// A RoundTripperFunc is a function used as the RoundTrip method for the
// http.RoundTripper interface.
type RoundTripperFunc func(req *http.Request) (*http.Response, error)

// RoundTrip implements the http.RoundTripper interface.
func (f RoundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

// WithTransportFunc returns a copy of a client that uses a given
// RoundTripperFunc as the transport. The client will default to
// http.DefaultClient when nil.
func WithTransportFunc(c *http.Client, f RoundTripperFunc) *http.Client {
	if c == nil {
		c = http.DefaultClient
	}
	clientWithTransport := *c
	clientWithTransport.Transport = f
	return &clientWithTransport
}
