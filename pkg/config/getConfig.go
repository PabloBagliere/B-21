package config

import (
	"github.com/PabloBagliere/B-21/pkg/errors"
)

// GetConfig returns the configuration instance
func GetConfig(nameService string) (map[string]interface{}, error) {
	config, error := ReadFile()
	if error != nil {
		return nil, error
	}
	serviceConfig, ok := config.Services[nameService].(map[string]interface{})
	if !ok {
		return nil, errors.NewConfigError("Error getting the service configuration", nil)
	}
	return serviceConfig, nil
}
