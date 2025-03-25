package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250325172610CreateStatsTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250325172610CreateStatsTable) Signature() string {
	return "20250325172610_create_stats_table"
}

// Up Run the migrations.
func (r *M20250325172610CreateStatsTable) Up() error {
	if !facades.Schema().HasTable("stats") {
		return facades.Schema().Create("stats", func(table schema.Blueprint) {
			table.String("name")
			table.Integer("value").Default(0)

			table.Primary("name")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250325172610CreateStatsTable) Down() error {
	return facades.Schema().DropIfExists("stats")
}
