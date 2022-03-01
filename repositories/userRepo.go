package repositories

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository struct {
	logger *zap.SugaredLogger
	dbCtx  *gorm.DB
}

func NewUserRepository(logger *zap.SugaredLogger, dbCtx *gorm.DB) *UserRepository {
	return &UserRepository{
		logger: logger,
		dbCtx:  dbCtx,
	}
}

var UserRepositoryModule = fx.Option(fx.Provide(NewUserRepository))
