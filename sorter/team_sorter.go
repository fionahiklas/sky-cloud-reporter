package sorter

import (
	"github.com/fionahiklas/sky-cloud-reporter/common/reporter"
	"log"
)

func SortIntoTeams(instances []reporter.MachineInstance, resultMap reporter.MachineReport, onlyRunning bool) {
	for _, machineInstance := range instances {

		if onlyRunning && machineInstance.State != "running" {
			log.Printf("Skipping ID '%s' in team '%s' because it's state '%s' is not running ",
				machineInstance.Id, machineInstance.Team, machineInstance.State)
			continue
		}

		log.Printf("Checking result map for team: %s", machineInstance.Team)
		teamInMap, ok := resultMap[machineInstance.Team]

		if ok == true {
			log.Printf("Found team '%s' in map", machineInstance.Team)
		} else {
			log.Printf("Did not find team '%s' in map, adding it", machineInstance.Team)
			teamInMap = new(reporter.TeamInstances) // Needs to be a new block of memory
			teamInMap.Instances = make([]*reporter.MachineInstance, 0, 10)
			resultMap[machineInstance.Team] = teamInMap
		}

		log.Printf("Adding to team '%s' the instance ID: %s", machineInstance.Team, machineInstance.Id)
		instanceInMap := new(reporter.MachineInstance)
		*instanceInMap = machineInstance // Go is sneaky with pointers and values *pointer means value :)

		teamInMap.Instances = append(teamInMap.Instances, instanceInMap)
		teamInMap.Count += 1
	}
}
