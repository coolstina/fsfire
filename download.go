package fsfire

import (
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
)

// DownloadFile DownloadFile will write download file into the HTTP response body.
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
