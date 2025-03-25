package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20250325172901CreateCaseItemDropRatesTable struct {
}

// Signature The unique signature for the migration.
func (r *M20250325172901CreateCaseItemDropRatesTable) Signature() string {
	return "20250325172901_create_case_item_drop_rates_table"
}

// Up Run the migrations.
func (r *M20250325172901CreateCaseItemDropRatesTable) Up() error {
	if !facades.Schema().HasTable("case_item_drop_rates") {
		return facades.Schema().Create("case_item_drop_rates", func(table schema.Blueprint) {
			table.ID()
			table.String("case_id")
			table.String("item_id")
			table.Float("item_drop_rate")
			table.Timestamps()

			table.Foreign("case_id").References("id").On("game_cases").CascadeOnDelete()
			table.Foreign("item_id").References("id").On("game_items").CascadeOnDelete()

			table.Index("case_id")

			table.Unique("case_id", "item_id")
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20250325172901CreateCaseItemDropRatesTable) Down() error {
	return facades.Schema().DropIfExists("case_item_drop_rates")
}
