package main

import (
	"io"
	"net/http"
)

// Response constains status and body of a http request
type Response struct {
	Body   string
	Status string
}

// DefaultResponse create a response with default values
func DefaultResponse() *Response {
	return &Response{
		Body:   "None",
		Status: "None",
	}
}

// HTTPRequest sends a HTTP request
func HTTPRequest(s spec, auth AuthInfo) (*Response, error) {
	response := &Response{}

	// Create a request by spec
	request, err := s.NewRequest(auth)
	if err != nil {
		return response, ErrHTTPRequest
	}

	// Send request by client
	client := http.DefaultClient
	resp, err := client.Do(request)
	if err != nil {
		return response, ErrHTTPResponse
	}
	defer resp.Body.Close()

	// Read request body and status
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, ErrHTTPRespBody
	}

	response.Body = string(body)
	response.Status = string(resp.Status)

	return response, nil
}
