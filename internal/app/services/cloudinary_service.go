package services

import (
	"context"

	"github.com/Andhika-GIT/wild_oasis_be/pkg/cloudinary"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/spf13/viper"
)

type CloudinaryService struct {
	viper *viper.Viper
}

func NewCloudinaryService(v *viper.Viper) *CloudinaryService {
	return &CloudinaryService{
		viper: v,
	}
}

func (cs *CloudinaryService) CheckImageAssets(c context.Context, publicID string) (*admin.AssetResult, error) {
	apiKey := cs.viper.GetString("CLOUDINARY_API_KEY")
	apiSecret := cs.viper.GetString("CLOUDINARY_API_SECRET")
	cloudName := cs.viper.GetString("CLOUDINARY_CLOUD_NAME")

	cloud, ctx, err := cloudinary.NewCloudinary(apiKey, apiSecret, cloudName)

	if err != nil {
		return nil, err
	}

	res, err := cloudinary.GetAssetInfo(cloud, ctx, publicID)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (cs *CloudinaryService) GetImagePublicUrl(c context.Context, publicID string) (string, error) {
	apiKey := cs.viper.GetString("CLOUDINARY_API_KEY")
	apiSecret := cs.viper.GetString("CLOUDINARY_API_SECRET")
	cloudName := cs.viper.GetString("CLOUDINARY_CLOUD_NAME")

	cloud, ctx, err := cloudinary.NewCloudinary(apiKey, apiSecret, cloudName)

	if err != nil {
		return "", err
	}

	res, err := cloudinary.GetAssetInfo(cloud, ctx, publicID)

	if err != nil {
		return "", err
	}

	return res.URL, nil

}
