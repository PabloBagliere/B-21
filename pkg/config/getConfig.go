package config

import "sync"

var lock = &sync.Mutex{}

var instance *Config

// GetConfig returns the configuration instance
func GetConfig() (*Config, error) {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		instance, err := ReadFile()
		if err != nil {
			return nil, err
		}
		return instance, nil
	}
	return instance, nil
}
