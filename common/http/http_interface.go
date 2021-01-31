//go:generate mockgen -package=mock_http -destination=../../mocks/mock_http/mock_http_interface.go . Response

package http

import "io"

type Response interface {
	StatusCode() int
	Body() io.ReadCloser
}


