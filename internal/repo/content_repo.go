package repo

import (
	"context"
	"log"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

type ContentRepository struct {
	db *cloudinary.Cloudinary
}

func NewContentRepository(db *cloudinary.Cloudinary) *ContentRepository {
	return &ContentRepository{db}
}

func (r *ContentRepository) UploadImage(ctx context.Context, imageName, dir string, imageFile any) (string, error) {
	log.Println("File is", imageFile)
	result, err := r.db.Upload.Upload(
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
