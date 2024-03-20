package imgupload

import (
	"log"

	"github.com/abdullahnettoor/connect-social-media/internal/config"
	"github.com/cloudinary/cloudinary-go"
)

func ConnectCloudinary(cfg *config.Config) (*cloudinary.Cloudinary, error) {
	cld, err := cloudinary.NewFromURL(cfg.ContentCloudUri)
	if err != nil {
		log.Println("Failed to initialize Cloudinary", err)
		return nil, err
	}
	log.Println("cloudinary connection established")
	return cld, err
}
