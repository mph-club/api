package server

import (
	"log"
	"mphclub-rest-server/database"
	"mphclub-rest-server/models"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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

	file, info, err := ctx.FormFile("photo")

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(generateJSONResponse(false, iris.Map{"server_error": err.Error()}))
		return
	}

	defer file.Close()
	filename := info.Filename

	log.Println(filename)

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(generateJSONResponse(false, iris.Map{"aws_error": err.Error()}))
		return
	}

	log.Printf("file uploaded to, %s\n", result.Location)
}
