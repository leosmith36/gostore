package store

import "time"

type item struct {
	key   string
	value string

	expireAt   time.Time
	lastUsedAt time.Time
	numUses    int
}
