package main

import (
	"log"
	"os"

	"github.com/h0tf1x/gotesttask/controllers"
	"github.com/h0tf1x/gotesttask/middleware"
	"github.com/h0tf1x/gotesttask/models"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
)

// Boot application
func Boot() *iris.Application {
	app := iris.New()

	// Migrate database
	db, err := gorm.Open("mysql", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Player{}, &models.Tournament{}, &models.Team{})
	db.Close()

	// Init Middleware
	app.Use(middleware.Database)

	// Users methods
	app.Get("/take", controllers.Take)
	app.Get("/fund", controllers.Fund)
	app.Get("/balance", controllers.Balance)

	// Tournaments
	app.Get("/announceTournament", controllers.Announce)
	app.Get("/joinTournament", controllers.Join)
	app.Post("/resultTournament", controllers.Join)

	// System
	app.Get("/reset", controllers.Reset)

	return app
}

func main() {
	app := Boot()
	app.Run(iris.Addr(":8080"))
}
