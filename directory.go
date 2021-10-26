package fsfire

import (
	"os"
	"path/filepath"
)

// RuntimeMasterDir returns the root path corresponding to the current directory.
func RuntimeMasterDir() string {
	base, _ := os.Getwd()
	return base
}

func MustGetFilePathWithFSPath(baseDir string, ops ...Option) string {
	options := options{
		path: GlobalDefaultDir,
	}

	for _, o := range ops {
		o.apply(&options)
	}

	path := filepath.Join(RuntimeMasterDir(), baseDir, options.path.String())
	if err := NotExistsMkdir(path); err != nil {
		panic(err)
	}

	return path
}

func GetFilePathWithFSPath(baseDir string, ops ...Option) (string, error) {
	options := options{
		path: GlobalDefaultDir,
	}

	for _, o := range ops {
		o.apply(&options)
	}

	path := filepath.Join(RuntimeMasterDir(), baseDir, options.path.String())
	if err := NotExistsMkdir(path); err != nil {
		return "", err
	}

	return path, nil
}

// NotExistsMkdir detects whether the target folder exists.
// and if it does not, create the folder.
func NotExistsMkdir(path string) error {
	if exists := IsNotExists(path); exists {
		if err := MkdirAll(path); err != nil {
			return err
		}
	}
	return nil
}

// IsNotExists check whether the target file or directory exists.
// Return FALSE if the file exists, TRUE otherwise.
func IsNotExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	}
	return false
}

// MkdirAll creates a directory named path.
func MkdirAll(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
