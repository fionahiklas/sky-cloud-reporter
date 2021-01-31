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

func (grabber *cloudGrabber) GrabInstances() *[]reporter.MachineInstance {
	urlString := grabber.cloudProvider.GetInstanceUrl()
	grabber.httpClient.Get(urlString)
	return new([]reporter.MachineInstance)
}
