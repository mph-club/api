package server

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// CreateAndListen exposes the listen and creation of the api
func CreateAndListen() {
	_api := echo.New()

	// API v1
	v1 := _api.Group("/api/v1")

	// Checking if not in production env
	// enables Swagger UI
	if appEnv := os.Getenv("APP_ENV"); appEnv != "production" {
		v1.Use(middleware.StaticWithConfig(middleware.StaticConfig{
			Root:  "swagger",
			Index: "index.html",
		}))

		// Support CORS
		_api.Use(cors())
	}

	// Connect logger
	v1.Use(connectLogger())

	registerRoutes(v1)

	_api.Start(":8080")
}
