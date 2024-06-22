package main

import (
	"fmt"

	"github.com/PabloBagliere/B-21/internal/auth"
	"github.com/PabloBagliere/B-21/pkg/config"
	"github.com/PabloBagliere/B-21/pkg/logger"
	"github.com/PabloBagliere/B-21/pkg/server"
	"github.com/rs/zerolog"
)

// @title Swagger API Auth B-21
// @version 0.0.1
// @description This is the API for the Auth service of the B-21 project.

// @contact.name API Auth B-21
// @contact.url https://github.com/PabloBagliere/B-21/issues
// @contact.email pablo.bagliere2@gmail.com

// @license.name MIT
// @license.url https://github.com/PabloBagliere/B-21/licence

// @host localhost:8080
func main() {
	serverName := "auth"
	p, err := config.GetConfig(serverName)
	if err != nil {
		panic(err)
	}
	// check si enabled está en la configuración y es true
	if _, ok := p["enabled"]; !ok || !p["enabled"].(bool) {
		panic(fmt.Sprintf("Error: %s not enabled in the configuration", serverName))
	}
	// Inicializar el logger
	logger.NewLogger(serverName, zerolog.InfoLevel)
	// Inicializar el JWT
	_, err = auth.InitJwt(p)
	if err != nil {
		panic(err)
	}
	jwt, err := auth.CreateResponse()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", jwt)
	isValid, err := auth.ValidateToken(jwt.AccessToken)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Token is valid: %v\n", isValid)
	// Check si el puerto está en la configuración
	if _, ok := p["port"]; !ok {
		panic("Error: port not found in the configuration")
	}

	// Extraer el port de la configuración y pasarlo al Start de echo
	port := fmt.Sprintf(":%v", p["port"])

	e := server.NewServer(serverName)
	e.Logger.Fatal(e.Start(port))
}
