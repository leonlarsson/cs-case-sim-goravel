package models

import (
	"github.com/goravel/framework/database/orm"
)

type GameItem struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Rarity      string `json:"rarity"`
	Phase       string `json:"phase"`
	orm.Timestamps
}
