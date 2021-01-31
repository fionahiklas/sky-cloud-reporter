package grab

import "github.com/fionahiklas/sky-cloud-reporter/common/reporter"

type cloudGrabber struct {
	httpClient HttpClient
	cloudProvider CloudProvider
}

func NewGrabber(client HttpClient, provider CloudProvider) *cloudGrabber {
	return &cloudGrabber{
		httpClient: client,
		cloudProvider: provider,
	}
}

func (grabber *cloudGrabber) GrabInstances() (instances *[]reporter.MachineInstance, err error) {
	urlString := grabber.cloudProvider.GetInstanceUrl()
	httpResponse, _ := grabber.httpClient.Get(urlString)
	instances, err = grabber.cloudProvider.ConvertResponseToMachineInstances(httpResponse)
	return instances, err
}
