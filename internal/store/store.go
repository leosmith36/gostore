package store

import "time"

func (s *Store) unsafeSet(key, value string) (err error) {
	s.cache[key] = &cacheItem{
		value: value,
	}

	return nil
}

func (s *Store) unsafeGet(key string) (value string, err error) {
	var (
		item *cacheItem
		ok   bool
	)

	if item, ok = s.cache[key]; !ok {
		return "", nil
	}

	if !item.expireAt.IsZero() && time.Now().After(item.expireAt) {
		delete(s.cache, key)

		return "", nil
	}

	return item.value, nil
}

func (s *Store) unsafeExpire(key string, exp time.Time) (succ bool, err error) {
	var (
		item *cacheItem
		ok   bool
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
