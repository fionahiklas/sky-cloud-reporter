package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}
	response, err := client.Get(os.Args[1])

	if err == nil {
		defer response.Body.Close()

		log.Printf("Status: %s\n", response.Status)
		log.Printf("Status code: %d\n", response.StatusCode)
		body, err := ioutil.ReadAll(response.Body)
		if err == nil {
			log.Printf("Body: %s\n", body)
		} else {
			log.Fatalf("Error getting body: %s", err)
		}
	} else {
		log.Fatalf("Failed get: %s", err.Error())
	}
}
