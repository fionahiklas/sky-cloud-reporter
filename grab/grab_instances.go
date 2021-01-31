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

func (grabber *cloudGrabber) GrabInstances() (instances []reporter.MachineInstance, err error) {
	resultCollector := make([]reporter.MachineInstance, 0, 10)

	for {
		urlString, done := grabber.cloudProvider.GenerateNextUrl()
		if done {
			break
		}
		httpResponse, _ := grabber.httpClient.Get(urlString)
		processedInstances, _ := grabber.cloudProvider.ProcessResponse(httpResponse)
		resultCollector = append(resultCollector, processedInstances...)
	}
	return resultCollector, nil
}
