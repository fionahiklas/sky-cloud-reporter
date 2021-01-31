package cloudone

import (
	"github.com/fionahiklas/sky-cloud-reporter/common/reporter"
	"net/http"
)

type provider struct {
	BaseUrl string
	Done bool
}

func NewProvider(baseUrl string) (*provider) {
	return &provider{
		BaseUrl: baseUrl,
		Done: false,
	}
}

func (provider *provider) RequiresPaging() bool {
	return false
}

func (provider *provider) GenerateNextUrl() (url string, done bool) {
	return "", false
}

func (provider *provider) ProcessResponse(response *http.Response) (machines *[]reporter.MachineInstance, err error) {
	return nil, nil
}

func (provider *provider) ResetFunction() func() {
	return func() {
		provider.Done = false
	}
}