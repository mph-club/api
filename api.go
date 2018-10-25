package main

import (
	"github.com/kataras/iris"
)

func main() {
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
	}
	
	_api.Run(iris.Addr(":8080"))
}
