package http

import (
	"io"
	"net/http"
)

func NewHTTPClient() *http.Client {
	return http.DefaultClient
}

func SendHTTPRequest(req *http.Request, client *http.Client) (*Response, error) {
	response := NewResponse()

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
