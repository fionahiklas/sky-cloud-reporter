package main

import (
	"encoding/json"
	"fmt"
	"github.com/fionahiklas/sky-cloud-reporter/common/reporter"
	"github.com/fionahiklas/sky-cloud-reporter/grab"
	"github.com/fionahiklas/sky-cloud-reporter/sorter"
	"log"
	"net/http"
)

type grabHandler struct {
	CloudProviderFactories []func() grab.CloudProvider
	GrabberFactory func(client grab.HttpClient, provider grab.CloudProvider) grab.Grabber
	HttpClientFactory func() grab.HttpClient
}

func NewGrabHandler(cloudProvidersFactories []func() grab.CloudProvider,
					grabFactory func(client grab.HttpClient, provider grab.CloudProvider) grab.Grabber,
					httpClientFactory func() grab.HttpClient) *grabHandler {
	return &grabHandler{
		CloudProviderFactories: cloudProvidersFactories,
		GrabberFactory: grabFactory,
		HttpClientFactory: httpClientFactory,
	}
}

func (handler *grabHandler) HttpHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Handler function called")
		instanceData := make([]reporter.MachineInstance,0, 20)

		for _, providerFactory := range handler.CloudProviderFactories {
			log.Printf("Creating Provider")
			provider := providerFactory()
			log.Printf("Created provider: %T", provider)
			httpClient := handler.HttpClientFactory()
			log.Printf("Creating grabber")
			grabber := handler.GrabberFactory(httpClient, provider)
			log.Printf("Grabbing instance information")
			providerInstances, _ := grabber.GrabInstances()
			instanceData = append(instanceData, providerInstances...)
		}

		log.Printf("Sorting data into Report ...")
		resultMap := make(reporter.MachineReport)
		sorter.SortIntoTeams(instanceData, resultMap, true)

		log.Printf("Result map ...\n")
		resultPrettyPrint, _ := json.MarshalIndent(resultMap, "", "    ")
		log.Printf("Returning result:\n\n%s\n", resultPrettyPrint)

		fmt.Fprintf(w, "%s\n", resultPrettyPrint)
		log.Printf("Done")
	}
}