package main

import (
	"github.com/karim-w/gin-orm/controllers"
	dbcontext "github.com/karim-w/gin-orm/helpers/dbContext"
	"github.com/karim-w/gin-orm/repositories"
	"github.com/karim-w/gin-orm/services"
	"github.com/karim-w/gin-orm/utils/logger"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		logger.Module,
		dbcontext.Module,
		repositories.UserRepositoryModule,
		services.UserServiceModule,
		controllers.UserControllerModule,
	)
	defer app.Run()
}
