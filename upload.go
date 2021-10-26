package fsfire

import (
	"net/http"

	"github.com/gabriel-vasile/mimetype"
)

func UploadIsAllowMIMEWithFormFile(request *http.Request, key string, allows []string) (bool, error) {
	// Gets the file for upload form.
	file, _, err := request.FormFile(key)
	if err != nil {
		return false, err
	}
	defer file.Close()

	// Gets the file content mime.
	mime, err := mimetype.DetectReader(file)
	if err != nil {
		return false, err
	}

	// Match allow mime.
	mimeStr := mime.String()
	allowed := false
	for _, allow := range allows {
		if allow == mimeStr {
			allowed = true
		}
	}

	return allowed, nil
}
