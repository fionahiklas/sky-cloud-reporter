package cloudtwo

import (
	"encoding/json"
	"github.com/fionahiklas/sky-cloud-reporter/common/reporter"
	"io/ioutil"
	"net/http"
	"strconv"
)

type provider struct {
	BaseUrl string
	CurrentPage int
	Total int
	ProcessedSoFar int
}

func NewProvider(baseUrl string) (*provider) {
	return &provider{
		BaseUrl: baseUrl,
		CurrentPage: 1,
		Total: 0,
		ProcessedSoFar: 0,
	}
}

func (provider *provider) RequiresPaging() bool {
	return true
}

func (provider *provider) GenerateNextUrl() (url string, done bool) {
	done = provider.Total > 0 && provider.ProcessedSoFar >= provider.Total

	if done == false {
		url = provider.BaseUrl + "/cloud/instances?page=" + strconv.Itoa(provider.CurrentPage)
	} else {
		url = ""
	}
	return
}

func (provider *provider) ProcessResponse(response *http.Response) (machines []reporter.MachineInstance, err error) {
	var instance CloudTwo
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	jsonErr := json.Unmarshal(bodyBytes, &instance)

	if jsonErr == nil {
		provider.Total = instance.Total
		provider.ProcessedSoFar += instance.Count
		provider.CurrentPage += 1
	}
	return nil, nil
}

func (provider *provider) ResetFunction() func() {
	return func() {
		provider.CurrentPage = 1
		provider.Total = 0
		provider.ProcessedSoFar = 0
	}
}