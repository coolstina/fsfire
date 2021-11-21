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
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
)

// DownloadFile will write download file into the HTTP response body.
func DownloadFile(resp http.ResponseWriter, path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Gets filename with file path.
	filename := filepath.Base(path)

	// Detect returns the MIME type found from the provided byte slice.
	detect := mimetype.Detect(data)

	// Setting http response header.
	resp.Header().Set("Content-Disposition", `attachment; filename="`+filename+`"`)
	resp.Header().Set("Content-Type", detect.String())

	// Write data into the response body.
	if _, err := resp.Write(data); err != nil {
		return err
	}

	// Write status code into the response header.
	resp.WriteHeader(http.StatusOK)

	return nil
}
