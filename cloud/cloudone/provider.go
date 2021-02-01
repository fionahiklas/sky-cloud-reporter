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

	return &provider{
		BaseUrl: baseUrl,
		Done: false,
	}
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
		machines = make([]reporter.MachineInstance, 0, len(instances))
		for _, instance := range instances {
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