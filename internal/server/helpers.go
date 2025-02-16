package server

import (
	"fmt"
	"lsmith/go-store/internal/constants"
)

func formatOutput(kwd constants.OutputKeyword, msg string) (output string) {
	return fmt.Sprintf("%s %s\n", kwd, msg)
}

func formatError(msg string) (output string) {
	return formatOutput(constants.OutputError, msg)
}