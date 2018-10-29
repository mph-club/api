package server

import (
	"log"

	"github.com/kataras/iris"
)

func cognitoAuth(ctx iris.Context) {
	log.Println(ctx.Request().Header.Get("Authorization"))

	isAuth, err := checkToken(ctx.Request().Header.Get("Authorization"))

	if isAuth {
		ctx.Next()
	} else {
		ctx.JSON(makeResponse(false, iris.Map{"error": err}))
	}
}
