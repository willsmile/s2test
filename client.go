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

func DefaultResponse() *Response {
	return &Response{
		Body:   "None",
		Status: "None",
	}
}

// HTTPRequest sends a HTTP request
func HTTPRequest(method string, url string, headers map[string]string, auth AuthInfo) (*Response, error) {
	response := &Response{}

	// Prepare request
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return response, ErrHTTPRequest
	}

	// Add headers to request if exists
	if len(headers) != 0 {
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	}
	// Attach AuthInfo to request if exists
	if auth != nil {
		auth.Attach(req)
	}

	// Send request by client
	client := http.DefaultClient
	resp, err := client.Do(req)
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
