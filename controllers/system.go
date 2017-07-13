package controllers

import (
	"github.com/h0tf1x/gotesttask/models"
	"github.com/h0tf1x/gotesttask/response"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/context"
)

// Reset - reset database to initial state
func Reset(ctx context.Context) {
	db := ctx.Values().Get("db").(*gorm.DB)
	tables := [3]interface{}{&models.Team{}, &models.Tournament{}, &models.Player{}}
	for _, model := range tables {
		if err := db.Unscoped().Delete(model).Error; err != nil {
			ctx.JSON(response.NewErrorResponse("Failed reset database"))
			return
		}
	}

	for i := 1; i <= 5; i++ {
		db.Create(&models.Player{})
	}
	ctx.JSON(response.NewSuccessResponse("Database was reset to initial state"))
}
