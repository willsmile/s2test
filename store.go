package main

// Store is a store for API specs
type Store map[string]spec

func (store *Store) Search(target string) (spec, error) {
	s := (*store)[target]
	if s.available() {
		return s, nil
	} else {
		return spec{}, ErrUndefinedAPI
	}
}
