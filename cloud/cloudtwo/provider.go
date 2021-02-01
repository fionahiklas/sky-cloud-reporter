package cloudtwo

import (
	"encoding/json"
	"github.com/fionahiklas/sky-cloud-reporter/common/reporter"
	"io/ioutil"
	"log"
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
	log.Printf("Creating CloudTwo provider")
	result := new(provider)
	*result = provider{
		BaseUrl: baseUrl,
		CurrentPage: 1,
		Total: 0,
		ProcessedSoFar: 0,
	}
	return result
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
	log.Printf("CloudTwo URL: %s, done: %t", url, done)
	return
}

func (provider *provider) ProcessResponse(response *http.Response) (machines []reporter.MachineInstance, err error) {
	log.Printf("Processing CloudTwo, CurrentPage: %d, Total: %d, ProcessedSoFar: %d",
		provider.CurrentPage, provider.Total, provider.ProcessedSoFar)

	var instances CloudTwo
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	jsonErr := json.Unmarshal(bodyBytes, &instances)

	if jsonErr == nil {
		// We know the total number of results here so the slice can have a fixed capacity
		machines = make([]reporter.MachineInstance, 0, instances.Count)
		for _, instance := range instances.Instances {
			// This should always be fine since the capacity should never be exceeded
			machines = append(machines, convertCloudStructToCommon(instance))
		}

		provider.Total = instances.Total
		provider.ProcessedSoFar += instances.Count
		provider.CurrentPage += 1
	}

	log.Printf("Processed CloudTwo, CurrentPage: %d, Total: %d, ProcessedSoFar: %d",
		provider.CurrentPage, provider.Total, provider.ProcessedSoFar)
	return
}

func (provider *provider) ResetFunction() func() {
	return func() {
		provider.CurrentPage = 1
		provider.Total = 0
		provider.ProcessedSoFar = 0
	}
}

// TODO: Maybe this would be more efficient to return a pointer
func convertCloudStructToCommon(cloudInstance CloudTwoInstance) reporter.MachineInstance {
	return reporter.MachineInstance{
		Id:      cloudInstance.InstanceId,
		Team:    cloudInstance.Team,
		Machine: cloudInstance.InstanceType,
		Ip:      cloudInstance.IpAddress,
		State:   cloudInstance.InstanceState,
		Region:  cloudInstance.Region,
	}
}