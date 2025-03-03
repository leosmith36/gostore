package store

import (
	"time"
)

func (s *Store[T]) Set(key string, value T) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.unsafeSet(key, value)
}

func (s *Store[T]) SetExpire(key string, value T, exp time.Time) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err = s.unsafeSet(key, value); err != nil {
		return err
	}

	if _, err = s.unsafeExpire(key, exp); err != nil {
		return err
	}

	return nil
}

func (s *Store[T]) Expire(key string, exp time.Time) (succ bool, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.unsafeExpire(key, exp)
}

func (s *Store[T]) Get(key string) (value T, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.unsafeGet(key)
}

func (s *Store[T]) Del(key string) (succ bool, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.unsafeDel(key)
}

func (s *Store[T]) Incr(key string) (value T, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.unsafeIncrBy(key, 1)
}

func (s *Store[T]) Decr(key string) (value T, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.unsafeIncrBy(key, -1)
}

func (s *Store[T]) IncrBy(key string, count int) (value T, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.unsafeIncrBy(key, count)
}

func (s *Store[T]) DecrBy(key string, count int) (value T, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.unsafeIncrBy(key, -count)
}
