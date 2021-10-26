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
