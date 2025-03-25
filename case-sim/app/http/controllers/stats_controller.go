package controllers

import (
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type StatsController struct {
	// Dependent services
}

func NewStatsController() *StatsController {
	return &StatsController{
		// Inject services
	}
}

func (r *StatsController) Index(ctx http.Context) http.Response {
	var stats []models.Stats
	facades.Orm().Query().Get(&stats)
	return ctx.Response().Success().Json(stats)
}
