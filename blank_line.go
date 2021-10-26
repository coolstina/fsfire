package fsfire

import (
	"bytes"
	"io/ioutil"
)

// CleanBlankLines clean blank lines with file.
func CleanBlankLines(data []byte) []byte {
	if len(data) == 0 {
		return nil
	}
	bu := bytes.Buffer{}
	ls := bytes.SplitAfter(data, []byte("\n"))
	for _, l := range ls {
		if len(bytes.TrimSpace(l)) == 0 {
			continue
		}
		bu.Write(l)
	}
	return bu.Bytes()
}

// CleanBlankLinesForFile clean blank lines with file.
func CleanBlankLinesForFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := CleanBlankLines(data)
	return lines, nil
}
