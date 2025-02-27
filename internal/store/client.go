package store

import (
	"context"
	"sync"
)

type Store struct {
	cache  map[string]*cacheItem
	mu     sync.RWMutex
	wg     *sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

func NewStore() (st *Store) {
	ctx, cancel := context.WithCancel(context.Background())

	return &Store{
		cache:  make(map[string]*cacheItem),
		wg:     &sync.WaitGroup{},
		ctx:    ctx,
		cancel: cancel,
	}
}

func (s *Store) Start() {
	// stub
}

func (s *Store) Stop() {
	s.cancel()
	s.wg.Wait()
}
