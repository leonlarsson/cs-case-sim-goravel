package seeders

import (
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
	return facades.Orm().Query().Create([]models.Stats{
		{Name: "total_unboxes_all", Value: 0},
		{Name: "total_unboxes_coverts", Value: 0},
	})
}
