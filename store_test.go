package main

import (
	"errors"
	"testing"
)

func TestStoreSearch_Available(t *testing.T) {
	store := Store{
		"GET a sample post": spec{
			URL:    "https://jsonplaceholder.typicode.com/posts/1",
			Method: "GET",
			Headers: map[string]string{
				"Content-type": "application/json; charset=utf-8",
			},
		},
		"GET a sample todo": spec{
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

func TestStoreSearch_NotAvailable(t *testing.T) {
	store := Store{
		"GET a sample post": spec{
			URL:    "https://jsonplaceholder.typicode.com/posts/1",
			Method: "GET",
			Headers: map[string]string{
				"Content-type": "application/json; charset=utf-8",
			},
		},
		"GET a sample todo": spec{
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

func TestStoreSearch_EmptyStore(t *testing.T) {
	store := Store{}

	_, err := store.Search("API from Empty Store")
	if !errors.Is(err, ErrUndefinedAPI) {
		t.Fatalf("store.Search(\"API from Empty Store\"), expected %s, got %s", ErrUndefinedAPI, err)
	}
}
