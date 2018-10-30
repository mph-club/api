package server

import (
	"log"

	"mphclub-rest-server/models"

	"github.com/kataras/iris"
)

func postListing(ctx iris.Context) {
	log.Println(ctx.FormValue("userData"))

	ctx.JSON(makeResponse(true, iris.Map{"userData": ctx.FormValue("userData")}))
}

func createUser(ctx iris.Context) {
	var u models.UserInfo

	if err := ctx.ReadJSON(&u); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(makeResponse(false, iris.Map{"error": err.Error()}))
		return
	}

	ctx.Writef("Received: %#+v\n", u)

}
