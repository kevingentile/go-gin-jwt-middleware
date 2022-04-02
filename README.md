# GO Gin JWT Middleware

[![GoDoc](https://pkg.go.dev/badge/github.com/kevingentile/go-gin-jwt-middleware.svg)](https://pkg.go.dev/github.com/kevingentile/go-gin-jwt-middleware/v2)
[![License](https://img.shields.io/github/license/kevingentile/go-gin-jwt-middleware.svg?style=flat-square)](https://github.com/kevingentile/go-gin-jwt-middleware/blob/master/LICENSE)
[![Release](https://img.shields.io/github/v/release/kevingentile/go-gin-jwt-middleware?include_prereleases&style=flat-square)](https://github.com/kevingentile/go-gin-jwt-middleware/releases)
[![Codecov](https://img.shields.io/codecov/c/github/kevingentile/go-gin-jwt-middleware?style=flat-square&token=fs2WrOXe9H)](https://codecov.io/gh/kevingentile/go-gin-jwt-middleware)
[![Tests](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Fauth0%2Fgo-jwt-middleware%2Fbadge%3Fref%3Dmaster&style=flat-square)](https://github.com/kevingentile/go-gin-jwt-middleware/actions?query=branch%3Amaster)
[![Stars](https://img.shields.io/github/stars/kevingentile/go-gin-jwt-middleware.svg?style=flat-square)](https://github.com/kevingentile/go-gin-jwt-middleware)
[![Contributors](https://img.shields.io/github/contributors/auth0/go-jwt-middleware?style=flat-square)](https://github.com/auth0/go-jwt-middleware/graphs/contributors)

---

Gin extension for the [auth0/go-jwt-middleware/v2](https://github.com/auth0/go-jwt-middleware) Golang middleware to check and validate [JWTs](jwt.io) in the request and add the valid token contents to the request 
context.

This is not an Auth0 project. For the latest features and fixes the offical middlware can be found at [auth0/go-jwt-middleware/v2](https://github.com/auth0/go-jwt-middleware)

-------------------------------------

## Installation

```shell
go get github.com/kevingentile/go-gin-jwt-middleware/v2
```


## Usage

```go
package main

import (
	"context"
	"log"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	jwtginmiddleware "github.com/kevingentile/go-gin-jwt-middleware/v2"
)

func main() {
	keyFunc := func(ctx context.Context) (interface{}, error) {
		// Our token must be signed using this data.
		return []byte("secret"), nil
	}

	// Set up the validator.
	jwtValidator, err := validator.New(
		keyFunc,
		validator.HS256,
		"https://<issuer-url>/",
		[]string{"<audience>"},
	)
	if err != nil {
		log.Fatalf("failed to set up the validator: %v", err)
	}
	auth0Middleware := jwtginmiddleware.New(
		jwtValidator.ValidateToken,
	)

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	authGroup := router.Group("auth")
	authGroup.Use(auth0Middleware.CheckJWT())

	//auth/claims
	authGroup.GET("/claims", func(c *gin.Context) {
		// Write claims JSON blob
		claims := c.Request.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		c.JSON(http.StatusOK, claims)
	})

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()

}
```

## License

This project is licensed under the MIT license. See the [LICENSE](LICENSE) file for more info.
