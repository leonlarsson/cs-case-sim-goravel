package seeders

import (
	"encoding/json"
	"fmt"
	"goravel/app/models"
	"net/http"

	"github.com/goravel/framework/facades"
)

// Temporary struct to match the API's JSON structure
type apiGameItem struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Rarity      struct {
		Name string `json:"name"`
	} `json:"rarity"`
	Phase string `json:"phase"`
}

type GameItemSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *GameItemSeeder) Signature() string {
	return "GameItemSeeder"
}

// Run executes the seeder logic.
func (s *GameItemSeeder) Run() error {
	resp, err := http.Get("https://bymykel.github.io/CSGO-API/api/en/skins.json")
	if err != nil {
		return fmt.Errorf("failed to fetch items: %w", err)
	}
	defer resp.Body.Close()

	var apiItems []apiGameItem
	if err := json.NewDecoder(resp.Body).Decode(&apiItems); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	// Convert API items to model items
	items := make([]models.GameItem, len(apiItems))
	for i, apiItem := range apiItems {
		items[i] = models.GameItem{
			Id:          apiItem.Id,
			Name:        apiItem.Name,
			Description: apiItem.Description,
			Image:       apiItem.Image,
			Rarity:      apiItem.Rarity.Name, // Extract rarity name
			Phase:       apiItem.Phase,
		}
	}

	fmt.Printf("Found %d  items. Inserting...\n", len(items))

	// For some reason, UpdateOrCreate doesn't accept a slice like Create does
	// return facades.Orm().Query().UpdateOrCreate(&items, models.GameItem{Id: "id"}, models.GameItem{Name: "name", Description: "description", Image: "image", Rarity: "rarity", Phase: "phase"})

	for _, gameIem := range items {
		err := facades.Orm().Query().UpdateOrCreate(
			&gameIem,
			models.GameItem{Id: gameIem.Id},
			models.GameItem{
				Name:        gameIem.Name,
				Description: gameIem.Description,
				Image:       gameIem.Image,
				Rarity:      gameIem.Rarity,
				Phase:       gameIem.Phase,
			},
		)
		if err != nil {
			return fmt.Errorf("failed to upsert item %s: %w", gameIem.Id, err)
		}

	}

	return nil
}
