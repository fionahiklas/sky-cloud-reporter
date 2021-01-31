package cloudtwo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const baseUrl = "http://ankh-morpork.disc/"
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
