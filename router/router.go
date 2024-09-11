package router

import (
	"go-webserver/controller"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitRouter(app *fiber.App, db *gorm.DB) {
	contol := controller.Controller{}
	app.Get("/testing", controller.TestingHandler)
	app.Post("/insert", contol.InsertData)
}
