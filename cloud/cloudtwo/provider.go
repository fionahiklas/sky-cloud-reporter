package cloudtwo

import (
	"github.com/fionahiklas/sky-cloud-reporter/common/reporter"
	"net/http"
)

type provider struct {
	BaseUrl string
	CurrentPage int
}

func NewProvider(baseUrl string) (*provider) {
	return &provider{
		BaseUrl: baseUrl,
		CurrentPage: 1,
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
	return nil
}