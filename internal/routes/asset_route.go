package routes

import (
	"github.com/rafidfadhil/Backend-Tubes-Mobile/internal/controllers"
	"github.com/rafidfadhil/Backend-Tubes-Mobile/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupAssetRoute(router fiber.Router, assetService services.AssetsServices) {
	assetController := controllers.NewAssetController(assetService)

	assets := router.Group("/assets")
	assets.Post("/add", assetController.Add)
	assets.Post("/damage-report/:id", assetController.CreateDamageReport)
	assets.Get("/:id", assetController.Show)
	assets.Get("/", assetController.ShowAll)
	assets.Put("/update/:id", assetController.Update)
	assets.Delete("/delete/:id", assetController.Delete)
}
