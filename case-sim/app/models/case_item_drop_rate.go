package models

import (
	"github.com/goravel/framework/database/orm"
)

type CaseItemDropRate struct {
	orm.Model
	CaseId       string  `json:"case_id"`
	ItemId       string  `json:"item_id"`
	ItemDropRate float64 `json:"item_drop_rate"`
}
