// Parts of this code modified from aws-controllers-k8s/pkg/compare
// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
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

package schema_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/anydotcloud/grm/pkg/types/resource/schema"
)

func TestJSON(t *testing.T) {
	require := require.New(t)

	ft := schema.FieldTypeList
	bytes, err := json.Marshal(&ft)
	require.Nil(err)
	require.Equal("\"list\"", string(bytes))

	var nft schema.FieldType
	err = json.Unmarshal(bytes, &nft)
	require.Nil(err)
	require.Equal(schema.FieldTypeList, nft)
}
