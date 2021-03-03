package main

import (
	"io/ioutil"
	"net/http"
)

// Response constains status and body of a http request
type Response struct {
	Body   string
	Status string
}

// NewResponse constructs an empty response.
func NewResponse() *Response {
	return &Response{}
}

// HTTPRequest sends a HTTP request
func HTTPRequest(method string, url string, headers map[string]string, cookies map[string]string) (*Response, error) {
	response := NewResponse()

	// Prepare request
	req, err := http.NewRequest(method, url, nil)
	// Add headers to request if exists
	if len(headers) != 0 {
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	}
	// Add cookies to request if exists
	if len(cookies) != 0 {
		for key, value := range cookies {
			cookie := &http.Cookie{Name: key, Value: value}
			req.AddCookie(cookie)
		}
	}
	if err != nil {
		return response, ErrHTTPRequest
	}

	// Send request by client
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return response, ErrHTTPResponse
	}
	defer resp.Body.Close()

	// Read request body and status
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, ErrHTTPRespBody
	}

	response.Body = string(body)
	response.Status = string(resp.Status)

	return response, nil
}
