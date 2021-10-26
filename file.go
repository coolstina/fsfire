package fsfire

import (
	"os"
	"path/filepath"
	"strings"
)

// FileIsExists Check file is exists.
func FileIsExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// FileIsDir Check file is directory.
func FileIsDir(filename string) (bool, error) {
	stat, err := os.Stat(filename)
	if err != nil {
		return false, err
	}

	if stat.IsDir() {
		return true, nil
	}
	return false, nil
}

// LastDirName Gets filename the last directory name.
func LastDirName(filename string) string {
	sli := strings.Split(filepath.Dir(filename), string(os.PathSeparator))
	if len(sli) > 1 {
		return sli[len(sli)-1]
	}
	if len(sli) == 1 {
		return sli[0]
	}
	return ""
}

// CreateFilename CreateFilename generate filename.
func CreateFilename(filename, extension string, ops ...Option) string {
	options := options{
		trim: true,
	}

	for _, o := range ops {
		o.apply(&options)
	}

	if options.trim {
		ext := GetFileExtension(filename,
			WithSpecificFileExtensionContainsDot(true))
		if ext != "" {
			filename = strings.TrimSuffix(filename, ext)
		}
	}

	return strings.Join([]string{filename, extension}, GlobalSeparatorWithFileName)
}

// GetFileExtension Gets the file extension.
func GetFileExtension(filename string, ops ...Option) string {
	options := options{
		dot: false,
	}

	for _, o := range ops {
		o.apply(&options)
	}

	return func() string {
		ext := filepath.Ext(filename)
		if !options.dot {
			return strings.TrimPrefix(ext, ".")
		}
		return ext
	}()
}
