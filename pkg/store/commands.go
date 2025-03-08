package store

import (
	"time"
)

func (s *Store) Set(key string, value string) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.unsafeSet(key, value)
}

func (s *Store) SetExpire(key string, value string, exp time.Time) (err error) {
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

func (s *Store) Expire(key string, exp time.Time) (succ bool, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.unsafeExpire(key, exp)
}

func (s *Store) Get(key string) (value string, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.unsafeGet(key)
}

func (s *Store) Del(key string) (succ bool, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.unsafeDel(key)
}

func (s *Store) Incr(key string) (value string, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.unsafeIncrBy(key, 1)
}

func (s *Store) Decr(key string) (value string, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.unsafeIncrBy(key, -1)
}

func (s *Store) IncrBy(key string, count int) (value string, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.unsafeIncrBy(key, count)
}

func (s *Store) DecrBy(key string, count int) (value string, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.unsafeIncrBy(key, -count)
}
