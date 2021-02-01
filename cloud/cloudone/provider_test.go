package cloudone

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
	assert.Equal(false, provider.RequiresPaging())
}

func TestGenerateNextUrlInitial(t *testing.T) {
	assert := assert.New(t)

	provider := NewProvider(baseUrl)
	url, done := provider.GenerateNextUrl()
	assert.Equal(baseUrl + "/instances", url)
	assert.Equal(false, done)
}

func TestGenerateNextUrlAfterProcessing(t *testing.T) {
	assert := assert.New(t)

	provider := NewProvider(baseUrl)

	provider.ProcessResponse(&http.Response{ StatusCode: 200, Body: convertJsonStringToReadCloser(testSimpleJson)})
	url, done := provider.GenerateNextUrl()
	assert.Equal("", url)
	assert.Equal(true, done)
}

func TestResetFunction(t *testing.T) {
	assert := assert.New(t)

	provider := NewProvider(baseUrl)

	provider.Done = true
	provider.ResetFunction()()

	assert.Equal(false, provider.Done)
}

func TestMappingFromCloudToReporter(t *testing.T) {
	assert := assert.New(t)

	testCloudInstance := CloudOneInstance{
		Id:             "SamVimes",
		TeamName:       "CityWatch",
		Machine:        "BSJohnson",
		IpAddress:      "treaclemine.road",
		DeployedRegion: "TheShades",
		State:          "Vetinari",
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

func convertJsonStringToReadCloser(jsonString string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(jsonString)))
}