package image

import "context"

type ImageService interface {
	UploadImage(
		ctx context.Context,
		res *UploadImageRes,
	) error
}
