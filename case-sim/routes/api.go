package routes

import (
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Api() {
	userController := controllers.NewUserController()
	facades.Route().Get("/users/{id}", userController.Show)

	gameCaseController := controllers.NewGameCaseController()
	facades.Route().Get("/cases", gameCaseController.Index)

	gameItemController := controllers.NewGameItemController()
	facades.Route().Get("/items", gameItemController.Index)

	statsController := controllers.NewStatsController()
	facades.Route().Get("/stats", statsController.Index)
}
