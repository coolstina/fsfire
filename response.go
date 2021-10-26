package fsfire

import (
	"net/http"
	"strings"
)

// GetResponseHeader Gets response header field value.
func GetResponseHeader(resp http.ResponseWriter, field string) string {
	return resp.Header().Get(field)
}

// GetResponseHeader Gets response header field value,
// processing of response field values by attaching custom functions.
func GetResponseHeaderFunc(resp http.ResponseWriter, field string, process func(val string) string) string {
	return process(GetResponseHeader(resp, field))
}

// ContentDispositionFilename Gets filename for response header Content-Disposition.
// Header such as Content-Disposition: attachment; filename="helloshaohua.txt"
func ContentDispositionFilename() func(val string) string {
	return func(val string) string {
		sli := strings.SplitAfter(val, ";")
		for _, item := range sli {
			if space := strings.TrimSpace(item); strings.Contains(space, "filename") {
				fsl := strings.SplitAfter(space, "=")
				return strings.Trim(fsl[len(fsl)-1], "\"")
			}
		}
		return val
	}
}
