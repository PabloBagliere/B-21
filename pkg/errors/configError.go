package errors

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

// ConfigError is the error type for configuration errors
type ConfigError struct {
	Message string
	Err     error
}

// Error returns the error message
func (e *ConfigError) Error() string {
	log.Error().Str("type", "CustomError").Err(e.Err).Msg(e.Message)
	return e.Message
}

// formatMessage formatea el mensaje de error usando Zerolog
func formatMessage(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

// NewConfigError creates a new ConfigError
func NewConfigError(message string, err error, args ...interface{}) *ConfigError {
	return &ConfigError{Message: formatMessage(message, args...), Err: err}
}

func IsConfigError(err error) bool {
	_, ok := err.(*ConfigError)
	return ok
}
