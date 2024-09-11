package controller

import (
	"fmt"
	"go-webserver/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Controller struct {
	db *gorm.DB
}

func (cx *Controller) InsertData(c *fiber.Ctx) error {
	cx.db.Create(&model.Task{Title: "Clean Room", Description: "Cleaning Room"})
	return nil
}

func printHelloWorld() {
	fmt.Println("Hello World")
}

func TestingHandler(c *fiber.Ctx) error {
	printHelloWorld()
	return nil
}
