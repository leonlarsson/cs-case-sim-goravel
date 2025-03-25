package seeders

import (
	"goravel/app/models"

	"github.com/goravel/framework/facades"
)

type UnboxSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *UnboxSeeder) Signature() string {
	return "UnboxSeeder"
}

// Run executes the seeder logic.
func (s *UnboxSeeder) Run() error {
	var unboxes []models.Unbox
	return facades.Orm().Factory().Count(10).Create(&unboxes)
}
