package main

import (
	"fmt"

	"github.com/PabloBagliere/B-21/pkg/config"
	"github.com/PabloBagliere/B-21/pkg/logger"
	"github.com/PabloBagliere/B-21/pkg/server"
	"github.com/rs/zerolog"
)

// @title Swagger API Example
// @version 0.0.1
// @description

// @contact.name API Support
// @contact.url https://github.com/PabloBagliere/B-21/issues
// @contact.email pablo.bagliere2@gmail.com

// @license.name MIT
// @license.url https://github.com/PabloBagliere/B-21/licence

// @host localhost:8080
func main() {
	serverName := "server"
	logger.NewLogger(serverName, zerolog.InfoLevel)
	p, err := config.GetConfig(serverName)
	if err != nil {
		panic(err)
	}

	// check si enabled está en la configuración y es true
	if _, ok := p["enabled"]; !ok || !p["enabled"].(bool) {
		panic(fmt.Sprintf("Error: %s not enabled in the configuration", serverName))
	}

	// Check si el puerto está en la configuración
	if _, ok := p["port"]; !ok {
		panic("Error: port not found in the configuration")
	}

	// Extraer el port de la configuración y pasarlo al Start de echo
	port := fmt.Sprintf(":%v", p["port"])

	e := server.NewServer(serverName)
	e.Logger.Fatal(e.Start(port))
}
