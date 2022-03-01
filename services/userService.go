package services

import (
	"github.com/karim-w/gin-orm/repositories"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type UserService struct {
	logger         *zap.SugaredLogger
	userRepository *repositories.UserRepository
}

func NewUserService(logger *zap.SugaredLogger, userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		logger:         logger,
		userRepository: userRepository,
	}
}

var UserServiceModule = fx.Option(fx.Provide(NewUserService))
