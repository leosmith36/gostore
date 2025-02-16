package store

import (
	"time"
)

func (s *Store) Set(key, value string) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.unsafeSet(key, value)
}

func (s *Store) SetExpire(key, value string, exp time.Time) (err error) {
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

func (s *Store) unsafeSet(key, value string) (err error) {
	s.cache[key] = &cacheItem{
		value: value,
	}

	return nil
}

func (s *Store) unsafeGet(key string) (value string, err error) {
	var (
		item *cacheItem
		ok bool
	)

	if item, ok = s.cache[key]; !ok {
		return "", nil
	}

	return item.value, nil
}


func (s *Store) unsafeExpire(key string, exp time.Time) (succ bool, err error) {
	var (
		item *cacheItem
		ok bool
	)

	if item, ok = s.cache[key]; !ok {
		return false, nil
	}

	item.expireAt = exp

	return true, nil
}

func (s *Store) unsafeDel(key string) (succ bool, err error) {
	if _, ok := s.cache[key]; !ok {
		return false, nil
	}

	delete(s.cache, key)

	return true, nil
}