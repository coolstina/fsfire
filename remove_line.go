package fsfire

import (
	"bytes"
	"io/ioutil"
	"regexp"
)

// CleanSpecificContentLines clean specific content lines.
func CleanSpecificContentLines(data []byte, findStr string) []byte {
	buffer := bytes.Buffer{}

	lines := bytes.SplitAfter(data, []byte("\n"))
	for _, line := range lines {
		if len(regexp.MustCompile(findStr).Find(line)) != 0 {
			continue
		}

		buffer.Write(line)
	}

	return buffer.Bytes()
}

// CleanSpecificContentLinesForFile clean specific content lines for file.
func CleanSpecificContentLinesForFile(filename string, findStr string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	buffer := bytes.Buffer{}
	lines := bytes.SplitAfter(data, []byte("\n"))
	for _, line := range lines {
		if len(regexp.MustCompile(findStr).Find(line)) != 0 {
			continue
		}

		buffer.Write(line)
	}

	return buffer.Bytes(), nil
}
