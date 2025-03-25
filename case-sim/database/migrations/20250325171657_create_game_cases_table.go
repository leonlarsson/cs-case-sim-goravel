package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250325171657CreateGameCasesTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250325171657CreateGameCasesTable) Signature() string {
	return "20250325171657_create_game_cases_table"
}

// Up Run the migrations.
func (r *M20250325171657CreateGameCasesTable) Up() error {
	if !facades.Schema().HasTable("game_cases") {
		return facades.Schema().Create("game_cases", func(table schema.Blueprint) {
			table.String("id")
			table.String("type").Nullable()
			table.String("name")
			table.Text("description").Nullable()
			table.Text("image")
			table.Timestamps()

			table.Primary("id")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250325171657CreateGameCasesTable) Down() error {
	return facades.Schema().DropIfExists("game_cases")
}
