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

import "fmt"

// A ContentTypeError describes an HTTP request responded with an unexpected
// content type.
type ContentTypeError struct {
	URL         string // URL of the request.
	Accept      string // Accept header of the request.
	ContentType string // Content-Type header of the response.
}

func (e ContentTypeError) Error() string {
	return fmt.Sprintf("httpext: %s responded with unexpected Content-Type \"%s\", Accept \"%s\"",
		e.URL, e.ContentType, e.Accept)
}
