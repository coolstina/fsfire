package fsfire

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
)

// UploadFileForMultipartFields Mock form upload file.
func UploadFileForMultipartFields(values map[string]interface{}) (contentType string, buffer bytes.Buffer, err error) {
	w := multipart.NewWriter(&buffer)
	defer w.Close()
	for key, r := range values {
		switch val := r.(type) {
		case *os.File:
			var fw io.Writer
			if fw, err = w.CreateFormFile(key, val.Name()); err != nil {
				return
			}
			if _, err = io.Copy(fw, val); err != nil {
				return
			}
		case string:
			if err = w.WriteField(key, val); err != nil {
				return
			}
		}
	}
	contentType = w.FormDataContentType()

	return
}
