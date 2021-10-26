package fsfire

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanSpecificContentLines(t *testing.T) {
	grids := []struct {
		file     string
		fineStr  string
		expected int
	}{
		{
			file:     "test/data/assets/textfiles/helloworld.txt",
			fineStr:  "山东",
			expected: 5,
		},
		{
			file:     "test/data/assets/textfiles/helloworld.txt",
			fineStr:  "中国",
			expected: 5,
		},
		{
			file:     "test/data/assets/textfiles/helloworld.txt",
			fineStr:  "hello",
			expected: 3,
		},
		{
			file:     "test/data/assets/textfiles/helloworld.txt",
			fineStr:  "world",
			expected: 5,
		},
	}

	for _, grid := range grids {
		data, err := ioutil.ReadFile(grid.file)
		assert.NoError(t, err)

		// Clean blank lines.
		data = CleanBlankLines(data)

		// Clean specific content lines
		data = CleanSpecificContentLines(data, grid.fineStr)
		assert.Equal(t, grid.expected, Lines(data))
	}
}

func TestCleanSpecificContentLinesForFile(t *testing.T) {
	grids := []struct {
		file     string
		fineStr  string
		expected int
	}{
		{
			file:     "test/data/assets/textfiles/helloworld.txt",
			fineStr:  "山东",
			expected: 11,
		},
		{
			file:     "test/data/assets/textfiles/helloworld.txt",
			fineStr:  "中国",
			expected: 11,
		},
		{
			file:     "test/data/assets/textfiles/helloworld.txt",
			fineStr:  "hello",
			expected: 9,
		},
		{
			file:     "test/data/assets/textfiles/helloworld.txt",
			fineStr:  "world",
			expected: 11,
		},
	}

	for _, grid := range grids {
		data, err := CleanSpecificContentLinesForFile(grid.file, grid.fineStr)
		assert.NoError(t, err)
		assert.Equal(t, grid.expected, Lines(data))
	}
}
