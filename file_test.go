package fsfire

import (
	"embed"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileIsExists(t *testing.T) {
	grids := []struct {
		filename       string
		expectedExists bool
	}{
		{
			filename:       "create.go",
			expectedExists: true,
		},
		{
			filename:       "none.go",
			expectedExists: false,
		},
	}

	for _, grid := range grids {
		actual := FileIsExists(grid.filename)
		assert.Equal(t, grid.expectedExists, actual)
	}
}

func TestFileIsDir(t *testing.T) {
	grids := []struct {
		filename string
		error    error
		expected bool
	}{
		{
			filename: "test",
			error:    nil,
			expected: true,
		},
		{
			filename: "none.go",
			error:    errors.New("stat none.go: no such file or directory"),
			expected: false,
		},
		{
			filename: "create.go",
			error:    nil,
			expected: false,
		},
	}

	for _, grid := range grids {
		actual, err := FileIsDir(grid.filename)
		if err != nil {
			assert.Equal(t, grid.error.Error(), err.Error())
		} else {
			assert.Equal(t, grid.expected, actual)
		}
	}
}

func TestLastDirName(t *testing.T) {

	grids := []struct {
		filename string
		expected string
	}{
		{
			filename: "hello/world/helloshaohua/a.txt",
			expected: "helloshaohua",
		},
		{
			filename: "hello/world/a.txt",
			expected: "world",
		},
		{
			filename: "helloshaohua/a.txt",
			expected: "helloshaohua",
		},
		{
			filename: "a.txt",
			expected: ".",
		},
	}

	for _, grid := range grids {
		actual := LastDirName(grid.filename)
		assert.Equal(t, grid.expected, actual)
	}
}

func TestCreateFilename(t *testing.T) {
	grids := []struct {
		filename  string
		extension string
		trim      bool
		expected  string
	}{
		{
			filename:  "hello/world/helloshaohua/abc.txt",
			extension: "png",
			trim:      true,
			expected:  "hello/world/helloshaohua/abc.png",
		},
		{
			filename:  "hello/world/abc.jpeg",
			extension: "txt",
			trim:      false,
			expected:  "hello/world/abc.jpeg.txt",
		},
	}

	for _, grid := range grids {
		actual := CreateFilename(grid.filename, grid.extension,
			WithSpecificTrimOriginalFileExtension(grid.trim))
		assert.Equal(t, grid.expected, actual)
	}
}

func TestGetFileExtension(t *testing.T) {
	grids := []struct {
		filename string
		dot      bool
		expected string
	}{
		{
			filename: "hello/world/helloshaohua/abc.txt",
			dot:      true,
			expected: ".txt",
		},
		{
			filename: "hello/world/abc.jpeg",
			dot:      false,
			expected: "jpeg",
		},
		{
			filename: "helloshaohua/abc.png",
			dot:      false,
			expected: "png",
		},
		{
			filename: "abc",
			dot:      true,
			expected: "",
		},
	}

	for _, grid := range grids {
		actual := GetFileExtension(grid.filename,
			WithSpecificFileExtensionContainsDot(grid.dot))
		assert.Equal(t, grid.expected, actual)
	}
}

//go:embed test/data/embed
var efs embed.FS

func TestGetFileContentWithFS(t *testing.T) {
	grids := []struct {
		filename        string
		expectedError   error
		expectedContext []byte
	}{
		{
			filename:        "a.txt",
			expectedError:   fmt.Errorf(`open a.txt: file does not exist`),
			expectedContext: []byte(`a.txt file.`),
		},
		{
			filename:        "test/data/embed/a.txt",
			expectedError:   nil,
			expectedContext: []byte(`a.txt file.`),
		},
		{
			filename:        "test/data/embed/b.txt",
			expectedError:   nil,
			expectedContext: []byte(`b.txt file.`),
		},
		{
			filename:        "test/data/embed/c.txt",
			expectedError:   nil,
			expectedContext: []byte(`c.txt file.`),
		},
		{
			filename:        "test/data/embed/VERSION",
			expectedError:   nil,
			expectedContext: []byte(`v1.2.1`),
		},
	}

	for _, grid := range grids {
		actual, err := GetFileContentWithFS(efs, grid.filename)
		if err != nil {
			assert.Equal(t, grid.expectedError.Error(), err.Error())
		} else {
			assert.Equal(t, grid.expectedContext, actual)
		}
	}
}
