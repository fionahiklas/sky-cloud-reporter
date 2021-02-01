package main

import (
	"github.com/fionahiklas/sky-cloud-reporter/common/reporter"
	"github.com/fionahiklas/sky-cloud-reporter/grab"
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

		// TODO: This might need to be bigger and also needs to expand
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


	}
}