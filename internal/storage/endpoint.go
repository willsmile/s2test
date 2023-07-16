package storage

import (
	"encoding/json"
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

func (epts Endpoints) GetEndpoint(target string) Endpoint {
	return epts[target]
}
