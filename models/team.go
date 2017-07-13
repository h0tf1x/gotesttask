package models

import (
	"github.com/jinzhu/gorm"
)

type Team struct {
	gorm.Model
	Users      []Player
	Tournament Tournament
}
