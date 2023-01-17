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

package manager

import (
	"github.com/jaypipes/grm/pkg/types/credentials"
	"github.com/jaypipes/grm/pkg/types/resource/schema"
)

// Factory returns a Manager that can be used to manage cloud resources for a
// particular Resource Kind and cloud credentials
type Factory interface {
	// Manager returns a Manager that manages cloud resources on behalf of a
	// particular set of cloud credentials
	Manager(
		schema.Kind,
		credentials.Credentials,
	) (Manager, error)
}
