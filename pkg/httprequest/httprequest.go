package httprequest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func PerformGetRequest(url string, apiKey string) {
	client := &http.Client{}

	// Create a new request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set the headers
	request.Header.Set("x-api-key", apiKey)
	request.Header.Set("Cache-Control", "no-cache")

	// Send the request
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Read the response body
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Response code", response.StatusCode)
}
