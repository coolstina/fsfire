// Copyright 2021 helloshaohua <wu.shaohua@foxmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fsfire

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateZIPFileWithFilename(t *testing.T) {
	grids := []struct {
		source   string
		expected string
	}{
		{
			source:   "./test",
			expected: "test.zip",
		},
		{
			source:   "./zip.go",
			expected: "zip.zip",
		},
	}

	for _, grid := range grids {
		actual, err := CreateZIPFileWithFilename(grid.source)
		assert.NoError(t, err)
		assert.Equal(t, grid.expected, actual)
	}
}
