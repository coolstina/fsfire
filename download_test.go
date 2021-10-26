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
		savePath    FSPath
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
