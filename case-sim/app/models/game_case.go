package models

import (
	"github.com/goravel/framework/database/orm"
)

type GameCase struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	orm.Timestamps
}
