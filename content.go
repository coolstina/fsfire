package fsfire

import (
	"bytes"
	"fmt"
)

// MarkContentPrefixClear Mark content prefix clear.
func MarkContentPrefixClear(data []byte, mark string, ops ...Option) []byte {
	options := options{
		containsMarkString: true,
	}

	for _, o := range ops {
		o.apply(&options)
	}

	index := bytes.Index(data, []byte(mark))
	if index > 0 {
		if !options.containsMarkString {
			return data[index+(len([]byte(mark))):]
		}
		return data[index:]
	}

	return data
}

// MarkContentSuffixClear Mark content suffix clear.
func MarkContentSuffixClear(data []byte, mark string, ops ...Option) []byte {
	options := options{
		containsMarkString: false,
	}

	for _, o := range ops {
		o.apply(&options)
	}

	index := bytes.Index(data, []byte(mark))
	if index > 0 {
		if options.containsMarkString {
			fmt.Printf("helloshaohua1")
			return data[:index+(len([]byte(mark)))]
		}
		fmt.Printf("helloshaohua2")
		return data[:index]
	}
	return data
}
