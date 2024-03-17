package imgupload

import (
	"context"
	"log"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/spf13/viper"
)

type UploadImage struct {
	cloud *cloudinary.Cloudinary
}

func NewUploadImage() (*UploadImage, error) {
	cld, err := cloudinary.NewFromURL(viper.GetString("IMG_CLOUD_URL"))
	if err != nil {
		log.Println("Failed to initialize Cloudinary", err)
		return nil, err
	}
	log.Println("cloudinary connection established")
	return &UploadImage{
		cloud: cld,
	}, err
}

func (h *UploadImage) Handler(ctx context.Context, imageName, dir string, imageFile any) (string, error) {
	log.Println("File is", imageFile)
	result, err := h.cloud.Upload.Upload(
		ctx,
		imageFile,
		uploader.UploadParams{
			PublicID: imageName,
			Folder:   "connectr/" + dir,
		})

	log.Println("Result is", result)
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Println("upload success")
	log.Println(result.SecureURL)
	return result.SecureURL, nil
}
