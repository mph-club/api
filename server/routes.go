package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func registerRoutes(g *echo.Group) {
	//  **** PUBLIC ****
	g.GET("/home", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "api home!!!!")
	})
	g.GET("/service", func(ctx echo.Context) error {
		return ctx.String(200, "api service!!!!")
	})

	g.GET("/vehicles", getCars)
	g.GET("/vehicles/:id", getCarDetail)
	g.GET("/users/:id", getHostDetail)
	g.GET("/explore", exploreCars)

	//  **** PRIVATE ****
	// GET
	g.GET("/getMyCars", getMyCars, cognitoAuth)
	g.GET("/driverLicense", getDriverLicense, cognitoAuth)
	g.GET("/account", getUser, cognitoAuth)
	g.GET("/reserve", getMyReservations, cognitoAuth)

	// POST
	g.POST("/updateUser", updateUser, cognitoAuth)
	g.POST("/listCar", upsertListing, cognitoAuth)
	g.POST("/uploadCarPhoto", uploadCarPhoto, cognitoAuth)
	g.POST("/uploadUserPhoto", uploadUserPhoto, cognitoAuth)
	g.POST("/driverLicense", uploadDriverLicense, cognitoAuth)
	g.POST("/reserve", makeReservation, cognitoAuth)
	g.POST("/insurance", addInsurance, cognitoAuth)
	g.POST("/cardInfo", addCardInfo, cognitoAuth)
}
