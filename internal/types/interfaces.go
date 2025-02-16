package types

import "time"

type StringCache interface {
	Get(key string) (value string, err error)
	Set(key, value string) (err error)
	SetExpire(key, value string, exp time.Time) (err error)
	Del(key string) (succ bool, err error)
	Expire(key string, exp time.Time) (succ bool, err error)
}