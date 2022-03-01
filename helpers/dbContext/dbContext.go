package dbcontext

import (
	"fmt"

	"go.uber.org/fx"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func NewDatabaseContext() *gorm.DB {
	connString := fmt.Sprintf("")

	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

var Module = fx.Option(fx.Provide(NewDatabaseContext))
