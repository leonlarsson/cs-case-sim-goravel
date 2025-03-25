package database

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/contracts/database/seeder"

	"goravel/database/migrations"
	"goravel/database/seeders"
)

type Kernel struct {
}

func (kernel Kernel) Migrations() []schema.Migration {
	return []schema.Migration{
		&migrations.M20240915060148CreateUsersTable{},
		&migrations.M20250325171657CreateGameCasesTable{},
		&migrations.M20250325172440CreateGameItemsTable{},
		&migrations.M20250325172610CreateStatsTable{},
		&migrations.M20250325172718CreateUnboxesTable{},
		&migrations.M20250325172901CreateCaseItemDropRatesTable{},
	}
}

func (kernel Kernel) Seeders() []seeder.Seeder {
	return []seeder.Seeder{
		&seeders.DatabaseSeeder{},
	}
}
