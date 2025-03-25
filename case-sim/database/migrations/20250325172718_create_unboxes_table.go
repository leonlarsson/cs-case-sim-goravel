package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250325172718CreateUnboxesTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250325172718CreateUnboxesTable) Signature() string {
	return "20250325172718_create_unboxes_table"
}

// Up Run the migrations.
func (r *M20250325172718CreateUnboxesTable) Up() error {
	if !facades.Schema().HasTable("unboxes") {
		return facades.Schema().Create("unboxes", func(table schema.Blueprint) {
			table.ID()
			table.String("case_id")
			table.String("item_id")
			table.Boolean("is_stat_trak").Default(false)
			table.String("unboxer_id").Nullable()
			table.Timestamps()

			table.Foreign("case_id").References("id").On("game_cases")
			table.Foreign("item_id").References("id").On("game_items")

			table.Index("unboxer_id")
			table.Index("created_at")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250325172718CreateUnboxesTable) Down() error {
	return facades.Schema().DropIfExists("unboxes")
}
