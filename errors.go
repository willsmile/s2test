package main

import "errors"

var (
	// ErrEmptyPath is returned when the path argument is empty
	ErrEmptyPath = errors.New("path argument is empty")
	// ErrReadFile is returned when failed to read a file
	ErrReadFile = errors.New("failed to read a file")
	// ErrHTTPRequest is returned when failed to make a request by http client
	ErrHTTPRequest = errors.New("failed to create a http request")
	// ErrHTTPResponse is returned when failed to receive a response by http client
	ErrHTTPResponse = errors.New("failed to receive a http response")
	// ErrHTTPRespBody is returned when failed to read a http response body
	ErrHTTPRespBody = errors.New("failed to read a http response body")
	// ErrUndefinedAPI is returned when the target API is undefined
	ErrUndefinedAPI = errors.New("target API is undefined")
	// ErrEmptyReport is returned when report is empty
	ErrEmptyReport = errors.New("nothing to print due to empty report")
)
