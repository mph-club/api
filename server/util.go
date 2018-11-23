package server

import (
	"image/jpeg"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/nfnt/resize"
)

func thumbnailPhoto(src multipart.File) (*os.File, error) {
	img, err := jpeg.Decode(src)
	if err != nil {
		return nil, err
	}

	thumb := resize.Thumbnail(200, 0, img, resize.NearestNeighbor)

	out, err := os.Create("thumbs.jpg")
	if err != nil {
		return nil, err
	}
	defer out.Close()

	jpeg.Encode(out, thumb, nil)

	return out, nil
}

func batchUpload(src multipart.File, thumbnail *os.File, vehicleID, filename string) error {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	})
	if err != nil {
		return err
	}

	uploader := s3manager.NewUploader(sess)

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
