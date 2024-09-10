package cloudinary

import (
	"context"
	"fmt"

	"github.com/Andhika-GIT/wild_oasis_be/internal/infrastructure/config"
	"github.com/cloudinary/cloudinary-go/v2"
)

func NewCloudinary() (*cloudinary.Cloudinary, context.Context) {
	v := config.NewViper()

	apiKey := v.GetString("CLOUDINARY_API_KEY")
	apiSecret := v.GetString("CLOUDINARY_API_SECRET")
	cloudName := v.GetString("CLOUDINARY_CLOUD_NAME")

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)

	if err != nil {
		fmt.Println("cloudinary setup failed", err)
	}

	return cld, context.Background()
}
