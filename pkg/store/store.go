package store

import (
	"time"
)

func (s *Store[T]) unsafeSet(key string, value T) (err error) {
	s.cache[key] = &item[T]{
		value: value,
	}

	return nil
}

func (s *Store[T]) unsafeGet(key string) (value T, err error) {
	var (
		item *item[T]
		ok   bool
	)

	if item, ok = s.cache[key]; !ok {
		return *new(T), nil
	}

	if !item.expireAt.IsZero() && time.Now().After(item.expireAt) {
		delete(s.cache, key)

		return *new(T), nil
	}

	return item.value, nil
}

func (s *Store[T]) unsafeExpire(key string, exp time.Time) (succ bool, err error) {
	var (
		item *item[T]
		ok   bool
	)

	if item, ok = s.cache[key]; !ok {
		return false, nil
	}

	item.expireAt = exp

	return true, nil
}

func (s *Store[T]) unsafeDel(key string) (succ bool, err error) {
	if _, ok := s.cache[key]; !ok {
		return false, nil
	}

	delete(s.cache, key)

	return true, nil
}

func (s *Store[T]) unsafeIncrBy(key string, count int) (value T, err error) {
	zero := getZero(value)

	if value, err = s.unsafeGet(key); err != nil {
		return zero, err
	}

	var (
		addval Addable
		ok     bool
	)
	if value != zero {
		if addval, ok = any(value).(Addable); !ok {
			return zero, ErrorInvalidArgument
		}
	}

	addval = addval.Add(count)
	if value, ok = addval.(T); !ok {
		return zero, ErrorInvalidArgument
	}

	if err = s.unsafeSet(key, value); err != nil {
		return zero, err
	}

	return value, nil
}
