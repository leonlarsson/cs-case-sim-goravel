package controllers

import (
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type GameCaseController struct {
	// Dependent services
}

func NewGameCaseController() *GameCaseController {
	return &GameCaseController{
		// Inject services
	}
}

func (r *GameCaseController) Index(ctx http.Context) http.Response {
	var gameCases []models.GameCase
	facades.Orm().Query().Get(&gameCases)
	return ctx.Response().Success().Json(gameCases)
}
