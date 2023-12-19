package controllers

import "github.com/rafidfadhil/Backend-Tubes-Mobile/internal/services"

type UserController struct {
	userService services.UserServices
}

func NewUserController(userService services.UserServices) *UserController {
	return &UserController{userService}
}
