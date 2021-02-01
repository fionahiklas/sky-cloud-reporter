package cloudtwo

import (
	"bytes"
	"github.com/fionahiklas/sky-cloud-reporter/common/reporter"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

const baseUrl = "http://ankh-morpork.disc"
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
}
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
}
`

var testReporterInstancesOne = []reporter.MachineInstance{
	{
		Id:      "LadyMargolotta",
		Team:    "vampires",
		Machine: "t2.large",
		Ip:      "240.99.253.110",
		State:   "dead",
		Region:  "uberwald",
	},
}

func TestNewProvider(t *testing.T) {
	assert := assert.New(t)

	result := NewProvider(baseUrl)
	assert.NotNil(result)
	assert.IsType(&provider{}, result)
	assert.Equal(baseUrl, result.BaseUrl)
}

func TestRequiresPaging(t *testing.T) {
	assert := assert.New(t)

	provider := NewProvider(baseUrl)
	assert.Equal(true, provider.RequiresPaging())
}

func TestGenerateNextUrlInitial(t *testing.T) {
	assert := assert.New(t)

	provider := NewProvider(baseUrl)
	url, done := provider.GenerateNextUrl()
	assert.Equal(baseUrl + "/cloud/instances?page=1", url)
	assert.Equal(false, done)
}

func TestGenerateNextUrlAfterProcessing(t *testing.T) {
	assert := assert.New(t)

	provider := NewProvider(baseUrl)

	provider.ProcessResponse(&http.Response{ StatusCode: 200, Body: convertJsonStringToReadCloser(testPagingJsonOne)})
	url, done := provider.GenerateNextUrl()
	assert.Equal(baseUrl + "/cloud/instances?page=2", url)
	assert.Equal(false, done)
}

func TestGenerateNextUrlAfterAllProcessing(t *testing.T) {
	assert := assert.New(t)

	provider := NewProvider(baseUrl)

	provider.ProcessResponse(&http.Response{ StatusCode: 200, Body: convertJsonStringToReadCloser(testPagingJsonOne)})
	provider.ProcessResponse(&http.Response{ StatusCode: 200, Body: convertJsonStringToReadCloser(testPagingJsonTwo)})
	url, done := provider.GenerateNextUrl()
	assert.Equal("", url)
	assert.Equal(true, done)
}


func TestResetFunction(t *testing.T) {
	assert := assert.New(t)

	provider := NewProvider(baseUrl)

	provider.CurrentPage = 2
	provider.Total = 2
	provider.ProcessedSoFar = 2
	provider.ResetFunction()()

	assert.Equal(1, provider.CurrentPage)
	assert.Equal(0, provider.Total)
	assert.Equal(0, provider.ProcessedSoFar)
}

func TestMappingFromCloudToReporter(t *testing.T) {
	assert := assert.New(t)

	testCloudInstance := CloudTwoInstance{
		InstanceId:    "SamVimes",
		Team:          "CityWatch",
		InstanceType:  "BSJohnson",
		IpAddress:     "treaclemine.road",
		Region:        "TheShades",
		InstanceState: "Vetinari",
	}

	expectedReporterInstance := reporter.MachineInstance{
		Id:      "SamVimes",
		Team:    "CityWatch",
		Machine: "BSJohnson",
		Ip:      "treaclemine.road",
		State:   "Vetinari",
		Region:  "TheShades",
	}

	result := convertCloudStructToCommon(testCloudInstance)

	assert.Equal(expectedReporterInstance, result)
}

func TestProcessResponseConversion(t *testing.T) {
	assert := assert.New(t)

	provider := NewProvider(baseUrl)
	result, err := 	provider.ProcessResponse(&http.Response{ StatusCode: 200, Body: convertJsonStringToReadCloser(testPagingJsonOne)})

	assert.Nil(err)
	assert.Equal(testReporterInstancesOne, result)
	assert.Equal(2, provider.Total)
	assert.Equal(2, provider.CurrentPage)
	assert.Equal(1, provider.ProcessedSoFar)
}

func convertJsonStringToReadCloser(jsonString string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(jsonString)))
}