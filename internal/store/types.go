package store

import "time"

type cacheItem struct {
	value string
	expireAt time.Time
}