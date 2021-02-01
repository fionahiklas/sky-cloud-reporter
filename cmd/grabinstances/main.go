package main

import (
	"encoding/json"
	"github.com/fionahiklas/sky-cloud-reporter/cloud/cloudone"
	"github.com/fionahiklas/sky-cloud-reporter/cloud/cloudtwo"
	"github.com/fionahiklas/sky-cloud-reporter/grab"
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

	log.Printf("Instances from one ...\n")
	onePrettyPrint, _ := json.MarshalIndent(grabOneInstances, "", "    ")
	log.Printf("%s\n", onePrettyPrint)

	log.Printf("Instances from two ...\n")
	twoPrettyPrint, _ := json.MarshalIndent(grabTwoInstances, "", "    ")
	log.Printf("%s\n", twoPrettyPrint)
}