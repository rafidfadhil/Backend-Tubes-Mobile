package routes

import (
	"github.com/rafidfadhil/Backend-Tubes-Mobile/internal/controllers"
	"github.com/rafidfadhil/Backend-Tubes-Mobile/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoute(router fiber.Router, userService services.UserServices) {
	userController := controllers.NewUserController(userService)

	router.Post("/register", userController.Register)
	router.Post("/login", userController.Login)
	
}