package server

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// CreateAndListen exposes the listen and creation of the api
func CreateAndListen() {
	_api := echo.New()

	// Support CORS
	_api.Use(cors())

	// API v1
	v1 := _api.Group("/api/v1")

	// Connect logger
	v1.Use(connectLogger())

	//  **** PUBLIC ****
	v1.GET("/vehicles", getCars)

	v1.GET("/home", func(ctx echo.Context) error {
		return ctx.String(200, "api home!!!!")
	})

	v1.GET("/service", func(ctx echo.Context) error {
		return ctx.String(200, "api service!!!!")
	})

	v1.GET("/explore", exploreCars)

	// Checking if not in production env
	// enables Swagger UI
	if appEnv := os.Getenv("APP_ENV"); appEnv != "production" {
		v1.Use(middleware.StaticWithConfig(middleware.StaticConfig{
			Root:  "swagger",
			Index: "index.html",
		}))
	}

	//  **** PRIVATE ****
	// GET
	v1.GET("/getMyCars", getMyCars, cognitoAuth)
	v1.GET("/driverLicense", getDriverLicense, cognitoAuth)
	v1.GET("/user", getUser, cognitoAuth)

	// POST
	v1.POST("/updateUser", updateUser, cognitoAuth)
	v1.POST("/listCar", upsertListing, cognitoAuth)
	v1.POST("/uploadCarPhoto", uploadCarPhoto, cognitoAuth)
	v1.POST("/uploadUserPhoto", uploadUserPhoto, cognitoAuth)
	v1.POST("/driverLicense", uploadDriverLicense, cognitoAuth)

	_api.Start(":8080")
}
