package main

import (
	"go-webserver/model"
	"go-webserver/router"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	db, err := gorm.Open(mysql.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&model.Product{})
	if err != nil {
		panic("failed to connect database")
	}
	router.InitRouter(app)

	app.Listen(":3000")
}
