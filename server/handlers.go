package server

import (
	"fmt"
	"mphclub-rest-server/database"
	"mphclub-rest-server/models"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo"
)

func upsertListing(ctx echo.Context) error {
	var v models.Vehicle

	v.UserID = ctx.Get("sub").(string)

	if err := ctx.Bind(&v); err != nil {
		return ctx.JSON(generateJSONResponse(false, http.StatusBadRequest, map[string]interface{}{"error": err.Error()}))
	}

	carID, result, err := database.UpsertListing(v)
	if err != nil {
		return ctx.JSON(generateJSONResponse(false, http.StatusBadRequest, map[string]interface{}{"database_error": err.Error()}))
	}

	resultString := fmt.Sprintf("vehicle was successfully %s", result)

	return ctx.JSON(generateJSONResponse(true, http.StatusOK, map[string]interface{}{"result": resultString, "id": carID}))
}

func updateUser(ctx echo.Context) error {
	var u models.User
	u.ID = ctx.Get("sub").(string)

	if err := ctx.Bind(&u); err != nil {
		return ctx.JSON(generateJSONResponse(false, http.StatusBadRequest, map[string]interface{}{"error": err.Error()}))
	}

	if err := database.UpsertUser(u); err != nil {
		return ctx.JSON(generateJSONResponse(false, http.StatusBadRequest, map[string]interface{}{"error": err.Error()}))
	}

	resultString := fmt.Sprintf("user was successfully updated")
	return ctx.JSON(generateJSONResponse(true, http.StatusOK, map[string]interface{}{"result": resultString}))
}

func getMyCars(ctx echo.Context) error {
	var u models.User
	u.ID = ctx.Get("sub").(string)

	list, err := database.GetMyCars(&u)

	if err != nil {
		return ctx.JSON(generateJSONResponse(false, http.StatusBadRequest, map[string]interface{}{"database_error": err.Error()}))
	}

	return ctx.JSON(generateJSONResponse(true, http.StatusOK, map[string]interface{}{"vehicles": list}))
}

func getCars(ctx echo.Context) error {
	list, err := database.GetCars()

	if err != nil {
		return ctx.JSON(generateJSONResponse(false, http.StatusBadRequest, map[string]interface{}{"database_error": err.Error()}))
	}

	return ctx.JSON(generateJSONResponse(true, http.StatusOK, map[string]interface{}{"vehicles": list}))
}

func exploreCars(ctx echo.Context) error {
	queryParams := ctx.Request().URL.Query()
	carType := ctx.QueryParam("type")

	list, err := database.GetExplore(carType, queryParams)

	if err != nil {
		return ctx.JSON(generateJSONResponse(false, http.StatusBadRequest, map[string]interface{}{"database_error": err.Error()}))
	}

	return ctx.JSON(generateJSONResponse(true, http.StatusOK, map[string]interface{}{"explore": list}))
}

func uploadToS3(ctx echo.Context) error {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	})
	if err != nil {
		return ctx.JSON(generateJSONResponse(false, http.StatusServiceUnavailable, map[string]interface{}{"aws_auth_err": err.Error()}))
	}

	uploader := s3manager.NewUploader(sess)

	vehicleID := ctx.FormValue("vehicle")

	file, err := ctx.FormFile("photo")
	if err != nil {
		return ctx.JSON(generateJSONResponse(false, http.StatusInternalServerError, map[string]interface{}{"form_file_error": err.Error()}))
	}

	src, err := file.Open()
	if err != nil {
		return ctx.JSON(generateJSONResponse(false, http.StatusInternalServerError, map[string]interface{}{"form_file_open_error": err.Error()}))
	}
	defer src.Close()

	filename := file.Filename

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("AWS_BUCKET")),
		Key:         aws.String(vehicleID + "/" + filename),
		Body:        src,
		ContentType: aws.String("image/jpeg"),
		ACL:         aws.String("public-read"),
	})

	if err != nil {
		return ctx.JSON(generateJSONResponse(false, http.StatusBadRequest, map[string]interface{}{"aws_error": err.Error()}))
	}

	err = database.EditPhotoURLArrayOnVehicle(vehicleID, result.Location)
	if err != nil {
		return ctx.JSON(generateJSONResponse(false, http.StatusBadRequest, map[string]interface{}{"db_error": err.Error()}))
	}

	resultString := fmt.Sprintf("photo was successfully uploaded to the bucket and attached to %s", vehicleID)

	return ctx.JSON(generateJSONResponse(true, http.StatusOK, map[string]interface{}{"result": resultString}))
}
