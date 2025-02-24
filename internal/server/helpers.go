package server

import (
	"errors"
	"fmt"
	"lsmith/gostore/internal/constants"
	"strconv"
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

	return time.Now().Add(time.Duration(secs)*time.Second), nil
}
