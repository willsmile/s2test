package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/willsmile/s2test/internal/storage"
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
	URL     string
	Method  string
	Headers http.Header
	Cookies []*http.Cookie
	Body    string
}

func NewRequest(ept storage.Endpoint, info []AuthInfo, vbs Variables, ua string) *Request {
	req := &Request{
		URL:     ept.URL,
		Method:  ept.Method,
		Headers: http.Header{},
		Cookies: []*http.Cookie{},
		Body:    "",
	}

	req.addHeaders(ept.Headers)
	if len(info) != 0 {
		for _, i := range info {
			i.attach(req)
		}
	}
	req.setBody(ept.Body, vbs)
	req.setUserAgent(ua)

	return req
}

// Add headers to request if exists
func (req *Request) addHeaders(headers map[string]string) {
	if len(headers) != 0 {
		for key, value := range headers {
			req.Headers.Add(key, value)
		}
	}
}

// Set body from raw and customized data of body
func (req *Request) setBody(raw json.RawMessage, vbs Variables) {
	var body string
	rawBody := string(raw)
	if len(vbs) != 0 {
		replacer := vbs.newReplacer()
		body = replacer.Replace(rawBody)
	} else {
		body = rawBody
	}
	req.Body = body
}

func (req *Request) setUserAgent(ua string) {
	req.Headers.Set("User-Agent", ua)
}

func (req *Request) HTTPRequest() (*http.Request, error) {
	var (
		hreq *http.Request
		err  error
	)

	// Check whether endpoint is available or not
	if !req.isAvailable() {
		return nil, ErrUndefinedAPI
	}

	// Prepare a request
	if req.Method == http.MethodPost {
		buf := bytes.NewBufferString(req.Body)
		hreq, err = http.NewRequest(req.Method, req.URL, buf)
	} else {
		hreq, err = http.NewRequest(req.Method, req.URL, nil)
	}
	if err != nil {
		return nil, ErrHTTPRequest
	}

	// Add headers to request if exists
	hreq.Header = req.Headers
	// Add cookies to request if exists
	for _, cookie := range req.Cookies {
		hreq.AddCookie(cookie)
	}

	return hreq, nil
}

func (req *Request) isAvailable() bool {
	if req.URL != "" && req.Method != "" {
		return true
	} else {
		return false
	}
}
