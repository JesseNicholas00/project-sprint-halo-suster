package image

import (
	"net/http"

	"github.com/JesseNicholas00/HaloSuster/services/image"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo/v4"
)

func (ctrl *imageController) uploadImage(c echo.Context) error {
	var res image.UploadImageRes

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("your-region"), // Change to your AWS region
	}))

	// Create an S3 service client
	svc := s3.New(sess)

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Prepare the parameters for uploading to S3
	params := &s3.PutObjectInput{
		Bucket: aws.String("your-bucket-name"), // Change to your bucket name
		Key:    aws.String(file.Filename),
		Body:   src,
	}

	// Upload the file to S3
	_, err = svc.PutObject(params)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "User logged in successfully",
		"data":    res,
	})
}
