package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/karim-w/gin-orm/controllers"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func SetupRoutes(log *zap.SugaredLogger, userController *controllers.UserController) *gin.Engine {
	log.Info("Setting up routes")
	r := gin.Default()

	r.Run()
	return r
}

func registerHooks(lifecycle fx.Lifecycle, ginRouter *gin.Engine, logger *zap.SugaredLogger) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Initializing server")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Terminating server")
			logger.Sync()
			return nil
		},
	})
}

var Module = fx.Options(fx.Provide(SetupRoutes), fx.Invoke(registerHooks))
