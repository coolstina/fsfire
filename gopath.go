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
