package main

// Store is a store for API endpoints
type Store map[string]Endpoint

func (store *Store) Search(target string) (Endpoint, error) {
	s := (*store)[target]
	if s.available() {
		return s, nil
	} else {
		return Endpoint{}, ErrUndefinedAPI
	}
}
