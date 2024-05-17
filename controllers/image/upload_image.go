package image

import (
	"context"
	"net/http"

	"github.com/JesseNicholas00/HaloSuster/services/image"
	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
)

func (ctrl *imageController) uploadImage(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	params := &s3.PutObjectInput{
		Bucket: aws.String(ctrl.bucket),
		Key:    aws.String(file.Filename),
		Body:   src,
	}

	result, err := ctrl.service.Upload(context.TODO(), params)
	if err != nil {
		return errorutil.AddCurrentContext(err)
	}

	res := image.UploadImageRes{
		ImageUrl: result.Location,
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Image uploaded successfully",
		"data":    res,
	})
}
