package grab

import (
	"bytes"
	"github.com/fionahiklas/sky-cloud-reporter/common/reporter"
	"github.com/fionahiklas/sky-cloud-reporter/mocks/mock_grab"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

const testSimpleJson = `
[{
"ID": "LadyMargolotta", 
"TeamName": "vampires",
"Machine": "t2.large",
"IPAddress": "240.99.253.110", 
"DeployedRegion": "uberwald",
"State": "dead"
},
{
"ID": "Angua", 
"TeamName": "werewolves",
"Machine": "t2.large",
"IPAddress": "240.99.253.123", 
"DeployedRegion": "ankhmorpork",
"State": "running"
}]
`

const testPagingJsonOne = `
{
"total": 2,
"count": 1,
"instances": [{
	"instance_id": "LadyMargolotta", 
	"team": "vampires",
	"instance_type": "t2.large",
	"ip_address": "240.99.253.110", 
	"region": "uberwald",
	"instance_state": "dead"
}]
`

const testPagingJsonTwo = `
{
"total": 2,
"count": 1,
"instances": [{
	"instance_id": "Angua", 
	"team": "werewolves",
	"instance_type": "t2.large",
	"ip_address": "240.99.253.123", 
	"region": "ankhmorpork",
	"instance_state": "running"
}]
`


func TestNewGrabber(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	httpClient := mock_grab.NewMockHttpClient(ctrl)
	cloudProvider := mock_grab.NewMockCloudProvider(ctrl)

	result := NewGrabber(httpClient, cloudProvider)
	assert.NotNil(result)
	assert.IsType(&cloudGrabber{}, result)
}

func TestGrabInstancesSimpleCloud(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	httpClient := mock_grab.NewMockHttpClient(ctrl)
	cloudProvider := mock_grab.NewMockCloudProvider(ctrl)

	grabber := NewGrabber(httpClient, cloudProvider)

	const urlString = "http://anhk.morpork/instances"
	httpResponse := http.Response{
		StatusCode: 200,
		Body: convertJsonStringToReadCloser(testSimpleJson),
	}
	machineInstances := []reporter.MachineInstance{}

	cloudProvider.EXPECT().
		ResetFunction().
		Return(func() { log.Printf("RESET FUNCTION CALLED") })

	gomock.InOrder(
		cloudProvider.EXPECT().GenerateNextUrl().Return(urlString, false),
		cloudProvider.EXPECT().GenerateNextUrl().Return("", true),
	)

	httpClient.EXPECT().
		Get(gomock.Eq(urlString)).
		Return(&httpResponse, nil)

	cloudProvider.EXPECT().
		ProcessResponse(gomock.Eq(&httpResponse)).
		Return(machineInstances, nil)

	result, resultError := grabber.GrabInstances()
	assert.NotNil(result)
	assert.Equal(machineInstances, result)
	assert.Nil(resultError)
}


func TestGrabInstancesCloudWithPaging(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	httpClient := mock_grab.NewMockHttpClient(ctrl)
	cloudProvider := mock_grab.NewMockCloudProvider(ctrl)

	grabber := NewGrabber(httpClient, cloudProvider)

	const urlStringOne = "http://anhk.morpork/instances?page=1"
	const urlStringTwo = "http://anhk.morpork/instances?page=2"
	httpResponseOne := http.Response{
		StatusCode: 200,
		Body: convertJsonStringToReadCloser(testPagingJsonOne),
	}
	httpResponseTwo := http.Response{
		StatusCode: 200,
		Body: convertJsonStringToReadCloser(testPagingJsonTwo),
	}

	machineInstancesOne := []reporter.MachineInstance{
		{
			Id: "LadyMargolotta",
			Team: "vampires",
		},
	}

	machineInstancesTwo := []reporter.MachineInstance{
		{
			Id: "Angua",
			Team: "werewolves",
		},
	}

	machineInstancesResult := make([]reporter.MachineInstance, 0, 2)
	machineInstancesResult = append(machineInstancesResult, machineInstancesOne...)
	machineInstancesResult = append(machineInstancesResult, machineInstancesTwo...)

	cloudProvider.EXPECT().
		ResetFunction().
		Return(func() { log.Printf("RESET FUNCTION CALLED") })

	gomock.InOrder(
		cloudProvider.EXPECT().GenerateNextUrl().Return(urlStringOne, false),
		cloudProvider.EXPECT().GenerateNextUrl().Return(urlStringTwo, false),
		cloudProvider.EXPECT().GenerateNextUrl().Return("", true),
	)

	gomock.InOrder(
		httpClient.EXPECT().Get(gomock.Eq(urlStringOne)).Return(&httpResponseOne, nil),
		httpClient.EXPECT().Get(gomock.Eq(urlStringTwo)).Return(&httpResponseTwo, nil),
	)

	gomock.InOrder(
		cloudProvider.EXPECT().ProcessResponse(gomock.Eq(&httpResponseOne)).Return(machineInstancesOne, nil),
		cloudProvider.EXPECT().ProcessResponse(gomock.Eq(&httpResponseTwo)).Return(machineInstancesTwo, nil),
	)

	result, resultError := grabber.GrabInstances()
	assert.NotNil(result)
	assert.Equal(machineInstancesResult, result)
	assert.Nil(resultError)
}


func TestGrabInstancesSimpleCloudResetCalled(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	resetCalled := false

	httpClient := mock_grab.NewMockHttpClient(ctrl)
	cloudProvider := mock_grab.NewMockCloudProvider(ctrl)

	grabber := NewGrabber(httpClient, cloudProvider)

	const urlString = "http://anhk.morpork/instances"
	httpResponse := http.Response{
		StatusCode: 200,
		Body: convertJsonStringToReadCloser(testSimpleJson),
	}
	machineInstances := []reporter.MachineInstance{}

	gomock.InOrder(
		cloudProvider.EXPECT().GenerateNextUrl().Return(urlString, false),
		cloudProvider.EXPECT().GenerateNextUrl().Return("", true),
	)

	cloudProvider.EXPECT().
		ProcessResponse(gomock.Eq(&httpResponse)).
		Return(machineInstances, nil)

	cloudProvider.EXPECT().
		ResetFunction().
		Return(func() {
			log.Printf("RESET FUNCTION CALLED")
			resetCalled = true
		})

	httpClient.EXPECT().
		Get(gomock.Eq(urlString)).
		Return(&httpResponse, nil)

	// Call the grabber
	result, resultError := grabber.GrabInstances()

	assert.NotNil(result)
	assert.Equal(machineInstances, result)
	assert.Nil(resultError)
	assert.Equal(resetCalled, true)
}



func convertJsonStringToReadCloser(jsonString string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(jsonString)))
}