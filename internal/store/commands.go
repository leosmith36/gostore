package store

func (s *Store) Set(key, value string) (err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.unsafeSet(key, value)
}

func (s *Store) unsafeSet(key, value string) (err error) {
	s.cache[key] = value

	return nil
}

func (s *Store) Get(key string) (value string, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.unsafeGet(key)
}

func (s *Store) unsafeGet(key string) (value string, err error) {
	return s.cache[key], nil
}

func (s *Store) Del(key string) (succ bool, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.unsafeDel(key)
}

func (s *Store) unsafeDel(key string) (succ bool, err error) {
	if _, ok := s.cache[key]; !ok {
		return false, nil
	}

	delete(s.cache, key)

	return true, nil
}