package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct {
	ServiceName string
	LogLevel    zerolog.Level
}

func InitLogger(config Config) {

	// Configura el formato de la fecha y hora
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// // Configura la salida del logger
	// log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Configura el nivel del logger
	zerolog.SetGlobalLevel(config.LogLevel)

	// Agrega un hook para incluir el nombre del servicio en cada log
	log.Logger = log.Logger.With().Str("service", config.ServiceName).Logger()
}

func NewLogger(serviceName string, logLevel zerolog.Level) {
	InitLogger(Config{
		ServiceName: serviceName,
		LogLevel:    logLevel,
	})
}
