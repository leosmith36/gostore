package server

import (
	"fmt"
	"lsmith/gostore/internal/constants"
	"lsmith/gostore/internal/types"
	"time"
)

func set(st types.StringCache, args ...string) (output string) {
  if len(args) < 1{
    return formatError(constants.ErrorMissingArguments)
  }
  if len(args) < 2 {
    return formatError("missing value for SET")
  }

	key := args[0]
  value := args[1]

	var (
		exp time.Time
		err error
	)

	if len(args) > 2 {
		for i := 2; i < len(args); i++ {
			arg := args[i]
			switch (arg) {
			case constants.OptionEx:
				exp, err = parseExpiration(args[i+1:]...)
				i++
			default:
				return formatError(fmt.Sprintf("unknown option: %s", arg))
			}

			if err != nil {
				return formatError(err.Error())
			}
		}
	}

	if !exp.IsZero() {
		if err = st.SetExpire(key, value, exp); err != nil {
			return formatError(err.Error())
		}
	} else if err = st.Set(key, value); err != nil {
		return formatError(err.Error())
	}


  return formatOutput(constants.OutputOk)
}

func get(st types.StringCache, args ...string) (output string) {
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
		return formatOutput(constants.OutputNull)
	}

	return formatOutput(fmt.Sprintf(`"%s"`, value))
}

func del(st types.StringCache, args ...string) (output string) {
  if len(args) < 1 {
    return formatError(constants.ErrorMissingArguments)
  }

  key := args[0]

  if _, err := st.Del(key); err != nil {
    return formatError(err.Error())
  }

	return formatOutput(constants.OutputOk)
}