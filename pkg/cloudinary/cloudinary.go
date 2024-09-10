package cloudinary

import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
)

func NewCloudinary(apiKey, apiSecret, cloudName string) (*cloudinary.Cloudinary, context.Context) {

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)

	if err != nil {
		fmt.Println("cloudinary setup failed", err)
	}

	return cld, context.Background()
}

func GetAssetInfo(cld *cloudinary.Cloudinary, c context.Context, publicID string) (*admin.AssetResult, error) {
	res, err := cld.Admin.Asset(c, admin.AssetParams{PublicID: publicID})

	if err != nil {
		return nil, fmt.Errorf("error getting assets")
	}

	return res, nil
}
