package server

import (
	"log"
	"lsmith/go-store/internal/constants"
	"strings"
)

func executeCommand(input string) (output string, err error) {
	split := strings.Split(input, " ")

	if len(split) < 1 {
		return formatError(string(constants.ErrorMissingCommand)), nil
	}

	cmd := constants.InputKeyword(split[0])
	switch (cmd) {
	case constants.InputSet:
		log.Printf("SET %v", split[1:])
	}

	return "OK", nil
}