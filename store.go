package main

// Store is a store for API specs
type Store map[string]spec

// spec is for information of a single API
type spec struct {
	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
}
