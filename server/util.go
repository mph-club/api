package server

import (
	"image"
	"image/jpeg"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/nfnt/resize"
)

func thumbnailPhoto(src multipart.File) (image.Image, error) {
	if _, err := src.Seek(0, 0); err != nil {
		return nil, err
	}

	img, _, err := image.Decode(src)
	if err != nil {
		return nil, err
	}

	t := resize.Thumbnail(400, 400, img, resize.NearestNeighbor)

	return t, nil
}

func batchUpload(file *multipart.FileHeader, vehicleID, filename string) error {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	})
	if err != nil {
		return err
	}

	uploader := s3manager.NewUploader(sess)

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	thumb, err := thumbnailPhoto(src)
	if err != nil {
		return err
	}

	thumbnail, err := os.Create("thumbs.jpg")
	if err != nil {
		return err
	}
	defer thumbnail.Close()

	jpeg.Encode(thumbnail, thumb, nil)

	if _, err := src.Seek(0, 0); err != nil {
		return err
	}

	if _, err := thumbnail.Seek(0, 0); err != nil {
		return err
	}

	objects := []s3manager.BatchUploadObject{
		{
			Object: &s3manager.UploadInput{
				Bucket:      aws.String(os.Getenv("AWS_BUCKET")),
				Key:         aws.String(vehicleID + "/" + filename),
				Body:        src,
				ContentType: aws.String("image/jpeg"),
				ACL:         aws.String("public-read"),
			},
		},
		{
			Object: &s3manager.UploadInput{
				Bucket:      aws.String(os.Getenv("AWS_BUCKET")),
				Key:         aws.String(vehicleID + "/thumb/" + filename),
				Body:        thumbnail,
				ContentType: aws.String("image/jpeg"),
				ACL:         aws.String("public-read"),
			},
		},
	}

	iter := &s3manager.UploadObjectsIterator{Objects: objects}

	if err := uploader.UploadWithIterator(aws.BackgroundContext(), iter); err != nil {
		return err
	}

	return nil
}
