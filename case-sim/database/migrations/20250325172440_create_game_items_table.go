package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250325172440CreateGameItemsTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250325172440CreateGameItemsTable) Signature() string {
	return "20250325172440_create_game_items_table"
}

// Up Run the migrations.
func (r *M20250325172440CreateGameItemsTable) Up() error {
	if !facades.Schema().HasTable("game_items") {
		return facades.Schema().Create("game_items", func(table schema.Blueprint) {
			table.String("id")
			table.String("name")
			table.Text("description").Nullable()
			table.Text("image")
			table.String("rarity")
			table.String("phase").Nullable()
			table.Timestamps()

			table.Primary("id")

			table.Index("rarity")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250325172440CreateGameItemsTable) Down() error {
	return facades.Schema().DropIfExists("game_items")
}
