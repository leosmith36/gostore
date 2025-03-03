package store

import (
	"context"
	"sync"
)

type Store[T comparable] struct {
	cache  map[string]*item[T]
	mu     sync.RWMutex
	wg     *sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

func NewStore[T comparable]() (st *Store[T]) {
	ctx, cancel := context.WithCancel(context.Background())

	return &Store[T]{
		cache:  make(map[string]*item[T]),
		wg:     &sync.WaitGroup{},
		ctx:    ctx,
		cancel: cancel,
	}
}

func (s *Store[T]) Start() {
	// stub
}

func (s *Store[T]) Stop() {
	s.cancel()
	s.wg.Wait()
}
