package server

import (
	"log"
	"mphclub-rest-server/database"
	"mphclub-rest-server/models"

	"github.com/kataras/iris"
)

func createListing(ctx iris.Context) {
	var v models.Vehicle

	if err := ctx.ReadJSON(&v); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(generateJSONResponse(false, iris.Map{"error": err.Error()}))
		return
	}
	err := database.CreateListing(v)
	if err != nil {
		ctx.JSON(generateJSONResponse(false, iris.Map{"database_error": err}))
	}

	log.Println(v)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(generateJSONResponse(true, iris.Map{"result": "vehicle was successfully inserted"}))
}

func createUser(ctx iris.Context) {
	var u models.UserInfo

	if err := ctx.ReadJSON(&u); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(generateJSONResponse(false, iris.Map{"error": err.Error()}))
		return
	}

	err := database.CreateUser(u)
	if err != nil {
		ctx.JSON(generateJSONResponse(false, iris.Map{"database_error": err}))
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(generateJSONResponse(true, iris.Map{"result": "user was successfully created"}))
}
