package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// spec is for information of a single API
type spec struct {
	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Body    json.RawMessage   `json:"body"`
}

func (s *spec) NewRequest(auth AuthInfo) (*http.Request, error) {
	var (
		req *http.Request
		err error
	)

	// Prepare a request
	if s.Method == http.MethodPost {
		buf := bytes.NewBuffer(s.Body)
		req, err = http.NewRequest(s.Method, s.URL, buf)
	} else {
		req, err = http.NewRequest(s.Method, s.URL, nil)
	}
	if err != nil {
		return nil, err
	}

	// Add headers to request if exists
	if len(s.Headers) != 0 {
		for key, value := range s.Headers {
			req.Header.Add(key, value)
		}
	}

	// Attach AuthInfo to request if exists
	if auth != nil {
		auth.Attach(req)
	}

	return req, nil
}

func (s spec) available() bool {
	if s.URL != "" && s.Method != "" {
		return true
	} else {
		return false
	}
}
