package controllers

import (
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type UnboxController struct {
	// Dependent services
}

func NewUnboxController() *UnboxController {
	return &UnboxController{
		// Inject services
	}
}

func (r *UnboxController) Index(ctx http.Context) http.Response {
	var unboxes []models.Unbox
	facades.Orm().Query().With("Case").With("Item").Get(&unboxes)
	return ctx.Response().Success().Json(unboxes)
}
