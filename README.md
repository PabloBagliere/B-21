# PROYECTO B-21

## ¿Que es el proyecto B-21?

En la actualidad exactamente no se. Es un proyecto echo en golang que estoy desarrollando para aprender el lenguaje y sus funcionalidades.

## ¿Que hace el proyecto B-21?

El proyecto B-21 en mi mente se trata de un sistema de microservicios en golang realizando todo desde el mismo repositorio. El mismo tendria entres sus servicios:

- Un API Gateway: Que se encargara de redirigir las peticiones a los servicios correspondientes.
- Un servicio de autenticacion: Que se encargara de manejar los usuarios el logeo con OAuth2 y openid connect. El mismo tambien maneja roles y permisos.
- Un servicio de base de dato: Que se encargara de manejar, conectar y ejecutar las query en la base de dato de los demas servicios.

Todos estos servicios tendran que exponer la documentacion de sus endpoints en un swagger-ui y las metricas de sus servicios en un prometheus.
Todos tendran una REST API y una gRPC API.
Todos tendran un sistema de logs y un sistema de monitoreo.
Todos tendran un endpoint de salud (healthcheck).

## ¿Como se ejecuta el proyecto B-21?

Para ejecutar el proyecto B-21 tienes que ejecutar el comando `go run` nombre de la carpeta del servicio que quieras ejecutar. Por ejemplo si quieres ejecutar el servicio de autenticacion tienes que ejecutar el comando `go run cmd/auth`.

## ¿Como se prueba el proyecto B-21?

Para probar el proyecto B-21 tienes que ejecutar el comando `go test` nombre de la carpeta del servicio que quieras probar. Por ejemplo si quieres probar el servicio de autenticacion tienes que ejecutar el comando `go test cmd/auth`.

## ¿Como se documenta el proyecto B-21 en swagger?

Para documentar el proyecto B-21 tienes que ejecutar el comando `swag init` nombre de la carpeta del servicio que quieras documentar. Por ejemplo si quieres documentar el servicio de autenticacion tienes que ejecutar el comando `swag init -g cmd\auth\main.go --output api/auth`. -g es la ruta del archivo principal del servicio y --output es la ruta donde se guardara la documentacion.
