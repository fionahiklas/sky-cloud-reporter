//go:generate mockgen -package=mock_grab -destination=../mocks/mock_grab/mock_http_interface.go . HttpClient

package grab

import "github.com/fionahiklas/sky-cloud-reporter/common/http"

type HttpClient interface {
	Get(string) *http.Response
}