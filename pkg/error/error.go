// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package error

import (
	"errors"
	"fmt"
)

// UnknownField indicates there was a bad field path supplied for a Resource's
// SetAt method call.
var (
	UnknownField = errors.New("Unknown field")
)

// UnknownFieldAtPath returns an UnknownField error for a supplied field path
// string
func UnknownFieldAtPath(p string) error {
	return fmt.Errorf("%w: %s", UnknownField, p)
}
