// Copyright 2021 helloshaohua <wu.shaohua@foxmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fsfire

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// CreateFilename Create the filename.
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

// IsFile Checks whether the specified file name is a file.
func IsFile(filename string) (bool, error) {
	stat, err := os.Stat(filename)
	if err != nil {
		return false, err
	}

	if stat.IsDir() {
		return false, nil
	}

	return true, nil
}

// Touch Create file, if parent directory not exists, then create.
func Touch(filename string) error {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
		if err != nil {
			return err
		}

		f, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer f.Close()
	} else {
		now := time.Now().Local()
		if err := os.Chtimes(filename, now, now); err != nil {
			return err
		}
	}

	return nil
}

// FilenameTrimPrefix Trim the filename prefix.
func FilenameTrimPrefix(filename, prefix string) string {
	return strings.TrimPrefix(filename, prefix)
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

// GetFileOrDirectoryName Gets the file or directory name.
func GetFileOrDirectoryName(path string) (string, bool, error) {
	stat, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return "", false, err
	}

	if stat.IsDir() {
		return filepath.Base(strings.TrimRight(path, string(os.PathSeparator))), true, nil
	}

	str := []byte(filepath.Base(path))
	return string(str[0:bytes.LastIndexAny(str, ".")]), false, nil
}

// GetFileContentWithEmbedFS Get file content with embed filesystem.
func GetFileContentWithEmbedFS(fs embed.FS, filename string) ([]byte, error) {
	data, err := fs.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetFileContentWithStringSlice Get the file contents as a string for slice.
func GetFileContentWithStringSlice(filename string, ops ...Option) ([]string, error) {
	options := &options{}

	for _, o := range ops {
		o.apply(options)
	}

	open, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer open.Close()

	var content = make([]string, 0, 256)
	var reader = bufio.NewReader(open)
	var prefix = bytes.Buffer{}
	for {
		data, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		if options.ignoreBlankLine {
			if len(bytes.TrimSpace(data)) == 0 {
				continue
			}
		}

		if isPrefix {
			prefix.Write(data)
			continue
		} else {
			prefix.Write(data)
		}

		content = append(content, prefix.String())
		prefix.Reset()
	}

	return content, nil
}

// GetContentWithRegularExpression Get the string content into slice's string, using a regular expression.
func GetContentWithRegularExpression(str, pattern string, ops ...Option) ([]string, error) {
	if str == "" {
		return nil, fmt.Errorf("str parameter can't empty")
	}

	compile, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	options := &options{}
	for _, o := range ops {
		o.apply(options)
	}

	var r = bytes.NewReader([]byte(str))
	var content = make([]string, 0, 256)
	var reader = bufio.NewReader(r)
	var prefix = bytes.Buffer{}
	for {
		data, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		if options.ignoreBlankLine {
			if len(bytes.TrimSpace(data)) == 0 {
				continue
			}
		}

		if isPrefix {
			prefix.Write(data)
			continue
		} else {
			prefix.Write(data)
		}

		if !compile.Match(prefix.Bytes()) {
			prefix.Reset()
			continue
		}

		content = append(
			content,
			compile.FindString(prefix.String()),
		)

		prefix.Reset()
	}

	return content, nil
}
