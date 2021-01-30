package internal

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestUnmarshalSimpleJson(t *testing.T) {
	jsonString := `{ "total": 5, "count": 1, "instances": [{ "id": "1265465", "name": "Gaspode" }, { "id": "09968567", "name": "Gavin" }]}`
	var unmarshalledJson interface{}

	err := json.Unmarshal([]byte(jsonString), &unmarshalledJson)
	if err != nil {
		t.Fatalf("Error unmarshalling JSON: %s", err)
	}
	log.Printf("Type of unmarshalled json: %T", unmarshalledJson)
}

func TestUnmarshalCloudTwo(t *testing.T) {

	assert := assert.New(t)

	type CloudTwoInstance struct {
		Instance_id    string
		Team           string
		Instance_type  string
		Ip_address     string
		Region         string
		Instance_state string
	}

	type CloudTwo struct {
		Total     int
		Count     int
		Instances []CloudTwoInstance
	}

	jsonString := `
{
	"total": 42,
  	"count": 2,
    "instances": [
		{
			"instance_id": "28b288ae-9c1a-4ae3-89a3-26bbb7235ee3", 
			"team": "red",
			"instance_type": "t2.large",
			"ip_address": "141.15.214.194",
			"region": "us-west-2",
			"instance_state": "running" 
		},
		{
			"instance_id": "3", 
			"team": "",
			"instance_type": "t2.large",
			"ip_address": "141.15.214.194",
			"region": "us-west-2",
			"instance_state": "running" 
		}

	]
}
`

	var unmarshalledJson CloudTwo
	err := json.Unmarshal([]byte(jsonString), &unmarshalledJson)

	assert.NoError(err, "Should not error unmarshalling: %s", err)
	assert.Equal(42, unmarshalledJson.Total)
	assert.Equal(2, unmarshalledJson.Count)
	assert.Equal(2, len(unmarshalledJson.Instances))

}
