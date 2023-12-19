package app

import (
	"os"

	"github.com/faruqii/Backend-Mobile-/internal/config"
	"github.com/faruqii/Backend-Mobile-/internal/repositories"
	"github.com/faruqii/Backend-Mobile-/internal/routes"
	"github.com/faruqii/Backend-Mobile-/internal/services"
	"github.com/gofiber/fiber/v2"
)

func StartApp() {
	app := fiber.New()

	// init database
	db, err := config.Connect()

	if err != nil {
		panic(err)
	}

	// init  user repository
	userRepositories := repositories.NewUserRepository(db)

	// init user  service
	userServices := services.NewUserServices(userRepositories)

	// init asset repository
	assetRepositories := repositories.NewAssetRepository(db)

	// init asset service
	assetServices := services.NewAssetsServices(assetRepositories)

	// init route
	apiEndpoint := app.Group("/api")
	routes.SetupUserRoute(apiEndpoint, userServices)
	routes.SetupAssetRoute(apiEndpoint, assetServices)

	err = app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		panic(err)
	}
}
