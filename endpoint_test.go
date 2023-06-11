package main

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestSearch_Available(t *testing.T) {
	store := Endpoints{
		"GET a sample post": Endpoint{
			URL:    "https://jsonplaceholder.typicode.com/posts/1",
			Method: "GET",
			Headers: map[string]string{
				"Content-type": "application/json; charset=utf-8",
			},
		},
		"GET a sample todo": Endpoint{
			URL:    "https://jsonplaceholder.typicode.com/todos/1/",
			Method: "GET",
			Headers: map[string]string{
				"Content-type": "application/json; charset=utf-8",
			},
		},
	}

	_, err := store.Search("GET a sample post")
	if err != nil {
		t.Fatalf("store.Search(\"GET a sample post\"), expected none error, got %s", err)
	}
}

func TestSearch_NotAvailable(t *testing.T) {
	store := Endpoints{
		"GET a sample post": Endpoint{
			URL:    "https://jsonplaceholder.typicode.com/posts/1",
			Method: "GET",
			Headers: map[string]string{
				"Content-type": "application/json; charset=utf-8",
			},
		},
		"GET a sample todo": Endpoint{
			URL:    "https://jsonplaceholder.typicode.com/todos/1/",
			Method: "GET",
			Headers: map[string]string{
				"Content-type": "application/json; charset=utf-8",
			},
		},
	}

	_, err := store.Search("Not Available API")
	if !errors.Is(err, ErrUndefinedAPI) {
		t.Fatalf("store.Search(\"Not Available API\"), expected %s, got %s", ErrUndefinedAPI, err)
	}
}

func TestSearch_Empty(t *testing.T) {
	store := Endpoints{}

	_, err := store.Search("API from Empty Store")
	if !errors.Is(err, ErrUndefinedAPI) {
		t.Fatalf("store.Search(\"API from Empty Store\"), expected %s, got %s", ErrUndefinedAPI, err)
	}
}

func TestNewRequest_ValidEndpoint(t *testing.T) {
	s := Endpoint{
		Method:  "GET",
		URL:     "https://jsonplaceholder.typicode.com/posts/1",
		Headers: map[string]string{},
		Body:    json.RawMessage{},
	}
	auth := cookie{}
	data := CustomizedData{}

	_, err := s.NewRequest(auth, data)
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
	data := CustomizedData{}

	_, err := s.NewRequest(auth, data)
	if err == nil {
		t.Fatalf("s.NewRequest(auth), expected error %s, got none", err)
	}
}
