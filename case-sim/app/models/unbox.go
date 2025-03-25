package models

import (
	"github.com/goravel/framework/contracts/database/factory"
	"github.com/goravel/framework/database/orm"

	"goravel/database/factories"
)

type Unbox struct {
	orm.Model
	CaseId     string `json:"case_id"`
	ItemId     string `json:"item_id"`
	IsStatTrak bool   `json:"is_stat_trak"`
	UnboxerId  string `json:"unboxer_id"`

	Case *GameCase `json:"case"`
	Item *GameItem `json:"item"`
}

func (u *Unbox) Factory() factory.Factory {
	return &factories.UnboxFactory{}
}
