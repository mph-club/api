package server

import (
	"log"
	"mphclub-rest-server/database"
	"mphclub-rest-server/models"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
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
	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession())

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	log.Println(ctx.FormFile("photo"))

	file, info, err := ctx.FormFile("uploadfile")

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(generateJSONResponse(false, iris.Map{"server_error": err.Error()}))
		return
	}

	defer file.Close()
	filename := info.Filename
	//
	// f, err := os.Open(filename)
	// if err != nil {
	// 	ctx.StatusCode(iris.StatusBadRequest)
	// 	ctx.JSON(generateJSONResponse(false, iris.Map{"os_error": err.Error()}))
	// }

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(generateJSONResponse(false, iris.Map{"aws_error": err.Error()}))
	}

	log.Printf("file uploaded to, %s\n", result.Location)
}
