package constants

type ErrorType string

const (
	ErrorInternal ErrorType = "internal server error"
	ErrorMissingCommand ErrorType = "no command provided"
)