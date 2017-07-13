package controllers

import (
	"fmt"

	"strings"

	"github.com/h0tf1x/gotesttask/models"
	"github.com/h0tf1x/gotesttask/response"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/context"
)

// Announce tournament
func Announce(ctx context.Context) {
	db := ctx.Values().Get("db").(*gorm.DB)
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
	db := ctx.Values().Get("db").(*gorm.DB)
	tournamentID, err := ctx.URLParamInt("tournamentId")
	if err != nil {
		ctx.JSON(response.NewErrorResponse("Tournament id required and should be integer"))
		return
	}
	tournament := models.Tournament{}
	if err := db.First(&tournament, tournamentID).Error; err != nil {
		ctx.JSON(response.NewErrorResponse("Tournament not found"))
		return
	}
	playerID, err := ctx.URLParamInt("playerId")
	if err != nil {
		ctx.JSON(response.NewErrorResponse("Player id required and should be integer"))
		return
	}
	player := models.Player{}
	if err := db.First(&player, playerID).Error; err != nil {
		ctx.JSON(response.NewErrorResponse("Player not found"))
		return
	}
	ctx.JSON(response.NewSuccessResponse("You have successfully joined tournament"))
	backerIds := strings.Split(ctx.URLParams()["backerId"], ",")
	fmt.Println(backerIds)
}

// Result - set winners
func Result(ctx context.Context) {

}
