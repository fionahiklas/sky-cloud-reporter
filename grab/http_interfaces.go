//go:generate mockgen -destination=../mocks/mock_http_interfaces.go . HttpResponse,HttpClient

package grab

import "io"

type HttpResponse interface {
	StatusCode() int
	Body() io.ReadCloser
}

type HttpClient interface {
	Get(string) *HttpResponse
}