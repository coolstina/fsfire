package fsfire

import "path/filepath"

// AbsPath Checks whether a directory is an absolute
// path and can be created by creating a path that does not exist
func AbsPath(path string, create bool) (string, error) {
	var err error

	if !filepath.IsAbs(path) {
		path, err = filepath.Abs(path)
		if err != nil {
			return "", err
		}
	}

	if create {
		err = NotExistsMkdirAll(path)
		if err != nil {
			return "", err
		}
	}

	return path, nil
}
