package grab

import (
	"bytes"
	"github.com/fionahiklas/sky-cloud-reporter/common/reporter"
	"github.com/fionahiklas/sky-cloud-reporter/mocks/mock_grab"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
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

	cloudProvider.EXPECT().GenerateNextUrl().
		Return(urlString, true).
		MaxTimes(1)

	httpClient.EXPECT().
		Get(gomock.Eq(urlString)).
		Return(&httpResponse, nil)

	cloudProvider.EXPECT().
		ProcessResponse(gomock.Eq(&httpResponse)).
		Return(&machineInstances, nil)

	result, resultError := grabber.GrabInstances()
	assert.NotNil(result)
	assert.Equal(&machineInstances, result)
	assert.Nil(resultError)
}


func convertJsonStringToReadCloser(jsonString string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(jsonString)))
}