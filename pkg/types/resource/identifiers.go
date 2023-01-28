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

package resource

import "github.com/anydotcloud/grm/pkg/types/resource/schema"

// Identifiers contains all the information callers would need to identify a
// particular resource
type Identifiers interface {
	// ValuesIter returns a slice, ordered by efficiency of fetch operation,
	// of maps, keyed by identifying field, of identifying field values.
	ValuesIter() []map[schema.Field]string
	// ValuesBy returns a slice of strings representing the values of
	// supplied identifying Fields
	ValuesBy(...schema.Field) []string
}
