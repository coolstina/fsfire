package fsfire

import (
	"bytes"
	"io/ioutil"
)

// Lines Count the number of data rows
func Lines(data []byte) int {
	if len(data) == 0 {
		return 0
	}
	return len(bytes.Split(data, []byte("\n")))
}

// LinesForFile Count the number of data rows from file.
func LinesForFile(filename string) (int, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	return Lines(data), nil
}

// Characters Count the number of data characters
func Characters(data []byte) int {
	if len(data) == 0 {
		return 0
	}
	data = CleanBlankLines(data)

	bu := bytes.Buffer{}
	defer bu.Reset()

	ls := bytes.SplitAfter(data, []byte("\n"))
	for _, l := range ls {
		bu.Write(bytes.TrimSpace(l))
	}

	return len(bytes.Runes(bu.Bytes()))
}

// CharactersForFile Count the number of data characters
func CharactersForFile(filename string) (int, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	return Characters(data), nil
}
