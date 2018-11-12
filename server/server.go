package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// CreateAndListen exposes the listen and creation of the api
func CreateAndListen() {
	_api := echo.New()

	v1 := _api.Group("api/v1")
	v1.Use(middleware.Logger())

	//  **** PUBLIC ****

	v1.GET("/getCars", getCars)

	v1.GET("/home", func(ctx echo.Context) error {
		return ctx.String(200, "api home!!!!")
	})

	v1.GET("/service", func(ctx echo.Context) error {
		return ctx.String(200, "api service!!!!")
	})

	// fs := http.FileServer(http.Dir("./swagger"))
	// v1.GET("/swagger/*", echo.WrapHandler(http.StripPrefix("/swagger/", fs)))

	v1.Use(middleware.Static("./swagger"))
	//  **** PRIVATE ****

	// GET
	v1.GET("/getMyCars", getMyCars, cognitoAuth)

	// POST
	v1.POST("/updateUser", updateUser, cognitoAuth)
	v1.POST("/listCar", upsertListing, cognitoAuth)
	v1.POST("/uploadPhoto", uploadToS3, cognitoAuth)

	_api.Start(":8080")
}
