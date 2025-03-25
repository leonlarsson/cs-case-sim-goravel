package seeders

import (
	"encoding/json"
	"fmt"
	"goravel/app/models"
	"net/http"

	"github.com/goravel/framework/facades"
)

type GameCaseSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *GameCaseSeeder) Signature() string {
	return "GameCaseSeeder"
}

// Run executes the seeder logic.
func (s *GameCaseSeeder) Run() error {
	res, err := http.Get("https://bymykel.github.io/CSGO-API/api/en/crates.json")
	if err != nil {
		return fmt.Errorf("failed to fetch cases: %w", err)
	}
	defer res.Body.Close()

	var cases []models.GameCase
	err = json.NewDecoder(res.Body).Decode(&cases)
	if err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	fmt.Printf("Found %d cases. Inserting...\n", len(cases))

	// For some reason, UpdateOrCreate doesn't accept a slice like Create does
	// return facades.Orm().Query().UpdateOrCreate(&cases, models.GameCase{Id: "id"}, models.GameCase{Name: "name", Description: "description", Image: "image"})

	for _, gameCase := range cases {
		err := facades.Orm().Query().UpdateOrCreate(
			&gameCase,
			models.GameCase{Id: gameCase.Id},
			models.GameCase{
				Name:        gameCase.Name,
				Description: gameCase.Description,
				Image:       gameCase.Image,
				Type:        gameCase.Type,
			},
		)
		if err != nil {
			return fmt.Errorf("failed to upsert case %s: %w", gameCase.Id, err)
		}
	}

	return nil
}
