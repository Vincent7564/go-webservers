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
	db, err := gorm.Open(mysql.Open("todolist.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.Task{})
	router.InitRouter(app, db)

	app.Listen(":3000")
}
