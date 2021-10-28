package fsfire

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInSlice(t *testing.T) {
	grids := []struct {
		source   []interface{}
		find     interface{}
		expected bool
	}{
		{
			source:   []interface{}{"hello", "world"},
			find:     "hello",
			expected: true,
		},
		{
			source:   []interface{}{"hello", "world"},
			find:     "not_exists",
			expected: false,
		},
		{
			source:   []interface{}{"hello", "world"},
			find:     5,
			expected: false,
		},
	}

	for _, grid := range grids {
		actual := InSlice(grid.source, grid.find)
		assert.Equal(t, actual, grid.expected)
	}
}
