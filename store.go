package main

// Store is a store for API specs
type Store map[string]spec

// spec is for information of a single API
type spec struct {
	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
}

func (store *Store) Search(target string) (spec, error) {
	s := (*store)[target]
	if s.available() {
		return s, nil
	} else {
		return spec{}, ErrUndefinedAPI
	}
}

func (s spec) available() bool {
	if s.URL != "" && s.Method != "" {
		return true
	} else {
		return false
	}
}
