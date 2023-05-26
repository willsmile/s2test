package main

import (
	"errors"
	"testing"
)

func TestHTTPRequest_ValidRequest(t *testing.T) {
	method := "GET"
	url := "https://jsonplaceholder.typicode.com/posts/1"
	headers := map[string]string{}
	auth := cookie{}

	_, err := HTTPRequest(method, url, headers, auth)
	if err != nil {
		t.Fatalf("HTTPRequest(method, url, headers, auth), expected none error, got %s", err)
	}
}

func TestHTTPRequest_InvalidRequest(t *testing.T) {
	method := "*?"
	url := "https://jsonplaceholder.typicode.com/posts/1"
	headers := map[string]string{}
	auth := cookie{}

	_, err := HTTPRequest(method, url, headers, auth)
	if !errors.Is(err, ErrHTTPRequest) {
		t.Fatalf("HTTPRequest(method, url, headers, auth), expected %s, got %s", ErrHTTPRequest, err)
	}
}
