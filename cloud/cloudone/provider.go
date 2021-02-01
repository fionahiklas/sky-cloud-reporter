package cloudone

import (
	"encoding/json"
	"github.com/fionahiklas/sky-cloud-reporter/common/reporter"
	"io/ioutil"
	"log"
	"net/http"
)

type provider struct {
	BaseUrl string
	Done bool
}

func NewProvider(baseUrl string) *provider {
	log.Printf("Creating CloudOne provider")
	result := new(provider)
	*result = provider{
		BaseUrl: baseUrl,
		Done: false,
	}
	return result
}

func (provider *provider) RequiresPaging() bool {
	return false
}

func (provider *provider) GenerateNextUrl() (url string, done bool) {
	done = provider.Done
	if done == false {
		url = provider.BaseUrl + "/instances"
	} else {
		url = ""
	}
	log.Printf("CloudOne URL: %s, done: %t", url, done)
	return
}

func (provider *provider) ProcessResponse(response *http.Response) (machines []reporter.MachineInstance, err error) {
	log.Printf("Processing CloudOne, done: %t", provider.Done)

	var instances CloudOne
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	jsonErr := json.Unmarshal(bodyBytes, &instances)

	if jsonErr == nil {
		// We know how big the list of instances is
		machines = make([]reporter.MachineInstance, 0, len(instances))
		for _, instance := range instances {
			// This should always be fine since the slice has the correct capacity
			machines = append(machines, convertCloudStructToCommon(instance))
		}
		provider.Done = true
	}

	log.Printf("Processed CloudOne, done: %t", provider.Done)
	return
}

func (provider *provider) ResetFunction() func() {
	return func() {
		provider.Done = false
	}
}

// TODO: Maybe this would be more efficient to return a pointer
func convertCloudStructToCommon(cloudInstance CloudOneInstance) reporter.MachineInstance {
	return reporter.MachineInstance{
		Id:      cloudInstance.Id,
		Team:    cloudInstance.TeamName,
		Machine: cloudInstance.Machine,
		Ip:      cloudInstance.IpAddress,
		State:   cloudInstance.State,
		Region:  cloudInstance.DeployedRegion,
	}
}