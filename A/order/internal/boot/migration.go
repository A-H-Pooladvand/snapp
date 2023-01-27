package boot

import (
	"order/internal/database/migrations"
	"order/internal/database/seeders"
)

type Migrator struct {
}

// Boot runs application's migrations
func (m Migrator) Boot() error {
	migrations.Migrate()
	seeders.Run()

	return nil
}
