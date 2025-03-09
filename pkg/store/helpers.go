package store

import "time"

func newItem(key, value string) *item {
	return &item{
		key:   key,
		value: value,

		lastUsedAt: time.Now(),
		numUses:    1,
	}
}
