package controllers

import "github.com/faruqii/Backend-Mobile-/internal/services"

type UserController struct {
	userService services.UserServices
}

func NewUserController(userService services.UserServices) *UserController {
	return &UserController{userService}
}
