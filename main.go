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
	dsn := "root@tcp(127.0.0.1:3306)/Todolist?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.Task{})
	router.InitRouter(app, db)

	app.Listen(":3000")
}
