package models

import (
	"github.com/jinzhu/gorm"
)

type Team struct {
	gorm.Model
	User       User
	Tournament Tournament
	Backer     bool
	Amount     int
}
