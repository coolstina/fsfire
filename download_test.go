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
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownloadFile(t *testing.T) {
	grids := []struct {
		baseDir     string
		savePath    FileSystemPath
		extension   FileExtension
		filename    string
		response    *httptest.ResponseRecorder
		contentType string
	}{
		{

			baseDir:     "test/data/downloads",
			savePath:    "textfiles",
			extension:   "txt",
			filename:    "test/data/assets/textfiles/helloshaohua.txt",
			response:    httptest.NewRecorder(),
			contentType: "text/plain",
		},
		{
			baseDir:     "test/data/downloads",
			savePath:    "images",
			extension:   "jpeg",
			filename:    "test/data/assets/images/helloshaohua.jpeg",
			response:    httptest.NewRecorder(),
			contentType: "image/jpeg",
		},
	}

	for _, grid := range grids {
		err := DownloadFile(grid.response, grid.filename)
		assert.NoError(t, err)

		actualFilename := GetResponseHeaderFunc(grid.response, "Content-Disposition", ContentDispositionFilename())
		if actualFilename != "" {
			actualFilename = func(str string) string {
				sli := strings.Split(str, ".")
				if len(sli) == 2 {
					return sli[0]
				}
				return str
			}(actualFilename)
		}

		actual, err := CreateFile(
			actualFilename,
			func() []byte {
				data, err := ioutil.ReadAll(grid.response.Body)
				assert.NoError(t, err)
				return data
			}(),
			WithSpecificBaseDir(grid.baseDir),
			WithSpecificFSPath(grid.savePath),
			WithSpecificFileExtension(grid.extension),
		)
		contains := filepath.Join(grid.baseDir, grid.savePath.String())
		assert.Contains(t, actual, contains)
	}
}
