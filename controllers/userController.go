package controllers

import (
	"github.com/karim-w/gin-orm/services"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type UserController struct {
	logger  *zap.SugaredLogger
	service *services.UserService
}

func NewUserController(logger *zap.SugaredLogger, service *services.UserService) *UserController {
	return &UserController{
		logger:  logger,
		service: service,
	}
}

var UserControllerModule = fx.Option(fx.Provide(NewUserController))
