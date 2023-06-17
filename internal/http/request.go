package http

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	"github.com/willsmile/s2test/internal/depository"
)

var (
	// ErrHTTPRequest is returned when failed to make a request by http client
	ErrHTTPRequest = errors.New("failed to create a http request")
	// ErrHTTPResponse is returned when failed to receive a response by http client
	ErrHTTPResponse = errors.New("failed to receive a http response")
	// ErrHTTPRespBody is returned when failed to read a http response body
	ErrHTTPRespBody = errors.New("failed to read a http response body")
	// ErrUndefinedAPI is returned when the target API is undefined
	ErrUndefinedAPI = errors.New("target API is undefined")
)

type Request struct {
	Endpoint depository.Endpoint
	Auth     depository.AuthInfo
	Data     depository.CustomizedData
}

func NewRequest(e depository.Endpoint, a depository.AuthInfo, d depository.CustomizedData) *Request {
	return &Request{e, a, d}
}

func NewHTTPClient() *http.Client {
	return http.DefaultClient
}

func (req *Request) HTTPRequest() (*http.Request, error) {
	var (
		hreq *http.Request
		err  error
	)

	// Check endpoint whether is available
	if !req.Endpoint.Available() {
		return nil, ErrUndefinedAPI
	}

	// Prepare a request
	if req.Endpoint.Method == http.MethodPost {
		body := req.Data.Apply(req.Endpoint.Body)
		buf := bytes.NewBufferString(body)
		hreq, err = http.NewRequest(req.Endpoint.Method, req.Endpoint.URL, buf)
	} else {
		hreq, err = http.NewRequest(req.Endpoint.Method, req.Endpoint.URL, nil)
	}
	if err != nil {
		return nil, ErrHTTPRequest
	}

	// Add headers to request if exists
	if len(req.Endpoint.Headers) != 0 {
		for key, value := range req.Endpoint.Headers {
			hreq.Header.Add(key, value)
		}
	}

	// Attach AuthInfo to request if exists
	if req.Auth != nil {
		req.Auth.Attach(hreq)
	}

	return hreq, nil
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
