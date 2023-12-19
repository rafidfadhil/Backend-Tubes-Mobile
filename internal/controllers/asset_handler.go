package controllers

import "github.com/faruqii/Backend-Mobile-/internal/services"

type AssetController struct {
	assetService services.AssetsServices
}

func NewAssetController(assetService services.AssetsServices) *AssetController {
	return &AssetController{assetService}
}