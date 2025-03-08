package server

import (
	"fmt"
	"lsmith/gostore/internal/constants"
	"lsmith/gostore/internal/types"
	"strconv"
	"time"
)

func set(st types.KeyValueStore, args ...string) (output string) {
	if len(args) < 2 {
		return formatError(constants.ErrorMissingArguments)
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
			switch arg {
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

func get(st types.KeyValueStore, args ...string) (output string) {
	if len(args) < 1 {
		return formatError(constants.ErrorMissingArguments)
	}
	if len(args) > 1 {
		return formatError(constants.ErrorTooManyArguments)
	}

	key := args[0]

	var (
		value string
		err   error
	)

	if value, err = st.Get(key); err != nil {
		return formatError(err.Error())
	}

	if value == "" {
		return formatOutput(constants.OutputNull)
	}

	return formatOutput(fmt.Sprintf(`"%s"`, value))
}

func del(st types.KeyValueStore, args ...string) (output string) {
	if len(args) < 1 {
		return formatError(constants.ErrorMissingArguments)
	}

	var (
		key    = args[0]
		exists bool
		res    int
		err    error
	)

	if exists, err = st.Del(key); err != nil {
		return formatError(err.Error())
	}

	if exists {
		res = 1
	} else {
		res = 0
	}

	return formatOutput(fmt.Sprint(res))
}

func ping(args ...string) (output string) {
	if len(args) < 1 {
		return formatOutput("PONG")
	}
	if len(args) > 1 {
		return formatError(constants.ErrorTooManyArguments)
	}

	return formatOutput(fmt.Sprintf(`"%s"`, args[0]))
}

func incr(st types.KeyValueStore, decr bool, args ...string) (output string) {
	if len(args) < 1 {
		return formatError(constants.ErrorMissingArguments)
	}
	if len(args) > 1 {
		return formatError(constants.ErrorTooManyArguments)
	}

	key := args[0]

	var (
		value string
		err   error
		count int
	)

	if decr {
		count = -1
	} else {
		count = 1
	}

	if value, err = st.IncrBy(key, count); err != nil {
		return formatError(err.Error())
	}

	return formatOutput(fmt.Sprintf(`"%s"`, value))
}

func incrBy(st types.KeyValueStore, decr bool, args ...string) (output string) {
	if len(args) < 2 {
		return formatError(constants.ErrorMissingArguments)
	}
	if len(args) > 2 {
		return formatError(constants.ErrorTooManyArguments)
	}

	key := args[0]
	scount := args[1]

	var (
		value string
		err   error
		count int
	)

	if count, err = strconv.Atoi(scount); err != nil {
		return formatError(constants.ErrInvalidArguments)
	}

	if decr {
		count *= -1
	}

	if value, err = st.IncrBy(key, count); err != nil {
		return formatError(err.Error())
	}

	return formatOutput(fmt.Sprintf(`"%s"`, value))
}
