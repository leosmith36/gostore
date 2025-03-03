package store

import "time"

type item[T comparable] struct {
	value    T
	expireAt time.Time
}

type Addable interface {
	Add(ops ...int) (sum Addable)
}
