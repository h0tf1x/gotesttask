package middleware

import (
	"os"

	"log"

	"github.com/h0tf1x/gotesttask/response"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris/context"
)

// Database connection middleware
func Database(ctx context.Context) {
	db, err := gorm.Open("mysql", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
		ctx.JSON(response.NewErrorResponse("Failed connect to database"))
		return
	}
	db.LogMode(true)

	defer db.Close()

	ctx.Values().Set("db", db)

	ctx.Next()
}
