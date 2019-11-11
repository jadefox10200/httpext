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

// IsInformationalStatus returns whether a status code is informational
// (100-199).
func IsInformationalStatus(code int) bool {
	return 100 <= code && code < 200
}

// IsSuccessStatus returns whether a status code indicates a success (200-299).
func IsSuccessStatus(code int) bool {
	return 200 <= code && code < 300
}

// IsRedirectStatus returns whether a status code indicates a redirect
// (300-399).
func IsRedirectStatus(code int) bool {
	return 300 <= code && code < 400
}

// IsClientErrorStatus returns whether a status code indicates a client error
// (400-499).
func IsClientErrorStatus(code int) bool {
	return 400 <= code && code < 500
}

// IsServerErrorStatus returns whether a status code indicates a server error
// (500-599).
func IsServerErrorStatus(code int) bool {
	return 500 <= code && code < 600
}

// IsErrorStatus returns whether a status code indicates an error (400-599).
func IsErrorStatus(code int) bool {
	return 400 <= code && code < 600
}
