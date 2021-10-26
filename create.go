package fsfire

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// CreateFile Create new file, your can specific multiple options.
func CreateFile(filename string, data []byte, ops ...Option) (string, error) {
	options := options{
		extension: FileExtensionWithDocumentOfFileText,
		baseDir:   ".",
	}

	// Exists file extension, use it.
	extension := GetFileExtension(
		filename,
		WithSpecificFileExtensionContainsDot(false),
	)
	if extension != "" {
		options.extension = FileExtension(extension)
	}

	for _, o := range ops {
		o.apply(&options)
	}

	path, err := GetFilePathWithFSPath(options.baseDir, ops...)
	if err != nil {
		return "", err
	}

	filename = filepath.Join(path, CreateFilename(
		filename,
		options.extension.String(),
		ops...,
	))

	err = ioutil.WriteFile(filename, data, os.ModePerm)
	if err != nil {
		return "", err
	}

	return filename, nil
}
