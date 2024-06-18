package config

import "github.com/BurntSushi/toml"

type Config struct {
	NameSystem string `toml:"nameSystem"`
	Version    string `toml:"version"`
	Server     struct {
		Enabled                bool     `toml:"enabled"`
		Port                   int      `toml:"port"`
		ReadTimeout            string   `toml:"readTimeout"`
		WriteTimeout           string   `toml:"writeTimeout"`
		IdleTimeout            string   `toml:"idleTimeout"`
		MaxHeaderBytes         int      `toml:"maxHeaderBytes"`
		MaxBodyBytes           int      `toml:"maxBodyBytes"`
		Cors                   bool     `toml:"cors"`
		CorsAllowOrigins       []string `toml:"corsAllowOrigins"`
		CorsAllowMethods       []string `toml:"corsAllowMethods"`
		CorsAllowHeaders       []string `toml:"corsAllowHeaders"`
		CorsAllowCredentials   bool     `toml:"corsAllowCredentials"`
		CorsMaxAge             string   `toml:"corsMaxAge"`
		CorsOptionsPassthrough bool     `toml:"corsOptionsPassthrough"`
		LogLevel               string   `toml:"logLevel"`
		LogFile                string   `toml:"logFile"`
	} `toml:"server"`
}

// ReadFile reads the configuration file
func ReadFile() (*Config, error) {
	var cfg Config
	if _, err := toml.DecodeFile("config.toml", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
