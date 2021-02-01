package main

import (
	"github.com/fionahiklas/sky-cloud-reporter/cloud/cloudone"
	"github.com/fionahiklas/sky-cloud-reporter/cloud/cloudtwo"
	"github.com/fionahiklas/sky-cloud-reporter/grab"
	"log"
	"net/http"
	"os"
)

func main() {

	cloudOneProviderFactory := func() grab.CloudProvider {
		return cloudone.NewProvider(os.Args[1])
	}

	cloudTwoProviderFactory := func() grab.CloudProvider {
		return cloudtwo.NewProvider(os.Args[2])
	}

	cloudFactories := []func() grab.CloudProvider {
		cloudOneProviderFactory,
		cloudTwoProviderFactory,
	}

	httpClientFactory := func() grab.HttpClient {
		return &http.Client{}
	}

	grabberFactory := func(client grab.HttpClient, provider grab.CloudProvider) grab.Grabber {
		return grab.NewGrabber(client, provider)
	}

 	grabHandler := NewGrabHandler(cloudFactories,grabberFactory,httpClientFactory)

 	log.Printf("Adding Handler ...")
 	http.HandleFunc("/report", grabHandler.HttpHandler())
 	log.Printf("Starting HTTP server ...")
 	log.Fatal(http.ListenAndServe(":8080", nil))
}