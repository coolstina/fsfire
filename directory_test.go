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

func TestMustNotExistsMkdir(t *testing.T) {
	grids := []struct {
		path  string
		error error
	}{
		{
			path:  "test/data/tem/images/post",
			error: nil,
		},
		{
			path:  "test/data/tem/docs/images",
			error: nil,
		},
		{
			path:  "test/data/tem/article/comment",
			error: nil,
		},
	}

	for _, grid := range grids {
		err := NotExistsMkdir(grid.path)
		assert.NoError(t, err)
	}
}

func TestNotExistsMkdir(t *testing.T) {
	grids := []struct {
		path  string
		error error
	}{
		{
			path:  "test/data/create/dir/a/helloshaohua",
			error: nil,
		},
		{
			path:  "test/data/create/dir/b/helloshaohua",
			error: nil,
		},
		{
			path:  "test/data/create/dir/c/helloshaohua",
			error: nil,
		},
	}

	for _, grid := range grids {
		err := NotExistsMkdir(grid.path)
		assert.NoError(t, err)
	}
}

func TestReadDir(t *testing.T) {
	grids := []struct {
		path  string
		error error
	}{
		{
			path:  "test",
			error: nil,
		},
		{
			path:  ".",
			error: nil,
		},
	}

	for _, grid := range grids {
		actual, err := ReadDir(grid.path, nil)
		assert.NoError(t, err)
		assert.Greater(t, len(actual), 0)
	}
}
