package models

import "github.com/goravel/framework/database/orm"

type Stats struct {
	orm.Model
	Name  string `json:"name"`
	Value int    `json:"value"`
}
