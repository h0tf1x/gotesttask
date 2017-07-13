package models

import (
	"github.com/jinzhu/gorm"
)

type Tournament struct {
	gorm.Model
	Deposit  int  `json:"deposit"`
	Finished bool `json:"finished"`
}
