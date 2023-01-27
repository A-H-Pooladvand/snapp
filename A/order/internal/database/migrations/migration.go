package migrations

import (
	"order/internal/models"
	"order/pkg/mysql"
)

var migrations = []any{
	new(models.Order),
}

func Migrate() {
	err := mysql.New().AutoMigrate(migrations...)

	if err != nil {
		panic(err)
	}
}
