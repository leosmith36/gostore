package store

import "sync"

type Store struct {
	cache map[string]string
	mu *sync.RWMutex
}

func NewStore() (st *Store) {
	return &Store{
		cache: make(map[string]string),
		mu: &sync.RWMutex{},
	}
}