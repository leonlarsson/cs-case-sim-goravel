package seeders

import (
	"fmt"
	"goravel/app/models"

	"github.com/goravel/framework/facades"
)

type StatsSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *StatsSeeder) Signature() string {
	return "StatsSeeder"
}

// Run executes the seeder logic.
func (s *StatsSeeder) Run() error {
	stats := []models.Stats{
		{Name: "total_unboxes_all"},
		{Name: "total_unboxes_coverts"},
	}

	for _, stat := range stats {
		err := facades.Orm().Query().UpdateOrCreate(
			&stat,
			models.Stats{Name: stat.Name},
			models.Stats{
				Name:  stat.Name,
				Value: stat.Value,
			},
		)

		if err != nil {
			return fmt.Errorf("failed to upsert stat %s: %w", stat.Name, err)
		}
	}

	return nil
}
