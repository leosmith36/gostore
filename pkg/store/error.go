package store

import "fmt"

type StoreError error

var (
	ErrorNotAnInteger StoreError = fmt.Errorf("value is not an integer")
)
