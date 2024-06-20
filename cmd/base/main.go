package main

import (
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
	e := server.NewServer("base", zerolog.InfoLevel)
	e.Logger.Fatal(e.Start(":8080"))
}
