package main

import "errors"

var (
	// ErrEmptyPath is returned when the path argument is empty
	ErrEmptyPath = errors.New("Path argument is empty")
	// ErrReadFile is returned when failed to read a file
	ErrReadFile = errors.New("Failed to read a file")
	// ErrHTTPRequest is returned when failed to make a request by http client
	ErrHTTPRequest = errors.New("Failed to create a http request")
	// ErrHTTPResponse is returned when failed to receive a response by http client
	ErrHTTPResponse = errors.New("Failed to receive a http response")
	// ErrHTTPRespBody is returned when failed to read a http response body
	ErrHTTPRespBody = errors.New("Failed to read a http response body")
)
