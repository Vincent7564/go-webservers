package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func printHelloWorld() {
	fmt.Println("Hello World")
}

func TestingHandler(c *fiber.Ctx) error {
	printHelloWorld()
	return nil //Pindahin ke folder Controller
}

func InitRouter(app *fiber.App) {
	app.Get("/testing", TestingHandler)
}
