package internal

import (
	"encoding/json"
	"log"
	"testing"
)

func TestUnmarshalSimpleJson(t *testing.T) {
	jsonString := `{ "total": 5, "count": 1, "instances": [{ "id": "1265465", "name": "Gaspode" }, { "id": "09968567", "name": "Gavin" }]}`
	var unmarshalledJson interface{}

	json.Unmarshal([]byte(jsonString), &unmarshalledJson)
	log.Printf("Type of unmarshalled json: %T", unmarshalledJson)
}
