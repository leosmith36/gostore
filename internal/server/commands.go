package server

import (
	"fmt"
	"lsmith/gostore/internal/constants"
	"lsmith/gostore/internal/store"
)

func set(st *store.Store, args ...string) (output string) {
  if len(args) < 1{
    return formatError(constants.ErrorMissingArguments)
  }
  if len(args) < 2 {
    return formatError("missing value for SET command")
  }

  key := args[0]
  value := args[1]

  if err := st.Set(key, value); err != nil {
    return formatError(err.Error())
  }

  return formatOutput(constants.OutOk)
}

func get(st *store.Store, args ...string) (output string) {
  if len(args) < 1 {
    return formatError(constants.ErrorMissingArguments)
  }

  key := args[0]

	var (
		value string
		err error
	)

  if value, err = st.Get(key); err != nil {
    return formatError(err.Error())
  }

  if value == "" {
		return formatOutput(constants.OutNull)
	}

	return formatOutput(fmt.Sprintf(`"%s"`, value))
}

func del(st *store.Store, args ...string) (output string) {
  if len(args) < 1 {
    return formatError(constants.ErrorMissingArguments)
  }

  key := args[0]

  if _, err := st.Del(key); err != nil {
    return formatError(err.Error())
  }

	return formatOutput(constants.OutOk)
}