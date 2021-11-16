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
