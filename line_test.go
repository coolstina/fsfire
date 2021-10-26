package fsfire

import (
	"errors"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLines(t *testing.T) {
	grids := []struct {
		file     string
		expected int
	}{
		{
			file:     "./test/data/assets/textfiles/hello.txt",
			expected: 0,
		},
		{
			file:     "./test/data/assets/textfiles/helloworld.txt",
			expected: 12,
		},
		{
			file:     "./test/data/assets/textfiles/helloshaohua.txt",
			expected: 22,
		},
		{
			file:     "./test/data/assets/textfiles/tom.txt",
			expected: 2,
		},
	}

	for _, grid := range grids {
		data, err := ioutil.ReadFile(grid.file)
		assert.NoError(t, err)

		actual := Lines(data)
		assert.Equal(t, grid.expected, actual)
	}
}

func TestLinesForFile(t *testing.T) {
	grids := []struct {
		file     string
		expected int
		err      error
	}{
		{
			file:     "./test/data/assets/textfiles/hello.txt",
			expected: 0,
			err:      nil,
		},
		{
			file:     "./test/data/assets/textfiles/helloworld.txt",
			expected: 12,
			err:      nil,
		},
		{
			file:     "./test/data/assets/textfiles/helloshaohua.txt",
			expected: 22,
			err:      nil,
		},
		{
			file:     "./test/data/assets/textfiles/helloshaohua-not-exists.txt",
			expected: 0,
			err:      errors.New("no such file or directory"),
		},
	}

	for _, grid := range grids {
		actual, err := LinesForFile(grid.file)
		if err != nil {
			assert.Contains(t, err.Error(), grid.err.Error())
		}
		assert.Equal(t, grid.expected, actual)
	}
}

func TestCharacters(t *testing.T) {
	grids := []struct {
		file     string
		expected int
	}{
		{
			file:     "./test/data/assets/textfiles/hellochina.txt",
			expected: 110,
		},
		{
			file:     "./test/data/assets/textfiles/helloshaohua.txt",
			expected: 99,
		},
		{
			file:     "./test/data/assets/textfiles/helloworld.txt",
			expected: 31,
		},
		{
			file:     "./test/data/assets/textfiles/hello.txt",
			expected: 0,
		},
	}

	for _, grid := range grids {
		file, err := ioutil.ReadFile(grid.file)
		assert.NoError(t, err)

		actual := Characters(file)
		assert.Equal(t, grid.expected, actual)
	}
}

func TestCharactersForFile(t *testing.T) {
	grids := []struct {
		file     string
		expected int
		err      error
	}{
		{
			file:     "./test/data/assets/textfiles/hellochina.txt",
			expected: 110,
			err:      nil,
		},
		{
			file:     "./test/data/assets/textfiles/helloshaohua.txt",
			expected: 99,
			err:      nil,
		},
		{
			file:     "./test/data/assets/textfiles/helloworld.txt",
			expected: 31,
			err:      nil,
		},
		{
			file:     "./test/data/assets/textfiles/hello.txt",
			expected: 0,
			err:      nil,
		},
		{
			file:     "./test/data/assets/textfiles/hello-not-exists.txt",
			expected: 0,
			err:      errors.New("no such file or directory"),
		},
	}

	for _, grid := range grids {
		actual, err := CharactersForFile(grid.file)
		if err != nil {
			assert.Contains(t, err.Error(), grid.err.Error())
		}

		assert.Equal(t, grid.expected, actual)
	}
}
