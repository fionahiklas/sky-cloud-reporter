package main

import (
	"encoding/json"
	"github.com/fionahiklas/sky-cloud-reporter/cloud/cloudone"
	"github.com/fionahiklas/sky-cloud-reporter/cloud/cloudtwo"
	"github.com/fionahiklas/sky-cloud-reporter/common/reporter"
	"github.com/fionahiklas/sky-cloud-reporter/grab"
	"github.com/fionahiklas/sky-cloud-reporter/sorter"
	"log"
	"net/http"
	"os"
)

func main() {
	providerOneBaseUrl := os.Args[1]
	providerTwoBaseUrl := os.Args[2]

	log.Printf("Creating provider one grabber with URL: %s", providerOneBaseUrl)
	providerOne := cloudone.NewProvider(providerOneBaseUrl)
	httpClientOne := &http.Client{}
	grabberOne := grab.NewGrabber(httpClientOne, providerOne)

	log.Printf("Creating provider two grabber with URL: %s", providerTwoBaseUrl)
	providerTwo := cloudtwo.NewProvider(providerTwoBaseUrl)
	httpClientTwo := &http.Client{}
	grabberTwo := grab.NewGrabber(httpClientTwo, providerTwo)

	log.Printf("Grabbing instances from provider one")
	grabOneInstances, _ := grabberOne.GrabInstances()

	log.Printf("Grabbing instances from provider two")
	grabTwoInstances, _ := grabberTwo.GrabInstances()

	log.Printf("Sticking the results together")
	allInstanceResults := make([]reporter.MachineInstance, 0, 10)
	allInstanceResults = append(allInstanceResults, grabOneInstances...)
	allInstanceResults = append(allInstanceResults, grabTwoInstances...)

	resultMap := make(reporter.MachineReport)
	sorter.SortIntoTeams(allInstanceResults, resultMap, true)

	log.Printf("Result map ...\n")
	resultPrettyPrint, _ := json.MarshalIndent(resultMap, "", "    ")
	log.Printf("%s\n", resultPrettyPrint)
}