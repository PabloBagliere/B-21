package errors

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

// ConfigException is the error type for configuration errors
type ConfigException struct {
	Message string
	typeErr string
	Err     error
}

// Error returns the error message
func (e *ConfigException) Error() string {
	log.Error().Str("type", e.typeErr).Err(e.Err).Msg(e.Message)
	return e.Message
}

// formatMessage formatea el mensaje de error usando Zerolog
func formatMessage(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

// NewConfigError creates a new ConfigError
func NewConfigError(message string, err error, args ...interface{}) *ConfigException {
	return &ConfigException{Message: formatMessage(message, args...), Err: err, typeErr: "ConfigError"}
}

// IsConfigError checks if the error is a ConfigError
func IsConfigError(err error) bool {
	_, ok := err.(*ConfigException)
	return ok
}

// NewJWTError creates a new JWTError
func NewJWTError(message string, err error, args ...interface{}) *ConfigException {
	return &ConfigException{Message: formatMessage(message, args...), Err: err, typeErr: "JWTError"}
}

// IsJWTError checks if the error is a JWTError
func IsJWTError(err error) bool {
	_, ok := err.(*ConfigException)
	return ok
}
