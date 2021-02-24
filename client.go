package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// HTTPRequest sends a HTTP request
func HTTPRequest(method string, url string, headers map[string]string, cookies map[string]string) (string, string) {
	// Prepare request
	request, error := http.NewRequest(method, url, nil)
	// Add headers to request if exists
	if len(headers) != 0 {
		for key, value := range headers {
			request.Header.Add(key, value)
		}
	}
	// Add cookies to request if exists
	if len(cookies) != 0 {
		for key, value := range cookies {
			cookie := &http.Cookie{Name: key, Value: value}
			request.AddCookie(cookie)
		}
	}
	if error != nil {
		log.Fatal("[HTTP Request Error] ", error)
	}

	// Send request by client
	client := http.DefaultClient
	response, error := client.Do(request)
	if error != nil {
		log.Fatal("[HTTP Response Error] ", error)
	}
	defer response.Body.Close()

	// Read request body and status
	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		log.Fatal("[IO Error] ", error)
	}
	status := response.Status

	return string(body), string(status)
}
