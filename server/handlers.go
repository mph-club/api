package server

import (
	"fmt"
	"mphclub-rest-server/database"
	"mphclub-rest-server/models"
	"net/http"

	"github.com/labstack/echo"
)

// POST HANDLERS
func upsertListing(ctx echo.Context) error {
	var v models.Vehicle

	v.UserID = ctx.Get("sub").(string)

	if err := ctx.Bind(&v); err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"json_bind_error": err.Error()}))
	}

	carID, result, err := database.UpsertListing(v)
	if err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"database_error": err.Error()}))
	}

	resultString := fmt.Sprintf("vehicle was successfully %s", result)

	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			map[string]interface{}{
				"result": resultString,
				"id":     carID,
			},
		))
}

func updateUser(ctx echo.Context) error {
	var u models.User

	if err := ctx.Bind(&u); err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"json_bind_error": err.Error()}))
	}

	if err := database.UpsertUser(u); err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"database_error": err.Error()}))
	}

	resultString := fmt.Sprintf("user was successfully updated")

	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			map[string]interface{}{"result": resultString},
		))
}

func uploadCarPhoto(ctx echo.Context) error {
	vehicleID := ctx.FormValue("vehicle")

	file, err := ctx.FormFile("photo")
	if err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"form_file_error": err.Error()}))
	}

	if err := batchUploadCarAndThumbPhoto(file, vehicleID, file.Filename); err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"batch_upload_error": err.Error()}))
	}

	if err := database.EditPhotoURLArrayOnVehicle(vehicleID, file.Filename); err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"database_error": err.Error()}))
	}

	resultString := fmt.Sprintf("photo was successfully uploaded to the bucket and attached to %s", vehicleID)

	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			map[string]interface{}{"result": resultString},
		))
}

func uploadUserPhoto(ctx echo.Context) error {
	userID := ctx.Get("sub").(string)

	file, err := ctx.FormFile("photo")
	if err != nil {
		return err
	}

	location, err := uploadUserPhotoToS3(file, userID, file.Filename)
	if err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"upload_to_s3_error": err.Error()}))
	}

	err = database.AddUserPhotoURL(userID, location)
	if err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"database_error": err.Error()}))
	}

	resultString := fmt.Sprintf("photo was successfully uploaded to the bucket and attached to %s", userID)

	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			map[string]interface{}{"result": resultString},
		))
}

func uploadDriverLicense(ctx echo.Context) error {
	userID := ctx.Get("sub").(string)
	var dl models.DriverLicense

	if err := ctx.Bind(&dl); err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"json_bind_error": err.Error()}))
	}

	if err := database.AddDriverLicense(userID, &dl); err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"databse_error": err.Error()}))
	}

	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			map[string]interface{}{"result": "drivers license added"},
		))
}

// GET HANDLERS
func getDriverLicense(ctx echo.Context) error {
	userID := ctx.Get("sub").(string)

	license, err := database.GetDriverLicense(userID)
	if err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"database_error": err.Error()}))
	}

	if license.ID == 0 {
		return ctx.JSON(
			response(
				true,
				http.StatusOK,
				map[string]interface{}{"license": "empty license"},
			))
	}

	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			map[string]interface{}{"license": license},
		))
}

func getMyCars(ctx echo.Context) error {
	var u models.User
	u.ID = ctx.Get("sub").(string)

	list, err := database.GetMyCars(&u)
	if err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"database_error": err.Error()}))
	}

	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			map[string]interface{}{"vehicles": list}))
}

func getCars(ctx echo.Context) error {
	urlQuery := ctx.Request().URL.Query()

	count, list, err := database.GetCars(urlQuery)
	if err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"database_error": err.Error()}))
	}

	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			map[string]interface{}{"vehicles": list, "count": count},
		))
}

func getCarDetail(ctx echo.Context) error {
	var v models.Vehicle
	v.ID = ctx.Param("id")

	detail, err := database.GetCarDetail(v)
	if err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"db_error": err.Error()}))
	}

	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			map[string]interface{}{
				"Vehicle": detail,
			},
		))
}

func getUser(ctx echo.Context) error {
	userID := ctx.Get("sub").(string)

	user, err := database.GetUser(userID)
	if err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"database_error": err.Error()}))
	}

	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			map[string]interface{}{"user": user},
		))
}

func exploreCars(ctx echo.Context) error {
	list, err := database.GetExplore()

	if err != nil {
		return ctx.JSON(response(false, http.StatusBadRequest, map[string]interface{}{"database_error": err.Error()}))
	}

	return ctx.JSON(
		response(
			true,
			http.StatusOK,
			list,
		))
}
