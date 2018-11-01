package server

import (
	"fmt"
	"mphclub-rest-server/database"
	"mphclub-rest-server/models"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/kataras/iris"
)

func upsertListing(ctx iris.Context) {
	var v models.Vehicle

	v.User = ctx.Values().Get("sub").(string)

	if err := ctx.ReadJSON(&v); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(generateJSONResponse(false, iris.Map{"error": err.Error()}))
		return
	}

	carID, result, err := database.UpsertListing(v)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(generateJSONResponse(false, iris.Map{"database_error": err.Error()}))
		return
	}

	resultString := fmt.Sprintf("vehicle was successfully %s", result)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(generateJSONResponse(true, iris.Map{"result": resultString, "id": carID}))
}

func updateUser(ctx iris.Context) {
	var u models.User
	u.Sub = ctx.Values().Get("sub").(string)

	if err := ctx.ReadJSON(&u); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(generateJSONResponse(false, iris.Map{"error": err.Error()}))
		return
	}

	if err := database.UpsertUser(u); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(generateJSONResponse(false, iris.Map{"error": err.Error()}))
		return
	}

	resultString := fmt.Sprintf("user was successfully updated")
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(generateJSONResponse(true, iris.Map{"result": resultString}))
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

func uploadToS3(ctx iris.Context) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	})
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(generateJSONResponse(false, iris.Map{"aws_auth_err": err.Error()}))
		return
	}

	uploader := s3manager.NewUploader(sess)

	vehicleID := ctx.FormValue("vehicle")

	file, info, err := ctx.FormFile("photo")

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(generateJSONResponse(false, iris.Map{"server_error": err.Error()}))
		return
	}

	defer file.Close()
	filename := info.Filename

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("AWS_BUCKET")),
		Key:         aws.String(vehicleID + "/" + filename),
		Body:        file,
		ContentType: aws.String("image/jpeg"),
		ACL:         aws.String("public-read"),
	})

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(generateJSONResponse(false, iris.Map{"aws_error": err.Error()}))
		return
	}

	err = database.EditPhotoURLArrayOnVehicle(vehicleID, result.Location)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(generateJSONResponse(false, iris.Map{"db_error": err.Error()}))
		return
	}

	resultString := fmt.Sprintf("photo was successfully uploaded to the bucket and attached to %s", vehicleID)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(generateJSONResponse(true, iris.Map{"result": resultString}))
}
