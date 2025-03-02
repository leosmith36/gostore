package store

import (
	"fmt"
	"strconv"
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

func (s *Store) IncrBy(key string, count int) (value string, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if value, err = s.unsafeGet(key); err != nil {
		return "", err
	}

	var ivalue int
	if value == "" {
		ivalue = 0
	} else if ivalue, err = strconv.Atoi(value); err != nil {
		return "", ErrorNotAnInteger
	}

	ivalue += count
	if err = s.unsafeSet(key, fmt.Sprint(ivalue)); err != nil {
		return "", err
	}

	return fmt.Sprint(ivalue), nil
}
