package http_body

import (
	"io"
	"net/http"
	"strconv"
)

func ReadHttpBody(body io.ReadCloser, header http.Header) ([]byte, error) {
	contentLength := 4096
	contentLengthHeader := header.Get("Content-Length")
	if contentLengthHeader != "" {
		contentLengthFromHeader, err := strconv.Atoi(contentLengthHeader)
		if err == nil {
			contentLength = contentLengthFromHeader
		}
	}
	respBody := make([]byte, 0, contentLength)
	var err error
	for {
		if len(respBody) == cap(respBody) {
			// Add more capacity (let append pick how much).
			respBody = append(respBody, 0)[:len(respBody)]
		}
		n := 0
		n, err = body.Read(respBody[len(respBody):cap(respBody)])
		respBody = respBody[:len(respBody)+n]
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
	}
	CloseBody(&body)
	return respBody, err
}

func CloseBody(body *io.ReadCloser) {
	if body != nil {
		_ = (*body).Close()
	}
}
