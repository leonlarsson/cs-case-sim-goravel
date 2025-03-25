package controllers

import (
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type GameItemController struct {
	// Dependent services
}

func NewGameItemController() *GameItemController {
	return &GameItemController{
		// Inject services
	}
}

func (r *GameItemController) Index(ctx http.Context) http.Response {
	var gameItems []models.GameItem
	facades.Orm().Query().Get(&gameItems)
	return ctx.Response().Success().Json(gameItems)
}
