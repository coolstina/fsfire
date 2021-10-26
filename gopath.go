package fsfire

import (
	"bytes"
	"os"
)

const (
	FileSystemOfGOPATH = "GOPATH"
)

// GetGOPATH Gets the current operating system GOPATH from environment.
func GetGOPATH() string {
	return os.Getenv(FileSystemOfGOPATH)
}

// GetGOPATHProjectSource Gets the current operating system GOPATH/src from environment and filesystem.
func GetGOPATHProjectSource() string {
	var buffer bytes.Buffer

	buffer.WriteString(GetGOPATH())
	buffer.WriteRune(os.PathSeparator)
	buffer.WriteString("src")
	defer buffer.Reset()

	return buffer.String()
}
