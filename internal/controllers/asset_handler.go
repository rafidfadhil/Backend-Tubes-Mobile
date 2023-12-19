package controllers

import "github.com/rafidfadhil/Backend-Tubes-Mobile/internal/services"

type AssetController struct {
	assetService services.AssetsServices
}

func NewAssetController(assetService services.AssetsServices) *AssetController {
	return &AssetController{assetService}
}