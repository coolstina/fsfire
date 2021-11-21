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

type FileExtension string

func (file FileExtension) String() string {
	return string(file)
}

const (
	FileExtensionWithImageOfJPEG FileExtension = "jpeg"
	FileExtensionWithImageOfPNG  FileExtension = "png"
	FileExtensionWithImageOfGIF  FileExtension = "gif"
	FileExtensionWithImageOfBMP  FileExtension = "bmp"
)

const (
	FileExtensionWithOfficeOfExcel FileExtension = "xlsx"
	FileExtensionWithOfficeOfWord  FileExtension = "docx"
	FileExtensionWithOfficeOfCSV   FileExtension = "csv"
)

const (
	FileExtensionWithDocumentOfMarkdown FileExtension = "md"
	FileExtensionWithDocumentOfFileText FileExtension = "txt"
)
