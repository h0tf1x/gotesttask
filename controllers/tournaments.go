package controllers

import (
	"fmt"

	"github.com/h0tf1x/gotesttask/models"
	"github.com/h0tf1x/gotesttask/response"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/context"
)

// Announce tournament
func Announce(ctx context.Context) {
	db := ctx.Value("db").(*gorm.DB)
	deposit, err := ctx.URLParamInt("deposit")
	if err != nil {
		ctx.JSON(response.NewErrorResponse("Deposit value required and should be int"))
		return
	}
	tournament := models.Tournament{Deposit: deposit}
	if err := db.Create(&tournament).Error; err != nil {
		ctx.JSON(response.NewErrorResponse("Failed to create tournament"))
		return
	}
	ctx.JSON(response.NewDataResponse(tournament))
}

// Join - join tournament
func Join(ctx context.Context) {
	playerID, err := ctx.URLParamInt("playerId")
	if err != nil {
		ctx.JSON(response.NewErrorResponse("Player id required and should be int"))
	}
	fmt.Println(playerID)
	ctx.JSON(response.NewSuccessResponse("You have successfully joined tournament"))
}

// Result - set winners
func Result(ctx context.Context) {

}
