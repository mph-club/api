package server

import (
	"log"

	"github.com/kataras/iris"
)

func postListing(ctx iris.Context) {
	log.Println(ctx.FormValue("userData"))

	ctx.JSON(makeResponse(true, iris.Map{"userData": ctx.FormValue("userData")}))
}
