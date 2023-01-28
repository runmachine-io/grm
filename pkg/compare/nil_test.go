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

package compare_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/anydotcloud/grm/pkg/compare"
)

func TestNilDifference(t *testing.T) {
	require := require.New(t)

	sampleString := ""
	var nullPtr *string
	nonNullPtr := &sampleString
	var nullMap map[string]string
	nonNullMap := make(map[string]string)
	var nullArray []string
	nonNullArray := make([]string, 0)
	var nullChan chan int
	nonNullChan := make(chan int, 0)

	require.False(compare.HasNilDifference(nil, nil))
	require.False(compare.HasNilDifference("a", "b"))
	require.True(compare.HasNilDifference(nil, "b"))
	require.True(compare.HasNilDifference("a", nil))

	//Pointer
	require.False(compare.HasNilDifference(nullPtr, nil))
	require.True(compare.HasNilDifference(nullPtr, nonNullPtr))

	//Map
	require.False(compare.HasNilDifference(nullMap, nil))
	require.True(compare.HasNilDifference(nullMap, nonNullMap))

	//Array or Slice
	require.False(compare.HasNilDifference(nullArray, nil))
	require.True(compare.HasNilDifference(nullArray, nonNullArray))

	//Chan
	require.False(compare.HasNilDifference(nullChan, nil))
	require.True(compare.HasNilDifference(nullChan, nonNullChan))

}
