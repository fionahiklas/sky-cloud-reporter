package grab

import (
	"github.com/fionahiklas/sky-cloud-reporter/common/http"
	"github.com/fionahiklas/sky-cloud-reporter/mocks/mock_grab"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

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

	cloudProvider.EXPECT().GetInstanceUrl().
		Return(urlString).
		MaxTimes(1)

	httpClient.EXPECT().
		Get(gomock.Eq(urlString)).
		Return(new(http.Response))

	result := grabber.GrabInstances()
	assert.NotNil(result)
}


