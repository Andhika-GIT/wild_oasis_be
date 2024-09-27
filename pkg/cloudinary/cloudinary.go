package cloudinary

import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
)

func NewCloudinary(apiKey, apiSecret, cloudName string) (*cloudinary.Cloudinary, context.Context, error) {

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)

	if err != nil {
		return &cloudinary.Cloudinary{}, context.Background(), fmt.Errorf("failed to setup cloudinary")
	}

	return cld, context.Background(), nil
}

func GetAssetInfo(cld *cloudinary.Cloudinary, c context.Context, publicID string) (*admin.AssetResult, error) {
	res, err := cld.Admin.Asset(c, admin.AssetParams{PublicID: publicID})

	if err != nil {
		return nil, fmt.Errorf("error getting assets")
	}

	return res, nil
}
