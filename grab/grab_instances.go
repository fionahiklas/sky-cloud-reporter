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

// TODO: I'm not sure how efficient returning a slice is - does it use reference of
// TODO: value and if the latter is is copying the underlying array or just pointers?
func (grabber *cloudGrabber) GrabInstances() (instances []reporter.MachineInstance, err error) {

	// This resets the cloud provider once we've grabbed all the data - the
	// syntax looks weird here since the ResetFunction method returns a
	// function which then needs the extra () to actually be executed
	defer grabber.cloudProvider.ResetFunction()()

	// TODO: The capacity needs to be a tunable value
	resultCollector := make([]reporter.MachineInstance, 0, 10)

	for {
		urlString, done := grabber.cloudProvider.GenerateNextUrl()
		if done {
			break
		}
		httpResponse, _ := grabber.httpClient.Get(urlString)
		processedInstances, _ := grabber.cloudProvider.ProcessResponse(httpResponse)

		// TODO: Currently this won't cope with exceeding the capacity of the slice
		resultCollector = append(resultCollector, processedInstances...)
	}
	return resultCollector, nil
}
