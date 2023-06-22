package http

import (
	"io"
	"net/http"
)

func NewClient() *http.Client {
	return http.DefaultClient
}

func Send(req *http.Request, client *http.Client) (*Response, error) {
	response := &Response{}

	// Send HTTP request by client
	resp, err := client.Do(req)
	if err != nil {
		return response, ErrHTTPResponse
	}
	defer resp.Body.Close()

	// Read HTTP request body and status
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, ErrHTTPRespBody
	}

	// Save body and status to response
	response.Body = string(body)
	response.Status = string(resp.Status)

	return response, nil
}
