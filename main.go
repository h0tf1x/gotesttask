package main

import (
	"github.com/h0tf1x/gotesttask/controllers"
	"github.com/h0tf1x/gotesttask/middleware"
	"github.com/kataras/iris"
)

// Boot application
func Boot() *iris.Application {
	app := iris.New()

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
