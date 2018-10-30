package server

import (
	"log"

	"mphclub-rest-server/database"
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

	err := database.CreateUser(u)
	if err != nil {
		ctx.JSON(makeResponse(false, iris.Map{"database_error": err}))
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(makeResponse(true, iris.Map{"result": "user was successfully created"}))

}
