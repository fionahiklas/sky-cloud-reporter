//go:generate mockgen -package=mock_grab -destination=../mocks/mock_grab/mock_http_interface.go . HttpClient

package grab

import "net/http"

type HttpClient interface {
	Get(string) (resp *http.Response, err error)
}