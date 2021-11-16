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
