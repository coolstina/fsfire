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
	"errors"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanBlankLines(t *testing.T) {
	grids := []struct {
		file     string
		expected int
		err      error
	}{
		{
			file:     "test/data/assets/textfiles//hellochina.txt",
			expected: 16,
			err:      nil,
		},
		{
			file:     "test/data/assets/textfiles//helloshaohua.txt",
			expected: 11,
			err:      nil,
		},
		{
			file:     "test/data/assets/textfiles//helloworld.txt",
			expected: 6,
			err:      nil,
		},
		{
			file:     "test/data/assets/textfiles//hello.txt",
			expected: 0,
			err:      nil,
		},
		{
			file:     "test/data/assets/textfiles//hello-not-exists.txt",
			expected: 0,
			err:      errors.New("no such file or directory"),
		},
	}

	for _, grid := range grids {

		data, err := ioutil.ReadFile(grid.file)
		if err != nil {
			assert.Contains(t, err.Error(), grid.err.Error())
		}

		actual := Lines(CleanBlankLines(data))
		assert.Equal(t, grid.expected, actual)
	}
}

func TestCleanBlankLinesForFile(t *testing.T) {
	grids := []struct {
		file     string
		expected int
		err      error
	}{
		{
			file:     "test/data/assets/textfiles//hellochina.txt",
			expected: 16,
			err:      nil,
		},
		{
			file:     "test/data/assets/textfiles//helloshaohua.txt",
			expected: 11,
			err:      nil,
		},
		{
			file:     "test/data/assets/textfiles//helloworld.txt",
			expected: 6,
			err:      nil,
		},
		{
			file:     "test/data/assets/textfiles//hello.txt",
			expected: 0,
			err:      nil,
		},
		{
			file:     "test/data/assets/textfiles//hello-not-exists.txt",
			expected: 0,
			err:      errors.New("no such file or directory"),
		},
	}

	for _, grid := range grids {

		data, err := CleanBlankLinesForFile(grid.file)
		if err != nil {
			assert.Contains(t, err.Error(), grid.err.Error())
		}

		assert.Equal(t, grid.expected, Lines(data))
	}
}
