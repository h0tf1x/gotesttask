package controllers

import (
	"github.com/h0tf1x/gotesttask/models"
	"github.com/h0tf1x/gotesttask/response"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/context"
)

// Reset - reset database to initial state
func Reset(ctx context.Context) {
	db := ctx.Value("db").(*gorm.DB)
	models := [3]interface{}{&models.Team{}, &models.Tournament{}, &models.User{}}
	for _, model := range models {
		if err := db.Delete(model).Error; err != nil {
			ctx.JSON(response.NewErrorResponse("Failed reset database"))
			return
		}
	}
	ctx.JSON(response.NewSuccessResponse("Database was reset to initial state"))
}
