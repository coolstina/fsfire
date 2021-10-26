package fsfire

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetResponseHeaderFunc(t *testing.T) {
	grids := []struct {
		filename         string
		response         *httptest.ResponseRecorder
		contentType      string
		expectedFilename string
	}{
		{
			filename:         "test/data/assets/textfiles/hello.txt",
			response:         httptest.NewRecorder(),
			contentType:      "text/plain",
			expectedFilename: "hello.txt",
		},
		{
			filename:         "test/data/assets/images/helloshaohua.jpeg",
			response:         httptest.NewRecorder(),
			contentType:      "image/jpeg",
			expectedFilename: "helloshaohua.jpeg",
		},
	}

	for _, grid := range grids {
		err := DownloadFile(grid.response, grid.filename)
		assert.NoError(t, err)

		actualFilename := GetResponseHeaderFunc(grid.response, "Content-Disposition", ContentDispositionFilename())
		assert.Equal(t, grid.expectedFilename, actualFilename)
	}

}
