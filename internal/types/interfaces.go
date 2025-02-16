package types

type StringCache interface {
	Get(key string) (value string, err error)
	Set(key, value string) (err error)
	Del(key string) (succ bool, err error)
}