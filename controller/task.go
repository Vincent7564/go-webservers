package controller

import (
	"fmt"
	"go-webserver/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
}

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Id          int64  `json:"id"`
}

func (cx *Controller) InsertData(c *fiber.Ctx) error {
	var task Task

	err := c.BodyParser(&task)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Request: " + err.Error())
	}

	cx.DB.Create(&model.Task{Title: task.Title, Description: task.Description})
	return nil
}

func (cx *Controller) DeleteData(c *fiber.Ctx) error {
	var task Task

	err := c.BodyParser(&task)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Request: " + err.Error())
	}

	cx.DB.Where("id = ?", task.Id).Delete(&Task{})

	return c.Status(fiber.StatusOK).SendString("Task deleted succesfully!")
}

func printHelloWorld() {
	fmt.Println("Hello World")
}

func TestingHandler(c *fiber.Ctx) error {
	printHelloWorld()
	return nil
}
