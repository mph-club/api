package server

import (
	"log"
	"mphclub-rest-server/database"
	"mphclub-rest-server/models"
	"strings"

	"github.com/kataras/iris"
)

func createListing(ctx iris.Context) {
	var v models.Vehicle

	if err := ctx.ReadJSON(&v); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(generateJSONResponse(false, iris.Map{"error": err.Error()}))
		return
	}

	if err := database.CreateListing(v); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(generateJSONResponse(false, iris.Map{"database_error": err.Error()}))
		return
	}

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

	if err := database.CreateUser(u); err != nil {
		pkExists := "ERROR #23505"
		var errorString string

		if strings.Contains(err.Error(), pkExists) {
			errorString = "user sub already exists in database"
		} else {
			errorString = err.Error()
		}

		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(generateJSONResponse(false, iris.Map{"database_error": errorString}))
		return
	}

	log.Println(u)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(generateJSONResponse(true, iris.Map{"result": "user was successfully created"}))
}

func getCars(ctx iris.Context) {
	list, err := database.GetCars()

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(generateJSONResponse(false, iris.Map{"database_error": err.Error()}))
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(generateJSONResponse(true, iris.Map{"vehicles": list}))
}
