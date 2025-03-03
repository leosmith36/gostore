package store

import "time"

type item struct {
	value    string
	expireAt time.Time
}
