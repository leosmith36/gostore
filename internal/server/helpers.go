package server

import (
	"fmt"
	"lsmith/gostore/internal/constants"
)

func formatOutput(msg string) (output string) {
	return fmt.Sprintf("%s\n", msg)
}

func formatOutputWithKeyword(kwd string, msg string) (output string) {
	return fmt.Sprintf("%s %s\n", kwd, msg)
}

func formatError(msg string) (output string) {
	return formatOutputWithKeyword(constants.OutError, msg)
}