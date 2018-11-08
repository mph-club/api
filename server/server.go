package server

import (
	"github.com/kataras/iris"
)

// CreateAndListen exposes the listen and creation of the api
func CreateAndListen() {
	_api := iris.New()

	v1 := _api.Party("api/v1")
	{
		v1.Use(requestLogger())

		//  **** PUBLIC ****

		v1.Get("/getCars", getCars)

		v1.Get("/home", func(ctx iris.Context) {
			ctx.Writef("api home!!!!")
		})

		v1.Get("/service", func(ctx iris.Context) {
			ctx.Writef("api service!!!!")
		})

		v1.Get("/swagger", func(ctx iris.Context) {
			ctx.ServeFile("./swagger/index.html", false)
		})
		//  **** PRIVATE ****

		// GET
		v1.Get("/getMyCars", cognitoAuth, getMyCars)

		// POST
		v1.Post("/updateUser", cognitoAuth, updateUser)
		v1.Post("/listCar", cognitoAuth, upsertListing)
		v1.Post("/uploadPhoto", cognitoAuth, uploadToS3)
	}

	_api.Run(iris.Addr(":8080"))
}
