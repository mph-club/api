package main

import (
	"github.com/kataras/iris"
)

func main() {
	_api := iris.New()

	_api.Get("/mphclub", func(ctx iris.Context) {
		ctx.ServeFile("./swagger/redoc-static.html", false)
	})

	v1 := _api.Party("api/v1")
	{
		v1.Get("/", func(ctx iris.Context) {
			ctx.Writef("api home!!!!")
		})

		v1.Get("/service", func(ctx iris.Context) {
			ctx.Writef("api service!!!!")
		})
	}

	_api.Run(iris.Addr(":8080"))
}
