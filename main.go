package main

import "github.com/kataras/iris"

func api() *iris.Application {
	app := iris.New()

	app.Get("/mphclub", func(ctx iris.Context) {
		ctx.ServeFile("./swagger-ui/index.html", false)
	})

	return app
}

func main() {
	_api := api()

	v1 := _api.Party("api/v1")
	{
		v1.Get("/", func(ctx iris.Context) {
			ctx.Writef("api home!!!!")
		})
	}

	_api.Run(iris.Addr(":8080"))
}
