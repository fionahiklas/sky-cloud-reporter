package sorter

import (
	"encoding/json"
	"github.com/fionahiklas/sky-cloud-reporter/common/reporter"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var testMachineInstances = []reporter.MachineInstance{
	{
		Id:      "Rincewind",
		Team:    "wizards",
		Machine: "Hex",
		Ip:      "8.1.1.1",
		State:   "running",
		Region:  "AnkhMorpork",
	},
	{
		Id:      "GrannyWeatherwax",
		Team:    "witches",
		Machine: "Broomstick",
		Ip:      "255.255.255.255",
		State:   "running",
		Region:  "Lancre",
	},
	{
		Id:      "EskarinaSmith",
		Team:    "wizards",
		Machine: "Staff",
		Ip:      "8.3.8.3",
		State:   "running",
		Region:  "AnkhMorpork",
	},
	{
		Id:      "Coin",
		Team:    "wizards",
		Machine: "Sourcery",
		Ip:      "8.8.8.8",
		State:   "gone",
		Region:  "Dimension",
	},
}


var expectedResult = reporter.MachineReport{

	"wizards": {
		Count: 2,
		Instances: []*reporter.MachineInstance{
			{
				Id:      "Rincewind",
				Team:    "wizards",
				Machine: "Hex",
				Ip:      "8.1.1.1",
				State:   "running",
				Region:  "AnkhMorpork",
			},

			{
				Id:      "EskarinaSmith",
				Team:    "wizards",
				Machine: "Staff",
				Ip:      "8.3.8.3",
				State:   "running",
				Region:  "AnkhMorpork",
			},
		},
	},

	"witches": {
		Count: 1,
		Instances: []*reporter.MachineInstance {
			{
				Id:      "GrannyWeatherwax",
				Team:    "witches",
				Machine: "Broomstick",
				Ip:      "255.255.255.255",
				State:   "running",
				Region:  "Lancre",
			},
		},

	},

}

func TestSortIntoTeams(t *testing.T) {
	assert := assert.New(t)

	resultMap := make(reporter.MachineReport)
	SortIntoTeams(testMachineInstances, resultMap, true)

	resultAsJson, resultErr := json.MarshalIndent(resultMap, "", "    ")

	if resultErr != nil {
		log.Printf("JsonError: %s", resultErr)
	} else {
		log.Printf("RESULT:\n%s", resultAsJson)
	}

	assert.Equal(2, len(resultMap))
	assert.Equal(expectedResult, resultMap)
}
