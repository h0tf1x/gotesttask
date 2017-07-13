package models

import (
	"github.com/jinzhu/gorm"
)

type Player struct {
	gorm.Model
	Balance int `json:"balance"`
}
