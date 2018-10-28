package server

import "github.com/kataras/iris"

// CreateAndListen exposes the listen and creation of the api
func CreateAndListen() {
	_api := iris.New()

	v1 := _api.Party("api/v1")
	{
		v1.Get("/", func(ctx iris.Context) {
			ctx.Writef("api home!!!!")
		})

		v1.Get("/service", func(ctx iris.Context) {
			ctx.Writef("api service!!!!")
		})

		v1.Get("/swagger", func(ctx iris.Context) {
			ctx.ServeFile("./swagger/index.html", false)
		})

		v1.Post("/listCar", postListing)
	}

	_api.Run(iris.Addr(":8080"))
}
