package server

import (
	"errors"
	"fmt"
	"lsmith/gostore/internal/constants"
	"strconv"
	"strings"
	"time"
)

func formatOutput(msg string) (output string) {
	return fmt.Sprintf("%s\n", msg)
}

func formatOutputWithKeyword(kwd string, msg string) (output string) {
	return fmt.Sprintf("%s %s\n", kwd, msg)
}

func formatError(msg string) (output string) {
	return formatOutputWithKeyword(constants.OutputError, msg)
}

func parseExpiration(args ...string) (exp time.Time, err error) {
	if len(args) < 1 {
		return time.Time{}, errors.New("missing seconds for EX")
	}

	var secs int
	if secs, err = strconv.Atoi(args[0]); err != nil {
		return time.Time{}, fmt.Errorf("invalid seconds for EX: %s", args[0])
	}

	if secs <= 0 {
		return time.Time{}, fmt.Errorf("invalid seconds for EX: %d", secs)
	}

	return time.Now().Add(time.Duration(secs) * time.Second), nil
}

func splitArgs(raw string) (args []string) {
	var (
		cur    string
		quoted bool
	)

	split := strings.Split(raw, "")
	for _, c := range split {
		switch c {
		case " ":
			if quoted {
				break
			}
			args = append(args, cur)
			cur = ""
			continue
		case "\"":
			if quoted {
				quoted = false
				args = append(args, cur)
				cur = ""
			} else {
				quoted = true
			}
			continue
		}

		cur += string(c)
	}

	if cur != "" {
		args = append(args, cur)
	}

	return
}
