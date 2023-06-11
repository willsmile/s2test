package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Endpoints is a store for API endpoints
type Endpoints map[string]Endpoint

// Endpoint is for information of a single API endpoint
type Endpoint struct {
	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Body    json.RawMessage   `json:"body"`
}

func (store *Endpoints) Search(target string) (Endpoint, error) {
	s := (*store)[target]
	if s.available() {
		return s, nil
	} else {
		return Endpoint{}, ErrUndefinedAPI
	}
}

func (e *Endpoint) NewRequest(auth AuthInfo, data CustomizedData) (*http.Request, error) {
	var (
		req *http.Request
		err error
	)

	// Prepare a request
	if e.Method == http.MethodPost {
		body := data.Apply(e.Body)
		buf := bytes.NewBufferString(body)
		req, err = http.NewRequest(e.Method, e.URL, buf)
	} else {
		req, err = http.NewRequest(e.Method, e.URL, nil)
	}
	if err != nil {
		return nil, err
	}

	// Add headers to request if exists
	if len(e.Headers) != 0 {
		for key, value := range e.Headers {
			req.Header.Add(key, value)
		}
	}

	// Attach AuthInfo to request if exists
	if auth != nil {
		auth.Attach(req)
	}

	return req, nil
}

func (e Endpoint) available() bool {
	if e.URL != "" && e.Method != "" {
		return true
	} else {
		return false
	}
}
