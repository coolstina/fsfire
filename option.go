package fsfire

type Option interface {
	apply(*options)
}

type options struct {
	extension          FileExtension
	dot                bool
	trim               bool
	path               FSPath
	baseDir            string
	containsMarkString bool
}

type optionFunc func(*options)

func (o optionFunc) apply(ops *options) {
	o(ops)
}

// WithSpecificFSPath Override the original value by specifying filepath.
func WithSpecificFSPath(path FSPath) Option {
	return optionFunc(func(ops *options) {
		ops.path = path
	})
}

// WithSpecificBaseDir Override the original value by specifying base directory.
func WithSpecificBaseDir(baseDir string) Option {
	return optionFunc(func(ops *options) {
		ops.baseDir = baseDir
	})
}

// WithSpecificFSPath Override the original value by specifying filepath.
func WithSpecificFileExtension(extension FileExtension) Option {
	return optionFunc(func(ops *options) {
		ops.extension = extension
	})
}

// WithSpecificFileExtensionContainsDot Override the original value by specifying dot.
func WithSpecificFileExtensionContainsDot(containsdot bool) Option {
	return optionFunc(func(ops *options) {
		ops.dot = containsdot
	})
}

// WithSpecificFileTrimOriginalExtension Override the original value by specifying dot.
func WithSpecificTrimOriginalFileExtension(trim bool) Option {
	return optionFunc(func(ops *options) {
		ops.trim = trim
	})
}

// WithSpecificContainsMarkString Override the original value by containsMarkString mark string.
func WithSpecificContainsMarkString(contains bool) Option {
	return optionFunc(func(ops *options) {
		ops.containsMarkString = contains
	})
}
