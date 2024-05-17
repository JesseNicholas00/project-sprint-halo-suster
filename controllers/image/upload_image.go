package image

import (
	"context"
	"net/http"
	"strings"

	"github.com/JesseNicholas00/HaloSuster/services/image"
	"github.com/JesseNicholas00/HaloSuster/utils/errorutil"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const maxSize int64 = 20 << 20
const minSize int64 = 10 << 10

func (ctrl *imageController) uploadImage(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "image is wrong",
		})
	}
	fileType := strings.Split(file.Header.Values("Content-Type")[0], "/")[1]
	if fileType != "jpg" && fileType != "jpeg" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "image is wrong",
		})
	}
	if file.Size < minSize || file.Size > maxSize {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "image is wrong",
		})
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	params := &s3.PutObjectInput{
		Bucket: aws.String(ctrl.bucket),
		Key:    aws.String(uuid.NewString() + "." + fileType),
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
		"message": "File uploaded successfully",
		"data":    res,
	})
}
