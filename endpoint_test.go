package main

import (
	"encoding/json"
	"testing"
)

func TestNewRequest_ValidEndpoint(t *testing.T) {
	s := Endpoint{
		Method:  "GET",
		URL:     "https://jsonplaceholder.typicode.com/posts/1",
		Headers: map[string]string{},
		Body:    json.RawMessage{},
	}
	auth := cookie{}

	_, err := s.NewRequest(auth)
	if err != nil {
		t.Fatalf("s.NewRequest(auth), expected none error, got %s", err)
	}
}

func TestNewRequest_InvalidEndpoint(t *testing.T) {
	s := Endpoint{
		Method:  "*?",
		URL:     "https://jsonplaceholder.typicode.com/posts/1",
		Headers: map[string]string{},
		Body:    json.RawMessage{},
	}
	auth := cookie{}

	_, err := s.NewRequest(auth)
	if err == nil {
		t.Fatalf("s.NewRequest(auth), expected error %s, got none", err)
	}
}
